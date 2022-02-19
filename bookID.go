package main

import (
	"database/sql"
	"fmt"
)

func mapIDToBook(db *sql.DB) map[int]string {
	bookMap := make(map[int]string)
	query := `SELECT b,n FROM key_english;`
	rows, err := db.Query(query)
	dbCheck(err)
	keyEnglish := KeyEnglish{}
	for rows.Next() {
		err = rows.Scan(&keyEnglish.ID, &keyEnglish.Book)
		check(err)
		bookMap[keyEnglish.ID] = keyEnglish.Book
	}
	return bookMap
}

func getBooks(db *sql.DB) []BookInfo {
	var books []BookInfo
	query := `SELECT "order","title_short","otnt","chapters" FROM book_info;`
	rows, err := db.Query(query)
	dbCheck(err)
	for rows.Next() {
		bookInfo := BookInfo{}
		err = rows.Scan(&bookInfo.ID, &bookInfo.Title, &bookInfo.Testament, &bookInfo.Chapters)
		check(err)
		if bookInfo.Testament == "OT" {
			bookInfo.Testament = "Old Testament"
		} else {
			bookInfo.Testament = "New Testament"
		}
		books = append(books, bookInfo)
	}
	return books
}

func listBooks(db *sql.DB) {
	books := getBooks(db)
	// {41 Mark New Testament 16}
	for _, v := range books {
		if v.Chapters > 1 {
			fmt.Printf("%s, %d Chapters, %s, The %d book of the bible\n", v.Title, v.Chapters, v.Testament, v.ID)
		} else {
			fmt.Printf("%s, %d Chapter, %s, The %d book of the bible\n", v.Title, v.Chapters, v.Testament, v.ID)
		}
	}
}
