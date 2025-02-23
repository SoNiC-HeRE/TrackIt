package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file if not in production
func LoadEnv() {
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not found")
		} else {
			log.Println(".env file loaded successfully")
		}
	}
}
