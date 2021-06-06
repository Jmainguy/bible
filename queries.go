package main

import (
	"database/sql"
	"fmt"
)

func rowQuery(query string, db *sql.DB) *sql.Row {
	row := db.QueryRow(query)
	return row
}

func rowsQuery(query string, db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(query)
	return rows, err
}

func runQuery(query string, db *sql.DB, translationName string) {
	rows, err := db.Query(query)
	check(err)

	verse := Verse{}
	bookMap := mapIDToBook(db)
	for rows.Next() {
		err = rows.Scan(&verse.ID, &verse.Book, &verse.Chapter, &verse.Verse, &verse.Text)
		check(err)
		fmt.Printf("%s %d:%d - %s\n", bookMap[verse.Book], verse.Chapter, verse.Verse, translationName)
		fmt.Println(" ", verse.Text)
		fmt.Println()
	}
}
