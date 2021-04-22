package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getVerseID(book int, chapterVerse string, db *sql.DB, desired string, version string) (verseID int) {
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
			panic(err)
		}
		verseID = verseIDInside
		//  Just a chaper, 3
	} else if len(chapterVerseArray) == 1 {
		// Return max verse for this chapter
		if desired == "max" {
			query := fmt.Sprintf("select id from %s where b is %d and c is %d;", version, book, chapter)
			rows, err := db.Query(query)
			check(err)

			verse := Verse{}
			for rows.Next() {
				err = rows.Scan(&verse.ID)
				check(err)
				verseID = verse.ID
			}
		} else {
			verseIDString = fmt.Sprintf("%02d%03d001", book, chapter)
			verseIDInside, err := strconv.Atoi(verseIDString)
			if err != nil {
				panic(err)
			}
			verseID = verseIDInside
		}
	}

	return
}

func parseVerseString(verseString, version string, db *sql.DB) (query string) {
	// Book
	bookID, bookName := getBook(verseString, db)

	// Must be a valid book to get this far.
	// Cut book name from requested query
	verseString = strings.TrimPrefix(verseString, bookName)
	verseArrayWithSpaces := strings.Split(verseString, " ")
	// Remove all blank spaces from array
	var verseArray []string
	for _, v := range verseArrayWithSpaces {
		if v != "" {
			// Bomb out if any non numeric or colons or dash
			validString := checkString(v)
			if !validString {
				fmt.Printf("Verse contains invalid characters, %s\n", v)
				os.Exit(1)
			}
			verseArray = append(verseArray, v)
		}
	}
	switch {
	// 1:1 - 3:16
	case len(verseArray) == 3:
		// second indice should be a - always
		if verseArray[1] != "-" {
			fmt.Printf("Range is invalid, must have a - in the middle of the two verses, %s\n", verseString)
			os.Exit(1)
		}
		// Get chapter and verse strings
		beginID := getVerseID(bookID, verseArray[0], db, "min", version)
		endID := getVerseID(bookID, verseArray[2], db, "max", version)
		query = fmt.Sprintf("SELECT * FROM %s WHERE id BETWEEN %d AND %d;", version, beginID, endID)

	// 3:16
	// 3
	case len(verseArray) == 1:
		// John 3
		if len(strings.Split(verseArray[0], ":")) == 1 {
			validChapter := checkNumeric(verseArray[0])
			if validChapter {
				query = fmt.Sprintf("SELECT * FROM %s WHERE b is %d and c is %s;", version, bookID, verseArray[0])
			}
		} else {
			verseID := getVerseID(bookID, verseArray[0], db, "max", version)
			query = fmt.Sprintf("SELECT * FROM %s WHERE id is %d;", version, verseID)
		}
	// John
	case len(verseArray) == 0:
		query = fmt.Sprintf("SELECT * FROM %s WHERE b is %d;", version, bookID)
	default:
		query = fmt.Sprintf("Length was %d, array is %v\n", len(verseArray), verseArray)
	}

	return
}

func getBook(verseString string, db *sql.DB) (bookID int, bookString string) {
	bookMap := mapIDToBook(db)

	for key, bookName := range bookMap {
		if strings.HasPrefix(strings.ToLower(verseString), strings.ToLower(bookName)) {
			bookID = key
			bookString = bookName
		}
	}
	if bookID == 0 {
		fmt.Println(verseString)
		fmt.Printf("Unable to find book in the bible\n")
		fmt.Println("Valid Books")
		for _, bookName := range bookMap {
			fmt.Println(bookName)
		}
		os.Exit(1)
	}

	return bookID, bookString
}
