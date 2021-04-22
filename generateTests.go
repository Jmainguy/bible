package main

import (
	"database/sql"
	"fmt"
	"github.com/divan/num2words"
	"github.com/iancoleman/strcase"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

func returnSingleInt(query string) (int, error) {
	var result int
	db, err := sql.Open("sqlite3", "database/bible-sqlite-jmainguy.db?cache=shared&mode=memory")
	if err != nil {
		return result, err
	}
	row := rowQuery(query, db)
	err = row.Scan(&result)
	if err != nil {
		return result, err
	}
	db.Close()
	return result, err
}

func longNames(book string) string {
	book = strings.ReplaceAll(book, "1 ", "First")
	book = strings.ReplaceAll(book, "2 ", "Second")
	book = strings.ReplaceAll(book, "3 ", "Third")
	book = strings.ReplaceAll(book, " of ", "Of")
	return book
}

func generateTests(db *sql.DB) {
	i := 1
	booksInBible := 66
	bookMap := mapIDToBook(db)
	for i <= booksInBible {
		bookName := longNames(bookMap[i])
		query := fmt.Sprintf(`select count(distinct c) from t_kjv where b = %d;`, i)
		chapters, err := returnSingleInt(query)
		check(err)
		test := fmt.Sprintf(`
func Test%sChapterCount(t *testing.T) {
        expectedCount := %d
        query :=  `+`"`+"select `table` "+`from bible_version_key;"`+`
        result := returnArray(t, query)
        bookID := %d
        for _, bibleVersion := range result {
            query := fmt.Sprintf("select count(distinct c) from %%s where b = %%d;", bibleVersion, bookID)
            actualCount, err := returnSingleInt(query)
            assert.Nil(t, err)
            msg := fmt.Sprintf("Should only be %%d chapters in %%s book in %%s translation", expectedCount, bookMap[bookID], bibleVersion)
            assert.Equal(t, expectedCount, actualCount, msg)
        }
}`, bookName, chapters, i)
		fmt.Println(test)
		// Test amount of verses
		currentChapter := 1
		for currentChapter <= chapters {
			prettyChapterName := num2words.Convert(currentChapter)
			prettyChapterName = strings.ReplaceAll(prettyChapterName, " ", "")
			prettyChapterName = strcase.ToCamel(prettyChapterName)
			query := fmt.Sprintf(`select count(distinct v) from t_kjv where b = %d and c = %d;`, i, currentChapter)
			expectedVerseCount, err := returnSingleInt(query)
			check(err)
			verseTest := fmt.Sprintf(`
func Test%sChapter%sTotalVerseCount(t *testing.T) {
        expectedCount := %d
        query :=  `+`"`+"select `table` "+`from bible_version_key;"`+`
        result := returnArray(t, query)
        bookID := %d
        chapterID := %d
        for _, bibleVersion := range result {
            query := fmt.Sprintf("select count(distinct v) from %%s where b = %%d and c = %%d;", bibleVersion, bookID, chapterID)
            actualCount, err := returnSingleInt(query)
            assert.Nil(t, err)
            msg := fmt.Sprintf("Should only be %%d verses in chapter %%d of %%s book in %%s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
            assert.Equal(t, expectedCount, actualCount, msg)
        }
}`, bookName, prettyChapterName, expectedVerseCount, i, currentChapter)
			fmt.Println(verseTest)
			currentChapter++
		}
		i++
	}
}
