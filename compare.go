package main

import (
	"fmt"
	"strings"
)

func compare(config Config) error {
	if config.Compare == "all" {
		for _, v := range config.TranslationsAvailable {
			query, err := parseVerseString(config, v.Table)
			if err != nil {
				return err
			}
			rows, err := config.DB.Query(query)
			if err != nil {
				return err
			}
			for rows.Next() {
				err = rows.Scan(&config.Verse.ID, &config.Verse.Book, &config.Verse.Chapter, &config.Verse.Verse, &config.Verse.Text)
				if err != nil {
					return err
				}
				fmt.Printf("%s %d:%d - %s\n", config.BookMap[config.Verse.Book], config.Verse.Chapter, config.Verse.Verse, v.Version)
				fmt.Println(" ", config.Verse.Text)
				fmt.Println()
			}

		}
	} else {
		translationsWantedList := strings.Fields(config.Compare)
		for _, v := range translationsWantedList {
			translation, okay := config.TranslationsAvailable[v]
			if okay {
				query, err := parseVerseString(config, v)
				if err != nil {
					return err
				}
				rows, err := config.DB.Query(query)
				if err != nil {
					return err
				}
				for rows.Next() {
					err = rows.Scan(&config.Verse.ID, &config.Verse.Book, &config.Verse.Chapter, &config.Verse.Verse, &config.Verse.Text)
					if err != nil {
						return err
					}
					fmt.Printf("%s %d:%d - %s\n", config.BookMap[config.Verse.Book], config.Verse.Chapter, config.Verse.Verse, translation.Version)
					fmt.Println(" ", config.Verse.Text)
					fmt.Println()
				}

			} else {
				err := fmt.Errorf("%s is not a translation in the %s database", v, config.Database)
				return err
			}
		}
	}
	return nil

}
