package repository

import (
	"github.com/google/uuid"
	"go_rocket_launch_sub/internal/pkg/model"
	"gorm.io/gorm"
	"time"
)

type SubscriptionPsql struct {
	db *gorm.DB
}

func NewSubscriptionPsql(db *gorm.DB) *SubscriptionPsql {
	return &SubscriptionPsql{db: db}
}

func (psql *SubscriptionPsql) FindByUser(userId string) (model.Subscription, error) {
	userUUID, _ := uuid.Parse(userId)
	var subscription = model.Subscription{}
	result := psql.db.First(&subscription, "user_id = ?", userUUID)

	return subscription, result.Error
}

func (psql *SubscriptionPsql) Create(userId string) (string, error) {
	userUUID, _ := uuid.Parse(userId)
	subscription := model.Subscription{UserId: userUUID}

	result := psql.db.Create(&subscription)

	return subscription.Id.String(), result.Error
}

func (psql *SubscriptionPsql) ListByUser(userId string) ([]model.Subscription, error) {
	var subscriptions []model.Subscription

	result := psql.db.Where("user_id = ?", userId).Find(&subscriptions)

	return subscriptions, result.Error
}

func (psql *SubscriptionPsql) Destroy(userId string, subId string) error {
	userUUID, _ := uuid.Parse(subId)
	id, _ := uuid.Parse(subId)

	subscription := model.Subscription{Id: id, UserId: userUUID}
	result := psql.db.Delete(&subscription)

	return result.Error
}

func (psql *SubscriptionPsql) UnNotifiedWithin(minutes int) ([]model.SubsWithUserEmail, error) {
	var subscriptions []model.SubsWithUserEmail

	timeBoundary := time.Now().Add(-time.Duration(minutes) * time.Minute)

	result := psql.db.Table("subscriptions").
		Select("subscriptions.*, users.email").
		Joins("left join users on users.id = subscriptions.user_id").
		Where("subscriptions.last_notification_run IS NULL OR subscriptions.last_notification_run <= ?", timeBoundary).
		Scan(&subscriptions)

	if result.Error != nil {
		return nil, result.Error
	}

	return subscriptions, nil
}
