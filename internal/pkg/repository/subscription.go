package repository

import (
	"go_rocket_launch_sub/internal/pkg/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionPsql struct {
	db *gorm.DB
}

func NewSubscriptionPsql(db *gorm.DB) *SubscriptionPsql {
	return &SubscriptionPsql{db: db}
}

func (psql *SubscriptionPsql) FindByUser(userId string) (model.Subscription, error) {
	var subscription = model.Subscription{}
	result := psql.db.First(&subscription, "user_id = ?", userId)

	return subscription, result.Error
}

func (psql *SubscriptionPsql) Create(userId string) (string, error) {
	subscription := model.Subscription{UserId: userId}

	result := psql.db.Create(&subscription)

	return subscription.Id.String(), result.Error
}

func (psql *SubscriptionPsql) ListByUser(userId string) ([]model.Subscription, error) {
	var subscriptions []model.Subscription

	result := psql.db.Where("user_id = ?", userId).Find(&subscriptions)

	return subscriptions, result.Error
}

func (psql *SubscriptionPsql) Destroy(userId string, subId string) error {
	id, _ := uuid.Parse(subId)

	subscription := model.Subscription{Id: id, UserId: userId}
	result := psql.db.Delete(&subscription)

	return result.Error
}
