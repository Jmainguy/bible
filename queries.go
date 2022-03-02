package main

import (
	"database/sql"
)

func rowQuery(query string, db *sql.DB) *sql.Row {
	row := db.QueryRow(query)
	return row
}

func rowsQuery(query string, db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(query)
	return rows, err
}
