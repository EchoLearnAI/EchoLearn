package main

import (
	"log"

	"github.com/EchoLearnAI/EchoLearn/api"
	"github.com/EchoLearnAI/EchoLearn/config"
	"github.com/EchoLearnAI/EchoLearn/data"
	"github.com/EchoLearnAI/EchoLearn/db"

	"github.com/gin-gonic/gin"
)

// @title EchoLearn API
// @version 1.0
// @description API for EchoLearn English learning app.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load configuration
	config.LoadConfig() // This will load from .env or defaults

	// Set Gin mode (optional, can be set via GIN_MODE env var)
	// gin.SetMode(gin.ReleaseMode) // Uncomment for production
	gin.SetMode(gin.DebugMode) // Default to debug mode

	// Initialize database connection
	_, err := db.Connect() // gorm.DB instance is stored internally by db package
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Seed database if it's empty
    // Check if the database file itself exists before trying to count questions.
    // This is a simple way to handle first-time run for SQLite.
    // For other DBs, the check might be different (e.g., check if tables exist).
    if !db.DatabaseExists(config.AppConfig.DBName) {
        log.Println("Database file not found. It will be created and seeded.")
        // db.Connect() already creates and migrates.
        // We can directly call Seed here.
        db.SeedDatabase(data.Seed) // Pass the actual Seed function
    } else {
        // If DB file exists, check if it needs seeding (e.g., no questions)
        db.SeedDatabase(data.Seed) 
    }

	// Setup router
	router := api.SetupRouter()

	serverAddr := ":" + config.AppConfig.ServerPort
	log.Printf("Server starting on %s", serverAddr)

	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
} 