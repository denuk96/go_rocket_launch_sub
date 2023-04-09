package service

import (
	"github.com/google/uuid"
	"go_rocket_launch_sub/internal/pkg/model"
	"go_rocket_launch_sub/internal/pkg/repository"
)

type Authorisation interface {
	CreateUser(user model.User) (uuid.UUID, error)
}

type Subscription interface {
}

type Service struct {
	Authorisation
	Subscription
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repository.Authorisation),
	}
}
