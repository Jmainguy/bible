package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func getHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	return dirname
}

func downloadDB(file string) error {

	// Ensure dir exists, if not create
	dirpath := filepath.Dir(file)
	dir, err := os.Stat(dirpath)
	if err != nil {
		err = os.MkdirAll(dirpath, os.ModePerm)
		if err != nil {
			fmt.Printf("So, %s didnt exist, we tried to download and create it, but ran into an error\n", file)
			fmt.Printf("Specifically, %s the directory didnt exist, and we got this error when trying to create it\n", dirpath)
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		if !dir.IsDir() {
			err = os.Remove(dirpath)
			if err != nil {
				fmt.Printf("So, %s didnt exist, we tried to download and create it, but ran into an error\n", file)
				fmt.Printf("Specifically, %s existed but was not a directory, so we tried to delete it, then recreate it as a dir but ran into this error\n", dirpath)
				fmt.Println(err)
				os.Exit(1)
			}
			err = os.MkdirAll(dirpath, os.ModePerm)
			if err != nil {
				fmt.Printf("So, %s didnt exist, we tried to download and create it, but ran into an error\n", file)
				fmt.Printf("Specifically, %s existed but was not a directory, so we tried to delete it, then recreate it as a dir but ran into this error\n", dirpath)
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	// Get the data
	resp, err := http.Get("https://github.com/Jmainguy/bible/raw/main/database/bible.db")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err

}

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
