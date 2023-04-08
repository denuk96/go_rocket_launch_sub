package service

import (
	"go_rocket_launch_sub/internal/pkg/repository"
)

type Authorisation interface {
}

type Subscription interface {
}

type Service struct {
	Authorisation
	Subscription
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
