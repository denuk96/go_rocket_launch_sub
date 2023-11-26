package model

import (
	"strings"
	"time"
)

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
