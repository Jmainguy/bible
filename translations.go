package main

import (
	"database/sql"
	"fmt"
)

func listTranslations(db *sql.DB) {
	var translations []Translation
	query := `SELECT "table","version" FROM bible_version_key;`
	rows, err := db.Query(query)
	check(err)
	for rows.Next() {
		translation := Translation{}
		err = rows.Scan(&translation.Table, &translation.Version)
		check(err)
		translations = append(translations, translation)
	}
	for _, v := range translations {
		fmt.Printf("%s: %s\n", v.Table, v.Version)
	}
}
