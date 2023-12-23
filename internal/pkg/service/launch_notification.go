package service

import (
	log "github.com/sirupsen/logrus"
	"go_rocket_launch_sub/internal/pkg/model"
	"go_rocket_launch_sub/internal/pkg/repository"
)

type LaunchNotificationService struct {
	repo repository.Subscription
	emailService EmailService
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
		unNotifiedLaunches := launchesService.selectUnNotifiedLaunches(*subscription.LastNotificationRun)

		if len(unNotifiedLaunches) == 0 {
			return
		}

		err := n.SendNotification(subscription, unNotifiedLaunches)
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

func (n *LaunchNotificationService) SendNotification(subscription model.SubsWithUserEmail, launches []model.Launch) error {
	log.Printf("sending email to %s", subscription.Email)

	return n.emailService.SendLaunchNotification(subscription.Email, launches)
}

func NewNotificationService(repo repository.Subscription, emailService EmailService) *LaunchNotificationService {
	return &LaunchNotificationService{repo: repo, emailService: emailService}
}
