package model

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	Id                  uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserId              uuid.UUID  `json:"user_id" binding:"required"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	LastNotificationRun *time.Time `json:"last_notification_run"`
}

type SubsWithUserEmail struct {
	Subscription
	Email string
}
