package main

import (
	"fmt"
	"math/rand"
)

func randomPassage(config Config) (Config, error) {
	idQuery := "SELECT ID from t_kjv;"
	allIDs, err := rowsQuery(idQuery, config.DB)
	if err != nil {
		return config, err
	}
	var randomIDs []int
	var id int
	for allIDs.Next() {
		err := allIDs.Scan(&id)
		if err != nil {
			return config, err
		}

		randomIDs = append(randomIDs, id)
	}
	randomID := randomIDs[rand.Intn(len(randomIDs))]
	query := fmt.Sprintf("SELECT * FROM %s WHERE id is %d;", config.Translation, randomID)
	passage := rowQuery(query, config.DB)

	err = passage.Scan(&config.Verse.ID, &config.Verse.Book, &config.Verse.Chapter, &config.Verse.Verse, &config.Verse.Text)
	if err != nil {
		return config, err
	}
	config.Passage = fmt.Sprintf("%s %d:%d", config.BookMap[config.Verse.Book], config.Verse.Chapter, config.Verse.Verse)
	return config, nil
}

func randomChapter(config Config) (Config, error) {
	book := config.BookInfoList[rand.Intn(len(config.BookInfoList))]
	chapter := rand.Intn(book.Chapters)
	if chapter == 0 {
		chapter++
	}
	config.Passage = fmt.Sprintf("%s %d", book.Title, chapter)
	return config, nil
}
