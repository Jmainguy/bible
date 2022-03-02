package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

func getVerseID(book int, chapterVerse string, db *sql.DB, minmax string, version string) (verseID int, err error) {
	chapterVerseArray := strings.Split(chapterVerse, ":")
	chapter, _ := strconv.Atoi(chapterVerseArray[0])
	var verse int
	var verseIDString string
	// Chapter and Verse, 3:16
	if len(chapterVerseArray) == 2 {
		verse, _ = strconv.Atoi(chapterVerseArray[1])
		verseIDString = fmt.Sprintf("%02d%03d%03d", book, chapter, verse)
		verseIDInside, err := strconv.Atoi(verseIDString)
		if err != nil {
			return 0, err
		}
		verseID = verseIDInside
		//  Just a chaper, 3
	} else if len(chapterVerseArray) == 1 {
		// Return max verse for this chapter
		if minmax == "max" {
			query := fmt.Sprintf("select id from %s where b is %d and c is %d;", version, book, chapter)
			rows, err := db.Query(query)
			if err != nil {
				return 0, err
			}

			verse := Verse{}
			for rows.Next() {
				err = rows.Scan(&verse.ID)
				if err != nil {
					return 0, err
				}
				verseID = verse.ID
			}
		} else if minmax == "min" {
			verseIDString = fmt.Sprintf("%02d%03d001", book, chapter)
			verseIDInside, err := strconv.Atoi(verseIDString)
			if err != nil {
				return 0, err
			}
			verseID = verseIDInside
		} else {
			err := fmt.Errorf("minmax must be set to min or max, was set to %s", minmax)
			return 0, err
		}
	}

	return verseID, nil
}

func parseVerseString(config Config, translation string) (query string, err error) {
	// Check for valid translation
	_, okay := config.TranslationsAvailable[translation]
	if !okay {
		err = fmt.Errorf("Translation is invalid")
		return "", err

	}
	// Book
	bookID, bookName, err := getBook(config)
	if err != nil {
		return "", err
	}
	// Must be a valid book to get this far.
	// Cut book name from requested query
	config.Passage = strings.TrimPrefix(config.Passage, bookName)
	verseArrayWithSpaces := strings.Split(config.Passage, " ")
	// Remove all blank spaces from array
	var verseArray []string
	for _, v := range verseArrayWithSpaces {
		if v != "" {
			// Bomb out if any non numeric or colons or dash
			validString := checkNumericPlusColonDash(v)
			if !validString {
				err := fmt.Errorf("Verse contains invalid characters, %s", v)
				return "", err
			}
			verseArray = append(verseArray, v)
		}
	}
	switch {
	// 1:1 - 3:16
	case len(verseArray) == 3:
		// second indice should be a - always
		if verseArray[1] != "-" {
			err := fmt.Errorf("range is invalid, must have a - in the middle of the two verses, %s", config.Passage)
			return "", err
		}
		// Get chapter and verse strings
		beginID, err := getVerseID(bookID, verseArray[0], config.DB, "min", translation)
		if err != nil {
			return "", err
		}
		endID, err := getVerseID(bookID, verseArray[2], config.DB, "max", translation)
		if err != nil {
			return "", err
		}
		query = fmt.Sprintf("SELECT * FROM %s WHERE id BETWEEN %d AND %d;", translation, beginID, endID)

	// 3:16
	// 3
	case len(verseArray) == 1:
		// John 3
		if len(strings.Split(verseArray[0], ":")) == 1 {
			validChapter := checkNumeric(verseArray[0])
			if validChapter {
				query = fmt.Sprintf("SELECT * FROM %s WHERE b is %d and c is %s;", translation, bookID, verseArray[0])
			}
		} else {
			verseID, err := getVerseID(bookID, verseArray[0], config.DB, "max", translation)
			if err != nil {
				return "", err
			}
			query = fmt.Sprintf("SELECT * FROM %s WHERE id is %d;", translation, verseID)
		}
	// John
	case len(verseArray) == 0:
		query = fmt.Sprintf("SELECT * FROM %s WHERE b is %d;", translation, bookID)
	default:
		err = fmt.Errorf("length was %d, array is %v", len(verseArray), verseArray)
		return "", err
	}

	return query, nil
}

func getBook(config Config) (bookID int, bookString string, err error) {
	bookMap, err := mapIDToBook(config.DB)
	if err != nil {
		return bookID, bookString, err
	}

	for key, bookName := range bookMap {
		if strings.HasPrefix(strings.ToLower(config.Passage), strings.ToLower(bookName)) {
			bookID = key
			bookString = bookName
		}
	}
	if bookID == 0 {
		err = fmt.Errorf("unable to find book in the bible, run 'bible -listBooks' for a valid list")
	}

	return bookID, bookString, err
}
