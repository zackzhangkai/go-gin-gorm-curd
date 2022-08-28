package initialenviroment

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
