package service

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"go_rocket_launch_sub/internal/config"
	"go_rocket_launch_sub/internal/pkg/model"
	"html/template"
	"net/smtp"
	"path/filepath"
)

type EmailService struct {
	smtpCreds config.SmtpCreds
}

func (e *EmailService) SendLaunchNotification(recipient string, launches []model.Launch) error {
	log.Infof("Number of launches: %d", len(launches))
	emailBody, err := parseTemplate("launch_notification.html", struct {
		Launches []model.Launch
	}{Launches: launches})
	if err != nil {
		log.Error(err)
	}

	// Set up the email headers and body
	msg := []byte("From: " + e.smtpCreds.Server + "\r\n" +
		"To: " + recipient + "\r\n" +
		"Subject: Launch Notification\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
		emailBody)

	// Send the email
	err = smtp.SendMail(e.smtpCreds.Server+":"+e.smtpCreds.Port, e.smtpCreds.Auth, e.smtpCreds.Server, []string{recipient}, msg)

	return err
}

func parseTemplate(templateName string, data interface{}) (string, error) {
	templatePath := filepath.Join("internal/pkg/templates", templateName)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func NewEmailService(smtpCreds config.SmtpCreds) *EmailService {
	return &EmailService{smtpCreds: smtpCreds}
}
