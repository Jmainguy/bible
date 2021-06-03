package main

import (
	"database/sql"
	"flag"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	verseStringPtr := flag.String("verse", "John 3:16", "Verse to return. Can be given in following syntax. 'John', '1 John 3', 'John 3:16', or for a range in the same book '1 John 1:1 - 3:16'")
	versionPtr := flag.String("version", "t_kjv", "Bible Version to use")
	databasePtr := flag.String("db", "database/bible-sqlite-jmainguy.db", "Bible database to use")
	generateTestsPtr := flag.Bool("generateTests", false, "Whether to generate and print tests to stdout")
	listBooksPtr := flag.Bool("listBooks", false, "List all books of the bible and their number of chapters")
	listTranslationsPtr := flag.Bool("listTranslations", false, "List all translations in database")
	flag.Parse()

	// Open sqlite3 database containing the bibles
	db, err := sql.Open("sqlite3", *databasePtr)
	check(err)

	if *generateTestsPtr {
		generateTests(db)
	} else if *listBooksPtr {
		listBooks(db)
	} else if *listTranslationsPtr {
		listTranslations(db)
	} else {
		query := parseVerseString(*verseStringPtr, *versionPtr, db)
		runQuery(query, db)
	}
}
