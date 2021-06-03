package main

// Verse : Verse structure of scrollmapper db
type Verse struct {
	ID      int
	Book    int
	Chapter int
	Verse   int
	Text    string
}

// KeyEnglish : Limited structure of scrollmapper key_english table
type KeyEnglish struct {
	ID   int
	Book string
}

// BookInfo : Information on a book
type BookInfo struct {
	ID        int
	Title     string
	Testament string
	Chapters  int
}

// Translation : Information on a translation
type Translation struct {
	ID            int
	Table         string
	Abbreviation  string
	Language      string
	Version       string
	InfoText      string
	InfoURL       string
	Publisher     string
	Copyright     string
	CopyrightInfo string
}
