package model

import "time"

// Book represents a book from a bookstore system.
type Book struct {
	ID           int64 `bun:"id,pk,autoincrement"`
	ISBN         string
	Title        string
	Author       string
	Genre        string
	Quantity     int
	PublicizedAt time.Time
}

// PricePolicy represents the price policies for a book, since the same book can have various prices in different periods.
type PricePolicy struct {
	ID        int64
	BookID    int64
	Price     float64
	StartDate time.Time
	EndDate   time.Time
}

func NewBooks(quantity int) []*Book {
	books := make([]*Book, quantity)
	for i := 0; i < quantity; i++ {
		books[i] = NewBook()
	}
	return books
}

func NewBook() *Book {
	return &Book{
		ISBN:         "978-3-16-148410-1",
		Title:        "Learning Go: An Idiomatic Approach to Real-World Go Programming",
		Author:       "Jon Bodner",
		Genre:        "Programming",
		Quantity:     20,
		PublicizedAt: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
}

func GetMaxID(books []*Book) int64 {
	if len(books) == 0 {
		return 0
	}
	maxID := books[0].ID
	for _, book := range books {
		if book.ID > maxID {
			maxID = book.ID
		}
	}
	return maxID
}
