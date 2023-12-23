package config

import (
	"net/smtp"
	"os"
)

type SmtpCreds struct {
	Auth   smtp.Auth
	Server string
	Port   string
	Sender string
}

func InitSmtp() (smtpCreds SmtpCreds) {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)

	return SmtpCreds{
		Auth:   auth,
		Server: os.Getenv("SMTP_HOST"),
		Port:   os.Getenv("SMTP_PORT"),
		Sender: os.Getenv("SMTP_USER"),
	}
}
