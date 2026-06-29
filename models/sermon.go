package models

type Sermon struct {
	ID         int
	Title      string
	BibleVerse string
	References string
	Content    string
	Category   string
	Date       string
	CreatedBy  string
	CreatedAt  string
}