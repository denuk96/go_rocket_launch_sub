package config

import (
	"github.com/joho/godotenv"
	"log"
)

func InitEnvs() {
	err := godotenv.Load() // load .env file by default

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Output(2, "File .env successfully loaded\n")
}
