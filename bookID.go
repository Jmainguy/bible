package main

import (
	"database/sql"
)

func mapIDToBook(db *sql.DB) map[int]string {
	bookMap := make(map[int]string)
	query := `SELECT b,n FROM key_english;`
	rows, err := db.Query(query)
	check(err)
	keyEnglish := KeyEnglish{}
	for rows.Next() {
		err = rows.Scan(&keyEnglish.ID, &keyEnglish.Book)
		check(err)
		bookMap[keyEnglish.ID] = keyEnglish.Book
	}
	return bookMap
}
