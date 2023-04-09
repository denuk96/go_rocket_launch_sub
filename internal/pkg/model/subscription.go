package model

import "time"

type Subscription struct {
	ID        string    `json:"-" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserId    string    `json:"user_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
