package main

import (
	"os"
	"testing"
)

func Test_downloadDB(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "DB Already exists",
			args: args{
				file: "database/bible.db",
			},
			wantErr: false,
		},
		{
			name: "Download db",
			args: args{
				file: "/tmp/bible.db",
			},
			wantErr: false,
		},
		{
			name: "Directory is invalid",
			args: args{
				file: "/dev/null/asd/bible.db",
			},
			wantErr: true,
		},
		{
			name: "Path is not directory, cannot remove",
			args: args{
				file: "/proc/mdstat/bible.db",
			},
			wantErr: true,
		},
		{
			name: "Path is not directory, can remove",
			args: args{
				file: "/tmp/bible.db/bible.db",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := downloadDB(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("downloadDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	// Clean up
	if err := os.Remove("/tmp/bible.db/bible.db"); err != nil {
		t.Fatalf("failed to remove /tmp/bible.db/bible.db: %v", err)
	}
	if err := os.Remove("/tmp/bible.db"); err != nil {
		t.Fatalf("failed to remove /tmp/bible.db: %v", err)
	}
}
