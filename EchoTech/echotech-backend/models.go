package main

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model        // Includes ID, CreatedAt, UpdatedAt, DeletedAt
	Email      string `gorm:"uniqueIndex;not null" json:"email"`
	Password   string `gorm:"not null" json:"-"` // json:"-" to prevent sending password hash in responses
}

// Category model
type Category struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null" json:"name"`
	Slug string `gorm:"uniqueIndex;not null" json:"slug"`
}

// Topic model
type Topic struct {
	gorm.Model
	Name       string   `gorm:"not null" json:"name"`
	Slug       string   `gorm:"uniqueIndex;not null" json:"slug"`
	CategoryID uint     `gorm:"not null" json:"categoryId"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

// Option struct for questions - will be stored as JSONB
type Option struct {
	ID          string `json:"id"`
	Text        string `json:"text"`
	IsCorrect   bool   `json:"isCorrect"`
	Explanation string `json:"explanation"`
}

// Question model
type Question struct {
	gorm.Model
	TopicID    uint           `gorm:"not null" json:"topicId"`
	Topic      Topic          `gorm:"foreignKey:TopicID" json:"topic,omitempty"`
	Difficulty string         `gorm:"not null" json:"difficulty"`
	Text       string         `gorm:"not null" json:"text"`
	Options    datatypes.JSON `gorm:"type:jsonb" json:"options"` // Storing options as JSONB
	// OriginalID can be used to store the ID from the data.go file if needed during seeding
	OriginalID string `gorm:"uniqueIndex" json:"originalId,omitempty"`
}

// Score model
type Score struct {
	gorm.Model
	UserID      uint      `gorm:"not null" json:"userId"`
	User        User      `gorm:"foreignKey:UserID" json:"-"`
	TopicID     uint      `gorm:"not null" json:"topicId"`
	Topic       Topic     `gorm:"foreignKey:TopicID" json:"topic,omitempty"` // Include Topic details if needed
	Difficulty  string    `gorm:"not null" json:"difficulty"`
	Points      int       `gorm:"not null" json:"points"` // Score for this attempt
	Total       int       `gorm:"not null" json:"total"`  // Total questions in this attempt
	AttemptedAt time.Time `gorm:"not null" json:"attemptedAt"`
}
