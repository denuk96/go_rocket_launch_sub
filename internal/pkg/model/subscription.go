package model

import "time"

type Subscription struct {
	ID        string    `json:"id" gorm:"primary_key"`
	UserId    string    `json:"user_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
