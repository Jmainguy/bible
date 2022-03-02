package main

import (
	"testing"
)

func Test_checkNumeric(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Numeric: Has has letters only",
			args: args{
				s: "Jon",
			},
			want: false,
		},

		{
			name: "Numeric: Has has letters and numbers",
			args: args{
				s: "Jon 1",
			},
			want: false,
		},
		{
			name: "Numeric: Has has numbers and spaces",
			args: args{
				s: "1 11",
			},
			want: false,
		},
		{
			name: "Numeric: Just numbers",
			args: args{
				s: "1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkNumeric(tt.args.s); got != tt.want {
				t.Errorf("checkNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkNumericPlusColonDash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "NumericPlus: Has has letters only",
			args: args{
				s: "Jon",
			},
			want: false,
		},

		{
			name: "NumericPlus: Has has letters and numbers",
			args: args{
				s: "Jon 1",
			},
			want: false,
		},
		{
			name: "NumericPlus: Has has numbers and spaces",
			args: args{
				s: "1 11",
			},
			want: false,
		},
		{
			name: "NumericPlus: Just numbers",
			args: args{
				s: "1",
			},
			want: true,
		},
		{
			name: "NumericPlus: numbers colon and -",
			args: args{
				s: "1-:",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkNumericPlusColonDash(tt.args.s); got != tt.want {
				t.Errorf("checkNumericPlusColonDash() = %v, want %v", got, tt.want)
			}
		})
	}
}
