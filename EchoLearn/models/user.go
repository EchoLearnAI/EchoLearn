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
	SurvivalMode GameMode = "survival"
	FiveTopicMode GameMode = "five_topic"
	InfiniteMode GameMode = "infinite"
)

// User represents a user in the system
type User struct {
	ID        string         `gorm:"type:varchar(36);primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `gorm:"uniqueIndex" json:"email,omitempty"` // Optional
	Progress  string         `gorm:"type:text" json:"progress,omitempty"` // Simple JSON string for progress
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Sessions  []GameSession `gorm:"foreignKey:UserID" json:"-"` // User can have multiple game sessions
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	return
} 