package main

import (
	"database/sql"
	"fmt"
)

func printTranslations(translations map[string]Translation) {
	for _, v := range translations {
		fmt.Printf("%s: %s\n", v.Table, v.Version)
	}
}

func getTranslations(db *sql.DB) map[string]Translation {
	translations := make(map[string]Translation)
	query := `SELECT "table","version" FROM bible_version_key;`
	rows, err := db.Query(query)
	dbCheck(err)
	for rows.Next() {
		translation := Translation{}
		err = rows.Scan(&translation.Table, &translation.Version)
		check(err)
		translations[translation.Table] = translation
	}
	return translations
}
