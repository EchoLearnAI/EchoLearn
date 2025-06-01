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

// CategoryDTO for Swagger documentation
type CategoryDTO struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
}

// TopicDTO for Swagger documentation
type TopicDTO struct {
	ID         uint        `json:"id"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	Name       string      `json:"name"`
	Slug       string      `json:"slug"`
	CategoryID uint        `json:"categoryId"`
	Category   CategoryDTO `json:"category,omitempty"` // Nested Category details
}

// OptionDTO for Swagger documentation (used within QuestionDTO)
type OptionDTO struct {
	ID          string `json:"id"`
	Text        string `json:"text"`
	IsCorrect   bool   `json:"isCorrect"`
	Explanation string `json:"explanation"`
}

// QuestionDTO for Swagger documentation
type QuestionDTO struct {
	ID         uint        `json:"id"`         // Database ID
	OriginalID string      `json:"originalId"` // e.g., "net-e-q1"
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	TopicID    uint        `json:"topicId"`
	Topic      *TopicDTO   `json:"topic,omitempty"` // Nested Topic details
	Difficulty string      `json:"difficulty"`
	Text       string      `json:"text"`
	Options    []OptionDTO `json:"options"`
}

// ScoreDTO for Swagger documentation (flattens gorm.Model)
type ScoreDTO struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserID      uint      `json:"userId"`
	TopicID     uint      `json:"topicId"`
	Topic       *TopicDTO `json:"topic,omitempty"` // Updated to use TopicDTO
	Difficulty  string    `json:"difficulty"`
	Points      int       `json:"points"`
	Total       int       `json:"total"`
	AttemptedAt time.Time `json:"attemptedAt"`
}
