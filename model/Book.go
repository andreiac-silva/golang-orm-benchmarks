package model

import "time"

// Book represents a book from a bookstore system.
type Book struct {
	ID           int64
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
