package models

import "time"

type Book struct {
	ID                 int       `gorm:"primaryKey" json:"id"`
	Title              string    `json:"title"`
	Author             string    `json:"author"`
	Genre              string    `json:"genre"`
	PublicationYear    int       `json:"publication_year"`
	AvailabilityStatus string    `json:"availability_status"`
	CreatedAt          time.Time `json:"created_at"`
}
