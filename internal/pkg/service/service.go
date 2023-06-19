package service

import (
	"go_rocket_launch_sub/internal/pkg/model"
	"go_rocket_launch_sub/internal/pkg/repository"
)

type Authorisation interface {
	SignUp(user model.User) (string, error)
	SignIn(email, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Subscription interface {
	Create(userId string) (string, error)
	Destroy(userId, subId string) error
	AllByUser(userId string) ([]model.Subscription, error)
}

type Service struct {
	Authorisation
	Subscription
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repository.Authorisation),
		Subscription:  NewSubscriptionService(repository.Subscription),
	}
}
