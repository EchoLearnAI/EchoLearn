package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user in the system.
// For MVP, auth is not included.
// Progress will be a simple JSON or map for now.

// GameMode represents the type of quiz mode.
type GameMode string

const (
	GameModeMistakes          GameMode = "mistakes"           // For games ending after X mistakes
	GameModeCategoryChallenge GameMode = "category_challenge" // For games focused on a specific category
	GameModeInfinite          GameMode = "infinite"           // For endless play

	// Existing modes - can be kept if distinct or removed/aliased if redundant
	// SurvivalMode  GameMode = "survival" // If this is same as Mistakes, choose one.
	// FiveTopicMode GameMode = "five_topic" // If this is related to Category Challenge, choose one.
)

// User represents a user in the system
type User struct {
	ID           string        `gorm:"type:varchar(36);primary_key" json:"id"`
	Name         string        `gorm:"type:varchar(100);not null" json:"name"`
	Email        string        `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	PasswordHash string        `gorm:"type:varchar(255);not null;default:''" json:"-"` // Store hashed password, not sent in JSON
	Progress     string        `gorm:"type:text" json:"progress,omitempty"`            // Simple JSON string for progress
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Sessions     []GameSession `gorm:"foreignKey:UserID" json:"sessions,omitempty"` // User can have multiple game sessions
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	return
}
