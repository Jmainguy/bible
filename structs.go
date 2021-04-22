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
