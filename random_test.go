package main

import (
	"database/sql"
	"os"
	"testing"
)

func Test_randomChapter(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name: "Jude one chapter",
			args: args{
				config: Config{
					BookInfoList: []BookInfo{
						{
							Title:    "Jude",
							Chapters: 1,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := randomChapter(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("randomChapter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_randomPassage(t *testing.T) {
	os.Remove("/tmp/bible.db")

	baddb, _ := sql.Open("sqlite", "/tmp/bible.db")
	gooddb, _ := sql.Open("sqlite", "database/bible.db")

	type args struct {
		config Config
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name: "Random Bad DB",
			args: args{
				config: Config{
					BookInfoList: []BookInfo{
						{
							Title:    "Jude",
							Chapters: 1,
						},
					},
					DB: baddb,
				},
			},
			wantErr: true,
		},
		{
			name: "Bad query",
			args: args{
				config: Config{
					BookInfoList: []BookInfo{
						{
							Title:    "Jude",
							Chapters: 1,
						},
					},
					DB:          gooddb,
					Translation: "t_jon",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := randomPassage(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("randomPassage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	os.Remove("/tmp/bible.db")
}
