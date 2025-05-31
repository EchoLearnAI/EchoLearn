package main

import "time"

// UserDTO for Swagger documentation (flattens gorm.Model)
type UserDTO struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Email     string    `json:"email"`
	// Password deliberately omitted
}

// ScoreDTO for Swagger documentation (flattens gorm.Model)
type ScoreDTO struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserID      uint      `json:"userId"`
	Topic       string    `json:"topic"`
	Difficulty  string    `json:"difficulty"`
	Points      int       `json:"points"`
	Total       int       `json:"total"`
	AttemptedAt time.Time `json:"attemptedAt"`
}
