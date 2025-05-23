package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Answer represents a user's answer to a question within a game session.
// It's embedded in GameSession for simplicity now, but could be its own table.

// GameSession represents a single play session for a user.
// It tracks the mode, questions presented, answers, score, and status.
type GameSession struct {
	ID              string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	UserID          string    `gorm:"type:varchar(36);index" json:"user_id"`
	Mode            GameMode  `json:"mode"` // e.g., "survival", "five_topic", "infinite"
	Score           int       `json:"score"`
	MistakesMade    int       `json:"mistakes_made"`
	MaxMistakes     int       `json:"max_mistakes,omitempty"` // For Survival Mode
	CurrentQuestion int       `json:"current_question_index"` // Index for modes like 5-topic
	TotalQuestions  int       `json:"total_questions,omitempty"` // For modes like 5-topic
	IsActive        bool      `gorm:"default:true" json:"is_active"`
	StartedAt       time.Time `json:"started_at"`
	EndedAt         time.Time `gorm:"index" json:"ended_at,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	// For simplicity in MVP, storing answered questions and their results directly.
	// In a larger system, this would be a separate table `SessionAnswers`.
	AnsweredQuestionsJSON string `gorm:"type:text" json:"-"` // JSON string of AnsweredQuestionDetail
}

// AnsweredQuestionDetail stores information about a question answered in a session.
// This is used for the session summary.
// We'll store an array of these as a JSON string in GameSession.AnsweredQuestionsJSON.
type AnsweredQuestionDetail struct {
	QuestionID      string `json:"question_id"`
	QuestionText    string `json:"question_text"`
	SelectedOptionID string `json:"selected_option_id"`
	CorrectOptionID string `json:"correct_option_id"`
	IsCorrect       bool   `json:"is_correct"`
	Explanation     string `json:"explanation,omitempty"` // Explanation for the correct answer
}

// BeforeCreate will set a UUID for the game session.
func (gs *GameSession) BeforeCreate(tx *gorm.DB) (err error) {
	if gs.ID == "" {
		gs.ID = uuid.New().String()
	}
	gs.StartedAt = time.Now()
	gs.IsActive = true
	return
} 