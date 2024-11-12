package models

import "time"

type Transaction struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	UserID       int       `json:"user_id"`
	BookID       int       `json:"book_id"`
	BorrowedDate time.Time `json:"borrowed_date"`
	DueDate      time.Time `json:"date"`
	Returned     bool      `json:"returned"`
	// UpdatedAt    time.Time `json:"updated_at"`
}

type BorrowRequest struct {
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}
