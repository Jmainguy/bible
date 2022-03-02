package main

import (
	"database/sql"
	"fmt"
)

func mapIDToBook(db *sql.DB) (bookMap map[int]string, err error) {
	bookMap = make(map[int]string)
	query := `SELECT b,n FROM key_english;`
	rows, err := db.Query(query)
	if err != nil {
		return bookMap, err
	}
	keyEnglish := KeyEnglish{}
	for rows.Next() {
		err = rows.Scan(&keyEnglish.ID, &keyEnglish.Book)
		if err != nil {
			return bookMap, err
		}
		bookMap[keyEnglish.ID] = keyEnglish.Book
	}
	return bookMap, err
}

func getBooks(db *sql.DB) (books []BookInfo, err error) {
	query := `SELECT "order","title_short","otnt","chapters" FROM book_info;`
	rows, err := db.Query(query)
	if err != nil {
		return books, err
	}
	for rows.Next() {
		bookInfo := BookInfo{}
		err = rows.Scan(&bookInfo.ID, &bookInfo.Title, &bookInfo.Testament, &bookInfo.Chapters)
		if err != nil {
			return books, err
		}
		if bookInfo.Testament == "OT" {
			bookInfo.Testament = "Old Testament"
		} else {
			bookInfo.Testament = "New Testament"
		}
		books = append(books, bookInfo)
	}
	return books, err
}

func listBooks(db *sql.DB) (msgs []string, err error) {
	books, err := getBooks(db)
	if err != nil {
		return msgs, err
	}

	// {41 Mark New Testament 16}
	for _, v := range books {
		if v.Chapters > 1 {
			msg := fmt.Sprintf("%s, %d Chapters, %s, The %d book of the bible", v.Title, v.Chapters, v.Testament, v.ID)
			msgs = append(msgs, msg)
		} else {
			msg := fmt.Sprintf("%s, %d Chapter, %s, The %d book of the bible", v.Title, v.Chapters, v.Testament, v.ID)
			msgs = append(msgs, msg)

		}
	}
	return msgs, err
}
