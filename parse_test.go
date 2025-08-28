package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func Test_getVerseID(t *testing.T) {
	db, err := sql.Open("sqlite", "database/bible.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	type args struct {
		book         int
		chapterVerse string
		db           *sql.DB
		minmax       string
		version      string
	}
	tests := []struct {
		name        string
		args        args
		wantVerseID int
		wantErr     bool
	}{
		{
			name: "3:16",
			args: args{
				book:         1,
				chapterVerse: "3:16",
				db:           db,
				minmax:       "",
				version:      "",
			},
			wantVerseID: 1003016,
			wantErr:     false,
		},
		{
			name: "3",
			args: args{
				book:         1,
				chapterVerse: "3",
				db:           db,
				minmax:       "",
				version:      "",
			},
			wantVerseID: 0,
			wantErr:     true,
		},
		{
			name: "3max",
			args: args{
				book:         1,
				chapterVerse: "3",
				db:           db,
				minmax:       "max",
				version:      "",
			},
			wantVerseID: 0,
			wantErr:     true,
		},
		{
			name: "3min",
			args: args{
				book:         1,
				chapterVerse: "3",
				db:           db,
				minmax:       "min",
				version:      "",
			},
			wantVerseID: 1003001,
			wantErr:     false,
		},
		{
			name: "3maxBadVersion",
			args: args{
				book:         1,
				chapterVerse: "3",
				db:           db,
				minmax:       "max",
				version:      "t_bad",
			},
			wantVerseID: 0,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVerseID, err := getVerseID(tt.args.book, tt.args.chapterVerse, tt.args.db, tt.args.minmax, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("getVerseID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotVerseID != tt.wantVerseID {
				t.Errorf("getVerseID() = %v, want %v", gotVerseID, tt.wantVerseID)
			}
		})
	}
}

func Test_parseVerseString(t *testing.T) {
	if err := os.Remove("/tmp/bible.db"); err != nil {
		t.Fatalf("failed to remove /tmp/bible.db: %v", err)
	}

	db, _ := sql.Open("sqlite", "database/bible.db")
	baddb, _ := sql.Open("sqlite", "/tmp/bible.db")

	translationsAvailable, err := getTranslations(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	type args struct {
		config      Config
		translation string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery string
		wantErr   bool
	}{
		{
			name: "Bad Book",
			args: args{
				config: Config{
					Passage:               "Jon",
					DB:                    db,
					TranslationsAvailable: translationsAvailable,
				},
				translation: "t_kjv",
			},
			wantQuery: "",
			wantErr:   true,
		},
		{
			name: "Bad Characters",
			args: args{
				config: Config{
					Passage:               "John asd-aa",
					DB:                    db,
					TranslationsAvailable: translationsAvailable,
				},
				translation: "t_kjv",
			},
			wantQuery: "",
			wantErr:   true,
		},
		{
			name: "Book and Chapter - blank translation",
			args: args{
				config: Config{
					Passage:               "John 3",
					DB:                    db,
					TranslationsAvailable: translationsAvailable,
				},
			},
			wantQuery: "",
			wantErr:   true,
		},
		{
			name: "Book and Chapter - kjv",
			args: args{
				config: Config{
					Passage:               "John 3",
					DB:                    db,
					TranslationsAvailable: translationsAvailable,
				},
				translation: "t_kjv",
			},
			wantQuery: "SELECT * FROM t_kjv WHERE b is 43 and c is 3;",
			wantErr:   false,
		},
		{
			name: "Book, Chapter, and verse - kjv",
			args: args{
				config: Config{
					Passage:               "John 3:16",
					DB:                    db,
					TranslationsAvailable: translationsAvailable,
				},
				translation: "t_kjv",
			},
			wantQuery: "SELECT * FROM t_kjv WHERE id is 43003016;",
			wantErr:   false,
		},
		{
			name: "multiple verses",
			args: args{
				config: Config{
					Passage:               "John 3:16 - 4",
					DB:                    db,
					TranslationsAvailable: translationsAvailable,
				},
				translation: "t_kjv",
			},
			wantQuery: "SELECT * FROM t_kjv WHERE id BETWEEN 43003016 AND 43004054;",
			wantErr:   false,
		},
		{
			name: "multiple verses no hyphen",
			args: args{
				config: Config{
					Passage:               "John 3:16  4",
					DB:                    db,
					TranslationsAvailable: translationsAvailable,
				},
				translation: "t_kjv",
			},
			wantQuery: "",
			wantErr:   true,
		},
		{
			name: "multiple verses no hyphen extra chars",
			args: args{
				config: Config{
					Passage:               "John 3:16 5 4",
					DB:                    db,
					TranslationsAvailable: translationsAvailable,
				},
				translation: "t_kjv",
			},
			wantQuery: "",
			wantErr:   true,
		},
		{
			name: "multiple verses no hyphen extra chars bad verses",
			args: args{
				config: Config{
					Passage:               "John 3:16 - 3:980",
					DB:                    db,
					TranslationsAvailable: translationsAvailable,
				},
				translation: "t_kjv",
			},
			wantQuery: "SELECT * FROM t_kjv WHERE id BETWEEN 43003016 AND 43003980;",
			wantErr:   false,
		},
		{
			name: "book only",
			args: args{
				config: Config{
					Passage:               "John",
					DB:                    db,
					TranslationsAvailable: translationsAvailable,
				},
				translation: "t_kjv",
			},
			wantQuery: "SELECT * FROM t_kjv WHERE b is 43;",
			wantErr:   false,
		},
		{
			name: "Bad chars in verse",
			args: args{
				config: Config{
					Passage:               "John 3:16..",
					DB:                    db,
					TranslationsAvailable: translationsAvailable,
				},
				translation: "t_kjv",
			},
			wantQuery: "",
			wantErr:   true,
		},
		{
			name: "Bad database",
			args: args{
				config: Config{
					Passage:               "John 3:16..",
					DB:                    baddb,
					TranslationsAvailable: translationsAvailable,
				},
				translation: "t_kjv",
			},
			wantQuery: "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, err := parseVerseString(tt.args.config, tt.args.translation)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseVerseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotQuery != tt.wantQuery {
				t.Errorf("parseVerseString() = %v, want %v", gotQuery, tt.wantQuery)
			}
		})
	}
	if err := os.Remove("/tmp/bible.db"); err != nil {
		t.Fatalf("failed to remove /tmp/bible.db: %v", err)
	}

}
