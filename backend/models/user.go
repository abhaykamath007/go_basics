package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
