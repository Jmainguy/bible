package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"math/rand"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	homeDir := getHomeDir()
	defaultBibleDB := homeDir + "/.bible/bible.db"

	passagePtr := flag.String("passage", "John 3:16", "Passage to return. Can be given in following syntax. 'John', '1 John 3', 'John 3:16', or for a range in the same book '1 John 1:1 - 3:16'")
	translationPtr := flag.String("translation", "t_kjv", "Bible translation to use")
	databasePtr := flag.String("db", defaultBibleDB, "Bible database to use")
	comparePtr := flag.String("compare", "", "Translations to compare passage against, set to 'all' for all translations in the database, or use a space seperated list")
	generateTestsPtr := flag.Bool("generateTests", false, "Whether to generate and print tests to stdout")
	listBooksPtr := flag.Bool("listBooks", false, "List all books of the bible and their number of chapters")
	listTranslationsPtr := flag.Bool("listTranslations", false, "List all translations in database")
	randomPassagePtr := flag.Bool("randomPassage", false, "Print a random passage")
	randomChapterPtr := flag.Bool("randomChapter", false, "Print a random Chapter")

	flag.Parse()

	// See if DB even exists, otherwise quit
	_, err := os.Stat(*databasePtr)
	if err != nil {
		fmt.Printf("Database %s does not seem to exist\n", *databasePtr)
		fmt.Println("Will download it now")
		err = downloadDB(*databasePtr)
		if err != nil {
			panic(err)
		}
	}

	// Open sqlite3 database containing the bibles
	db, err := sql.Open("sqlite3", *databasePtr)
	if err != nil {
		fmt.Println("Had trouble opening database")
		fmt.Println(err)
		os.Exit(1)
	}
	rand.Seed(time.Now().UnixNano())

	if *randomChapterPtr {
		books := getBooks(db)
		book := books[rand.Intn(len(books))]
		chapter := rand.Intn(book.Chapters)
		if chapter == 0 {
			chapter++
		}
		*passagePtr = fmt.Sprintf("%s %d", book.Title, chapter)
	} else if *randomPassagePtr {
		idQuery := "SELECT ID from t_kjv;"
		allIDs, err := rowsQuery(idQuery, db)
		if err != nil {
			fmt.Println("Unable to return a random passage")
			fmt.Println(err)
			os.Exit(1)
		}
		var randomIDs []int
		var id int
		for allIDs.Next() {
			err := allIDs.Scan(&id)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			randomIDs = append(randomIDs, id)
		}
		randomID := randomIDs[rand.Intn(len(randomIDs))]
		query := fmt.Sprintf("SELECT * FROM %s WHERE id is %d;", *translationPtr, randomID)
		passage := rowQuery(query, db)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		verse := Verse{}
		bookMap := mapIDToBook(db)
		err = passage.Scan(&verse.ID, &verse.Book, &verse.Chapter, &verse.Verse, &verse.Text)
		check(err)
		*passagePtr = fmt.Sprintf("%s %d:%d", bookMap[verse.Book], verse.Chapter, verse.Verse)
	}

	if *generateTestsPtr {
		generateTests(db)
	} else if *listBooksPtr {
		listBooks(db)
	} else if *listTranslationsPtr {
		translations := getTranslations(db)
		printTranslations(translations)
	} else if *comparePtr != "" {
		translationsAvailable := getTranslations(db)
		if *comparePtr == "all" {
			for _, v := range translationsAvailable {
				query := parseVerseString(*passagePtr, v.Table, db)
				runQuery(query, db, v.Version)
			}
		} else {
			translationsWantedList := strings.Fields(*comparePtr)
			for _, v := range translationsWantedList {
				translation, okay := translationsAvailable[v]
				if okay {
					query := parseVerseString(*passagePtr, v, db)
					runQuery(query, db, translation.Version)
				} else {
					fmt.Printf("%s is not a translation in the %s database\n", v, *databasePtr)
				}
			}
		}
	} else {
		query := parseVerseString(*passagePtr, *translationPtr, db)
		translationsAvailable := getTranslations(db)
		translation, okay := translationsAvailable[*translationPtr]
		if okay {
			runQuery(query, db, translation.Version)
		} else {
			fmt.Printf("%s is not a translation in the %s database\n", *translationPtr, *databasePtr)
		}
	}
}
