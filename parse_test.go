package main

import (
	"database/sql"
	"testing"
)

func Test_getVerseID(t *testing.T) {
	db, err := sql.Open("sqlite", "database/bible.db")
	check(err)
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
				version:      "t_asdasd",
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
	db, err := sql.Open("sqlite", "database/bible.db")
	check(err)

	type args struct {
		verseString string
		translation string
		db          *sql.DB
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
				verseString: "Jon",
				translation: "t_kjv",
				db:          db,
			},
			wantQuery: "",
			wantErr:   true,
		},
		{
			name: "Bad Characters",
			args: args{
				verseString: "John asd-aa",
				translation: "t_kjv",
				db:          db,
			},
			wantQuery: "",
			wantErr:   true,
		},
		{
			name: "Book and Chapter - blank translation",
			args: args{
				verseString: "John 3",
				translation: "",
				db:          db,
			},
			wantQuery: "",
			wantErr:   true,
		},
		{
			name: "Book and Chapter - kjv",
			args: args{
				verseString: "John 3",
				translation: "t_kjv",
				db:          db,
			},
			wantQuery: "SELECT * FROM t_kjv WHERE b is 43 and c is 3;",
			wantErr:   false,
		},
		{
			name: "Book, Chapter, and verse - kjv",
			args: args{
				verseString: "John 3:16",
				translation: "t_kjv",
				db:          db,
			},
			wantQuery: "SELECT * FROM t_kjv WHERE id is 43003016;",
			wantErr:   false,
		},
		{
			name: "multiple verses",
			args: args{
				verseString: "John 3:16 - 4",
				translation: "t_kjv",
				db:          db,
			},
			wantQuery: "SELECT * FROM t_kjv WHERE id BETWEEN 43003016 AND 43004054;",
			wantErr:   false,
		},
		{
			name: "multiple verses no hyphen",
			args: args{
				verseString: "John 3:16  4",
				translation: "t_kjv",
				db:          db,
			},
			wantQuery: "",
			wantErr:   true,
		},
		{
			name: "multiple verses no hyphen extra chars",
			args: args{
				verseString: "John 3:16 5 4",
				translation: "t_kjv",
				db:          db,
			},
			wantQuery: "",
			wantErr:   true,
		},
		{
			name: "multiple verses no hyphen extra chars bad verses",
			args: args{
				verseString: "John 3:16 - 3:980",
				translation: "t_kjv",
				db:          db,
			},
			wantQuery: "SELECT * FROM t_kjv WHERE id BETWEEN 43003016 AND 43003980;",
			wantErr:   false,
		},
		{
			name: "book only",
			args: args{
				verseString: "John",
				translation: "t_kjv",
				db:          db,
			},
			wantQuery: "SELECT * FROM t_kjv WHERE b is 43;",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, err := parseVerseString(tt.args.verseString, tt.args.translation, tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseVerseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotQuery != tt.wantQuery {
				t.Errorf("parseVerseString() = %v, want %v", gotQuery, tt.wantQuery)
			}
		})
	}
}
