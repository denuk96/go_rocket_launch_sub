package service

import (
	"go_rocket_launch_sub/internal/pkg/model"
	"time"
)

type LaunchesService struct {
	Launches []model.Launch
}

func NewLaunchService(launches []model.Launch) *LaunchesService {
	return &LaunchesService{Launches: launches}
}

func (lc *LaunchesService) getEarliestAndLatestLaunchTime() (time.Time, time.Time) {
	if len(lc.Launches) == 0 {
		return time.Time{}, time.Time{}
	}

	var earliestLaunch, latestLaunch *model.Launch
	earliestLaunch = &lc.Launches[0]
	latestLaunch = &lc.Launches[0]

	for i := 1; i < len(lc.Launches); i++ {
		launch := lc.Launches[i]
		if launch.LaunchTime.Before(earliestLaunch.LaunchTime.Time) {
			earliestLaunch = &lc.Launches[i]
		}
		if launch.LaunchTime.After(latestLaunch.LaunchTime.Time) {
			latestLaunch = &lc.Launches[i]
		}
	}

	return earliestLaunch.LaunchTime.Time, latestLaunch.LaunchTime.Time
}
