package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id           uuid.UUID `json:"-" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name         string    `json:"name" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	PasswordHash string    `json:"password" binding:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
