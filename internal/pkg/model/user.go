package model

import "time"

type User struct {
	Id        string    `json:"-" db:"id"`
	Name      string    `json:"name" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
