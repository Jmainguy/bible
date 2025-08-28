package main

import (
	"database/sql"
	"os"
	"testing"
)

func Test_printPassage(t *testing.T) {
	if err := os.Remove("/tmp/bible.db"); err != nil {
		t.Fatalf("failed to remove /tmp/bible.db: %v", err)
	}

	baddb, _ := sql.Open("sqlite", "/tmp/bible.db")
	gooddb, _ := sql.Open("sqlite", "database/bible.db")

	type args struct {
		config Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Print bad db",
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
			name: "Print bad translation",
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
			if err := printPassage(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("printPassage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if err := os.Remove("/tmp/bible.db"); err != nil {
		t.Fatalf("failed to remove /tmp/bible.db: %v", err)
	}

}
