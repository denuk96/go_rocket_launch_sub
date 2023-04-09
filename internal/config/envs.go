package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func InitEnvs() {
	err := godotenv.Load() // load .env file by default

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Info("File .env successfully loaded\n")
}
