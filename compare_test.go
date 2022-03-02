package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func Test_compare(t *testing.T) {

	gooddb, _ := sql.Open("sqlite", "database/bible.db")

	translationsAvailable, err := getTranslations(gooddb)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	type args struct {
		config Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Compare bad translation",
			args: args{
				config: Config{
					BookInfoList: []BookInfo{
						{
							Title:    "Jude",
							Chapters: 1,
						},
					},
					DB:                    gooddb,
					Compare:               "t_jon t_kjv",
					TranslationsAvailable: translationsAvailable,
					Passage:               "John 3:16",
				},
			},
			wantErr: true,
		},
		{
			name: "Compare bad translation no passage",
			args: args{
				config: Config{
					BookInfoList: []BookInfo{
						{
							Title:    "Jude",
							Chapters: 1,
						},
					},
					DB:                    gooddb,
					Compare:               "t_jon t_kjv",
					TranslationsAvailable: translationsAvailable,
				},
			},
			wantErr: true,
		},
		{
			name: "Compare all translations",
			args: args{
				config: Config{
					BookInfoList: []BookInfo{
						{
							Title:    "Jude",
							Chapters: 1,
						},
					},
					DB:                    gooddb,
					Compare:               "all",
					TranslationsAvailable: translationsAvailable,
					Passage:               "John 3:16",
				},
			},
			wantErr: false,
		},
		{
			name: "Compare all translations no passage",
			args: args{
				config: Config{
					BookInfoList: []BookInfo{
						{
							Title:    "Jude",
							Chapters: 1,
						},
					},
					DB:                    gooddb,
					Compare:               "all",
					TranslationsAvailable: translationsAvailable,
				},
			},
			wantErr: true,
		},

		{
			name: "Compare leb and kjv",
			args: args{
				config: Config{
					BookInfoList: []BookInfo{
						{
							Title:    "Jude",
							Chapters: 1,
						},
					},
					DB:                    gooddb,
					Compare:               "t_leb t_kjv",
					TranslationsAvailable: translationsAvailable,
					Passage:               "John 3:16",
				},
			},
			wantErr: false,
		},
		{
			name: "Compare leb and kjv no passage",
			args: args{
				config: Config{
					BookInfoList: []BookInfo{
						{
							Title:    "Jude",
							Chapters: 1,
						},
					},
					DB:                    gooddb,
					Compare:               "t_leb t_kjv",
					TranslationsAvailable: translationsAvailable,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := compare(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("compare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
