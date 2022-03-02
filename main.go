package main

import (
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	config, err := setup()
	checkErrorPrintExitOne(err)

	if config.RandomChapter {
		config, err = randomChapter(config)
		checkErrorPrintExitOne(err)
	} else if config.RandomPassage {
		config, err = randomPassage(config)
		checkErrorPrintExitOne(err)
	}

	if config.ListBooks {
		for _, book := range config.BookList {
			fmt.Println(book)
		}
	} else if config.ListTranslations {
		translations, err := getTranslations(config.DB)
		checkErrorPrintExitOne(err)
		for _, v := range translations {
			fmt.Printf("%s: %s\n", v.Table, v.Version)
		}
	} else if config.Compare != "" {
		err := compare(config)
		checkErrorPrintExitOne(err)
	} else {
		err := printPassage(config)
		checkErrorPrintExitOne(err)
	}
}
