package repository

import (
	"go_rocket_launch_sub/internal/pkg/model"
	"gorm.io/gorm"
)

type SubscriptionPsql struct {
	db *gorm.DB
}

func NewSubscriptionPsql(db *gorm.DB) *SubscriptionPsql {
	return &SubscriptionPsql{db: db}
}

func (psql *SubscriptionPsql) Create(userId string) (string, error) {
	subscription := model.Subscription{UserId: userId}

	result := psql.db.Create(&subscription)

	return subscription.ID, result.Error
}

func (psql *SubscriptionPsql) Destroy(userId, subId string) error {
	subscription := model.Subscription{ID: subId, UserId: userId}
	result := psql.db.Delete(&subscription)

	return result.Error
}
