package main

import (
	"fmt"
)

func printPassage(config Config) error {
	query, err := parseVerseString(config, config.Translation)
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
		fmt.Printf("%s %d:%d - %s\n", config.BookMap[config.Verse.Book], config.Verse.Chapter, config.Verse.Verse, config.Translation)
		fmt.Println(" ", config.Verse.Text)
		fmt.Println()
	}
	return nil
}
