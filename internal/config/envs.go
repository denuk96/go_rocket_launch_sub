package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func InitEnvs() {
	curDir, err := os.Getwd()
	if err != nil {
		log.Info(err)
	}
	loadErr := godotenv.Load(curDir + "/.env")
	if loadErr != nil {
		log.Fatal("can't load env file from current directory: " + curDir)
	}

	log.Info("File .env successfully loaded\n")
}
