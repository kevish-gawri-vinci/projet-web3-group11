package entity

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	IsAdmin   bool      `json:"isadmin"`
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time
}
