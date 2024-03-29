package repository

import (
	"go_rocket_launch_sub/internal/pkg/model"
	"gorm.io/gorm"
	"time"
)

type Authorisation interface {
	CreateUser(user model.User) (string, error)
	FindUserByEmail(email string) (model.User, error)
	FindUserById(uuid string) (model.User, error)
}

type Subscription interface {
	Create(userId string) (string, error)
	Update(subscriptionId string, params map[string]interface{}) error
	Destroy(userId, subId string) error
	FindByUser(userId string) (model.Subscription, error)
	ListByUser(userId string) ([]model.Subscription, error)
	UnNotifiedAfter(time time.Time) ([]model.SubsWithUserEmail, error)
}

type Repository struct {
	Authorisation
	Subscription
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorisation: NewUserPsql(db),
		Subscription:  NewSubscriptionPsql(db),
	}
}
