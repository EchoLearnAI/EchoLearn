package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configuration values
type Config struct {
	ServerPort string
	DBName     string
}

var AppConfig *Config

// LoadConfig loads configuration from .env file or environment variables
func LoadConfig() {
	// Attempt to load .env file, but don't fail if it doesn't exist
	// This allows for environment variables to take precedence or be the sole source
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables or defaults.")
	}

	AppConfig = &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBName:     getEnv("DB_NAME", "echolearn.db"),
	}

	log.Printf("Configuration loaded: Port=%s, DBName=%s", AppConfig.ServerPort, AppConfig.DBName)
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
} 