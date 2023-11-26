package service

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"go_rocket_launch_sub/internal/pkg/model"
	"io"
	"net/http"
)

const apiUrl = "https://fdo.rocketlaunch.live/json/launches/next/5"

type apiResponse struct {
	Count  int            `json:"count"`
	Result []model.Launch `json:"result"`
}

type LaunchApiService struct {
	Launches []model.Launch
}

func (lc *LaunchApiService) FetchUpcomingLaunches() error {
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"url":   apiUrl,
		}).Error("Failed to fetch data")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"status_code": resp.StatusCode,
			"url":         apiUrl,
		}).Error("Non-200 response received")
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Error("Failed to read response body")
		return err
	}

	var apiResponse apiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.WithError(err).Error("Failed to unmarshal response")
		return err
	}

	lc.Launches = apiResponse.Result

	return nil
}

func NewLaunchApiService() *LaunchApiService {
	return &LaunchApiService{Launches: []model.Launch{}}
}
