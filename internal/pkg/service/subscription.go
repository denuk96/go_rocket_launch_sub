package service

import (
	"go_rocket_launch_sub/internal/pkg/repository"
)

type SubscriptionService struct {
	repo repository.Subscription
}

func (s *SubscriptionService) Create(userId string) (string, error) {
	subId, error := s.repo.Create(userId)
	if error != nil {
		return "", error
	}

	return subId, nil
}

func (s *SubscriptionService) Destroy(userId, subId string) error {
	error := s.repo.Destroy(userId, subId)

	return error
}

func NewSubscriptionService(repo repository.Subscription) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}
