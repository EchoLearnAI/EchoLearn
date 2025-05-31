package main

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model        // Includes ID, CreatedAt, UpdatedAt, DeletedAt
	Email      string `gorm:"uniqueIndex;not null" json:"email"`
	Password   string `gorm:"not null" json:"-"` // json:"-" to prevent sending password hash in responses
}

// Score model
type Score struct {
	gorm.Model
	UserID      uint      `gorm:"not null" json:"userId"`
	User        User      `gorm:"foreignKey:UserID" json:"-"`
	Topic       string    `gorm:"not null" json:"topic"` // One of the 19 categories
	Difficulty  string    `gorm:"not null" json:"difficulty"`
	Points      int       `gorm:"not null" json:"points"` // Score for this attempt
	Total       int       `gorm:"not null" json:"total"`  // Total questions in this attempt
	AttemptedAt time.Time `gorm:"not null" json:"attemptedAt"`
}
