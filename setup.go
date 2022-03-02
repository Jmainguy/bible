package main

import (
	"database/sql"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

// Config : Struct of common things bible needs
type Config struct {
	Passage               string                 `json:"passage"`
	HomeDir               string                 `json:"homeDir"`
	Translation           string                 `json:"translation"`
	Database              string                 `json:"database"`
	DB                    *sql.DB                `json:"db"`
	Compare               string                 `json:"compare"`
	ListBooks             bool                   `json:"listBooks"`
	ListTranslations      bool                   `json:"listTranslations"`
	RandomPassage         bool                   `json:"randomPassage"`
	RandomChapter         bool                   `json:"randomChapter"`
	BookMap               map[int]string         `json:"bookMap"`
	BookList              []string               `json:"booklist"`
	BookInfoList          []BookInfo             `json:"bookinfolist"`
	TranslationsAvailable map[string]Translation `json:"transitionsAvailable"`
	Verse                 Verse                  `json:"verse"`
}

func setup() (config Config, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return config, err
	}
	defaultBibleDB := homeDir + "/.bible/bible.db"

	passagePtr := flag.String("passage", "John 3:16", "Passage to return. Can be given in following syntax. 'John', '1 John 3', 'John 3:16', or for a range in the same book '1 John 1:1 - 3:16'")
	translationPtr := flag.String("translation", "t_kjv", "Bible translation to use")
	databasePtr := flag.String("db", defaultBibleDB, "Bible database to use")
	comparePtr := flag.String("compare", "", "Translations to compare passage against, set to 'all' for all translations in the database, or use a space seperated list")
	listBooksPtr := flag.Bool("listBooks", false, "List all books of the bible and their number of chapters")
	listTranslationsPtr := flag.Bool("listTranslations", false, "List all translations in database")
	randomPassagePtr := flag.Bool("randomPassage", false, "Print a random passage")
	randomChapterPtr := flag.Bool("randomChapter", false, "Print a random Chapter")

	flag.Parse()

	// Download DB if not present
	_, err = os.Stat(*databasePtr)
	if err != nil {
		fmt.Printf("Database %s does not seem to exist\n", *databasePtr)
		fmt.Println("Will download it now")
		err = downloadDB(*databasePtr)
		if err != nil {
			return config, err
		}
	}

	// Open sqlite3 database containing the bibles
	config.DB, err = sql.Open("sqlite", *databasePtr)
	if err != nil {
		return config, err
	}

	config.Verse = Verse{}
	config.BookMap, err = mapIDToBook(config.DB)
	if err != nil {
		return config, err
	}

	rand.Seed(time.Now().UnixNano())

	config.Passage = *passagePtr
	config.Translation = *translationPtr
	config.Compare = *comparePtr
	config.ListBooks = *listBooksPtr
	config.ListTranslations = *listTranslationsPtr
	config.RandomPassage = *randomPassagePtr
	config.RandomChapter = *randomChapterPtr
	config.Database = *databasePtr

	config.BookList, err = listBooks(config.DB)
	if err != nil {
		return config, err
	}
	config.BookInfoList, err = getBooks(config.DB)
	if err != nil {
		return config, err
	}

	config.TranslationsAvailable, err = getTranslations(config.DB)
	if err != nil {
		return config, err
	}

	return config, nil
}
