package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// config func to get env value from key ...
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
