package service

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

const apiUrl = "https://fdo.rocketlaunch.live/json/launches/next/5"

type Launch struct {
	ID                 int        `json:"id"`
	Name               string     `json:"name"`
	MissionDescription string     `json:"mission_description"`
	LaunchDescription  string     `json:"launch_description"`
	LaunchTime         customTime `json:"t0"`
	Vehicle            Vehicle
	Pad                Pad
}

type Vehicle struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Pad struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location Location
}

type Location struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	State     string `json:"state"`
	StateName string `json:"statename"`
	Country   string `json:"country"`
}

type ApiResponse struct {
	Count  int      `json:"count"`
	Result []Launch `json:"result"`
}

func GetNextLaunches() ([]Launch, error) {
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"url":   apiUrl,
		}).Error("Failed to fetch data")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"status_code": resp.StatusCode,
			"url":         apiUrl,
		}).Error("Non-200 response received")
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Error("Failed to read response body")
		return nil, err
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.WithError(err).Error("Failed to unmarshal response")
		return nil, err
	}

	return apiResponse.Result, nil
}

// handle custom time in response
type customTime struct {
	time.Time
}

func (ct *customTime) UnmarshalJSON(b []byte) error {
	strInput := string(b)
	strInput = strings.Trim(strInput, "\"")
	if strInput == "null" || strInput == "" {
		return nil
	}
	newTime, err := time.Parse("2006-01-02T15:04Z", strInput)
	if err != nil {
		return err
	}
	ct.Time = newTime
	return nil
}
