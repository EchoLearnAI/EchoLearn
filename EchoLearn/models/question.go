package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Question represents a multiple-choice question.
// It includes the question text, options, correct answer, explanation,
// and associated grammar rule and category.

// Option represents one choice in a multiple-choice question.
type Option struct {
	ID          string `gorm:"type:varchar(36);primaryKey" json:"id"`
	QuestionID  string `gorm:"type:varchar(36);index" json:"question_id"` // Foreign key to Question
	Label       string `json:"label"`       // e.g., "A", "B", "C", "D" or the text itself
	Text        string `json:"text"`        // The actual option text
	IsCorrect   bool   `json:"is_correct"`
	Explanation string `gorm:"type:text" json:"explanation,omitempty"` // Explanation for why this option is correct/incorrect
}

// BeforeCreate will set a UUID for the option.
func (opt *Option) BeforeCreate(tx *gorm.DB) (err error) {
	if opt.ID == "" {
		opt.ID = uuid.New().String()
	}
	return
}

// GrammarRule defines the grammar concept related to a question.
// Examples could be a JSON array of strings or a single string with newlines.
type GrammarRule struct {
	ID          string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Examples    string    `gorm:"type:text" json:"examples"` // Could be JSON array of strings or markdown
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// BeforeCreate will set a UUID for the grammar rule.
func (gr *GrammarRule) BeforeCreate(tx *gorm.DB) (err error) {
	if gr.ID == "" {
		gr.ID = uuid.New().String()
	}
	return
}

// Question represents a single question in the quiz
type Question struct {
	ID            string      `gorm:"type:varchar(36);primaryKey" json:"id"`
	Text          string      `gorm:"type:text" json:"text"`
	Options       []Option    `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE;" json:"options"`
	GrammarRuleID string      `gorm:"type:varchar(36);index" json:"grammar_rule_id"`
	GrammarRule   GrammarRule `gorm:"foreignKey:GrammarRuleID" json:"grammar_rule"`
	Category      string      `gorm:"index" json:"category"` // e.g., "Articles", "Prepositions"
	Difficulty    string      `json:"difficulty,omitempty"`  // e.g., "Easy", "Medium", "Hard"
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

// BeforeCreate will set a UUID for the question.
func (q *Question) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == "" {
		q.ID = uuid.New().String()
	}
	return
} 