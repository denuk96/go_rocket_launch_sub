package service

import (
	log "github.com/sirupsen/logrus"
	"go_rocket_launch_sub/internal/pkg/model"
	"go_rocket_launch_sub/internal/pkg/repository"
)

type LaunchNotificationService struct {
	repo repository.Subscription
}

func (n *LaunchNotificationService) NotifyAll() {
	launchApiClient := NewLaunchApiService()

	err := launchApiClient.FetchUpcomingLaunches()
	if err != nil {
		log.WithError(err).Error("Failed to get launches")
		return
	}
	launchesService := NewLaunchService(launchApiClient.Launches)
	earliestLaunchTime, latestLaunchTime := launchesService.getEarliestAndLatestLaunchTime()

	subscriptions, err := n.repo.UnNotifiedAfter(earliestLaunchTime)
	if err != nil {
		log.WithError(err).Error("Failed to get fetch subscriptions")
		return
	}

	for _, subscription := range subscriptions {
		err := n.SendNotification(subscription)
		if err != nil {
			return
		}

		updateParams := map[string]interface{}{
			"last_notification_run": latestLaunchTime,
		}
		err = n.repo.Update(subscription.Id.String(), updateParams)
		if err != nil {
			log.Error(err.Error())
		}
	}
}

func (n *LaunchNotificationService) SendNotification(subscription model.SubsWithUserEmail) error {
	//TODO: write SMTP service
	log.Infof(subscription.Email)

	return nil
}

func NewNotificationService(repo repository.Subscription) *LaunchNotificationService {
	return &LaunchNotificationService{repo: repo}
}
