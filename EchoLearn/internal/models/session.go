package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameMode string

const (
	SurvivalMode GameMode = "survival"
	TopicMode    GameMode = "topic"
	InfiniteMode GameMode = "infinite"
)

type Session struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID     uuid.UUID      `json:"user_id" gorm:"type:uuid"`
	User       User           `json:"user" gorm:"foreignKey:UserID"`
	Mode       GameMode       `json:"mode" gorm:"not null"`
	Score      int            `json:"score" gorm:"default:0"`
	Mistakes   int            `json:"mistakes" gorm:"default:0"`
	IsComplete bool           `json:"is_complete" gorm:"default:false"`
	Answers    []Answer       `json:"answers" gorm:"foreignKey:SessionID"`
	StartedAt  time.Time      `json:"started_at"`
	EndedAt    *time.Time     `json:"ended_at"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Answer struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	SessionID  uuid.UUID      `json:"session_id" gorm:"type:uuid"`
	QuestionID uuid.UUID      `json:"question_id" gorm:"type:uuid"`
	Question   Question       `json:"question" gorm:"foreignKey:QuestionID"`
	OptionID   uuid.UUID      `json:"option_id" gorm:"type:uuid"`
	Option     Option         `json:"option" gorm:"foreignKey:OptionID"`
	IsCorrect  bool           `json:"is_correct"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type SessionSummary struct {
	TotalQuestions int           `json:"total_questions"`
	CorrectAnswers int           `json:"correct_answers"`
	Mistakes       int           `json:"mistakes"`
	Score          int           `json:"score"`
	GameMode       GameMode      `json:"game_mode"`
	Duration       time.Duration `json:"duration"`
	ByCategory     []CategoryScore `json:"by_category"`
}

type CategoryScore struct {
	Category       string `json:"category"`
	TotalQuestions int    `json:"total_questions"`
	CorrectAnswers int    `json:"correct_answers"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	s.StartedAt = time.Now()
	return nil
}

func (a *Answer) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
} 