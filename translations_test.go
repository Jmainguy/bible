package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func Test_getTranslations(t *testing.T) {
	db, err := sql.Open("sqlite", "database/bible.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	emptyDB, err := sql.Open("sqlite", "/tmp/empty.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no error",
			args: args{
				db: db,
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				db: emptyDB,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getTranslations(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTranslations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
