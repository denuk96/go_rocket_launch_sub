package service

import (
	"errors"
	"go_rocket_launch_sub/internal/pkg/model"
	"go_rocket_launch_sub/internal/pkg/repository"

	"gorm.io/gorm"
)

type SubscriptionService struct {
	repo repository.Subscription
}

func (s *SubscriptionService) Create(userId string) (string, error) {
	subscription, err := s.repo.FindByUser(userId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			subId, err := s.repo.Create(userId)
			if err != nil {
				return "", err
			}

			return subId, err
		}

		return "", err
	}

	return subscription.Id.String(), nil
}

func (s *SubscriptionService) Destroy(userId, subId string) error {
	err := s.repo.Destroy(userId, subId)

	return err
}

func (s *SubscriptionService) AllByUser(userId string) ([]model.Subscription, error) {
	return s.repo.ListByUser(userId)
}

func NewSubscriptionService(repo repository.Subscription) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}
