package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!", err)
		return
	}

	log.Println("Database connection successful.")

	// Auto-migrate schemas
	err = database.AutoMigrate(&User{}, &Score{}, &Category{}, &Topic{}, &Question{})
	if err != nil {
		log.Fatal("Failed to migrate database!", err)
		return
	}
	log.Println("Database migrated successfully.")

	DB = database
}
