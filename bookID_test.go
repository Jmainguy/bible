package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func Test_mapIDToBook(t *testing.T) {
	db, err := sql.Open("sqlite", "database/bible.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "There are 66 books in bible",
			args: args{
				db: db,
			},
			want:    66,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapIDToBook(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapIDToBook()  error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("mapIDToBook() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_getBooks(t *testing.T) {
	db, err := sql.Open("sqlite", "database/bible.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "There are 66 books in bible",
			args: args{
				db: db,
			},
			want:    66,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBooks(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBooks()  error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("getBooks() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_listBooks(t *testing.T) {
	db, err := sql.Open("sqlite", "database/bible.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "There are 66 books in bible",
			args: args{
				db: db,
			},
			want:    66,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := listBooks(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("listBooks()  error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("listBooks() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
