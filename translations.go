package main

import (
	"database/sql"
)

func getTranslations(db *sql.DB) (map[string]Translation, error) {
	translations := make(map[string]Translation)
	query := `SELECT "table","version" FROM bible_version_key;`
	rows, err := db.Query(query)
	if err != nil {
		return translations, err
	}
	for rows.Next() {
		translation := Translation{}
		err = rows.Scan(&translation.Table, &translation.Version)
		if err != nil {
			return translations, err
		}
		translations[translation.Table] = translation
	}
	return translations, nil
}
