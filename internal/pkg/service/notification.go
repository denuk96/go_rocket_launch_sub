package service

import (
	log "github.com/sirupsen/logrus"
	"go_rocket_launch_sub/internal/pkg/model"
	"go_rocket_launch_sub/internal/pkg/repository"
)

type NotificationService struct {
	repo repository.Subscription
}

func (n *NotificationService) NotifyAllWithin(minutes int) {
	subscriptions, _ := n.repo.UnNotifiedWithin(minutes)

	for _, subscription := range subscriptions {
		err := n.SendNotification(subscription)
		if err != nil {
			log.Error(err.Error())
		}
	}
}

func (n *NotificationService) SendNotification(subscription model.SubsWithUserEmail) error {
	log.Infof(subscription.Email)

	return nil
}

func NewNotificationService(repo repository.Subscription) *NotificationService {
	return &NotificationService{repo: repo}
}
