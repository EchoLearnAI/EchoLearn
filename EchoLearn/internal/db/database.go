package db

import (
	"fmt"
	"log"
	"os"

	"echolearn/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

func InitDB() {
	var err error

	// Load environment variables from .env file
	godotenv.Load()

	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "sqlite"
	}

	switch dbType {
	case "sqlite":
		DB, err = connectSQLite()
	case "postgres":
		DB, err = connectPostgres()
	default:
		log.Fatalf("Unsupported database type: %s", dbType)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate models
	migrateModels()
}

func connectSQLite() (*gorm.DB, error) {
	dbPath := os.Getenv("SQLITE_PATH")
	if dbPath == "" {
		dbPath = "data/echolearn.db"
	}

	// Ensure directory exists
	dir := "data"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, fmt.Errorf("failed to create data directory: %w", err)
		}
	}

	return gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
}

func connectPostgres() (*gorm.DB, error) {
	config := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == "" {
		config.Port = "5432"
	}
	if config.SSLMode == "" {
		config.SSLMode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func migrateModels() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.GrammarRule{},
		&models.Question{},
		&models.Option{},
		&models.Session{},
		&models.Answer{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database models: %v", err)
	}
} 