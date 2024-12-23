package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	Database    DatabaseConfig
	Env         string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var AppConfig Config

// LoadConfig loads environment variables and sets the AppConfig values
func LoadConfig() {
	// Load .env file for development environments
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set values for AppConfig from environment variables
	AppConfig = Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("ENV", "development"),
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "user"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "defaultdb"),
		},
	}
}

// getEnv is a helper function to get environment variables with default fallback
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
