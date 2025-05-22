package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Text      string         `json:"text" gorm:"not null"`
	Category  string         `json:"category" gorm:"not null"`
	Options   []Option       `json:"options" gorm:"foreignKey:QuestionID"`
	GrammarID uuid.UUID      `json:"grammar_id" gorm:"type:uuid"`
	Grammar   GrammarRule    `json:"grammar" gorm:"foreignKey:GrammarID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Option struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	QuestionID  uuid.UUID      `json:"question_id" gorm:"type:uuid"`
	Label       string         `json:"label" gorm:"not null"`
	IsCorrect   bool           `json:"is_correct"`
	Explanation string         `json:"explanation"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type GrammarRule struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	Examples    string         `json:"examples"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (q *Question) BeforeCreate(tx *gorm.DB) error {
	if q.ID == uuid.Nil {
		q.ID = uuid.New()
	}
	return nil
}

func (o *Option) BeforeCreate(tx *gorm.DB) error {
	if o.ID == uuid.Nil {
		o.ID = uuid.New()
	}
	return nil
}

func (g *GrammarRule) BeforeCreate(tx *gorm.DB) error {
	if g.ID == uuid.Nil {
		g.ID = uuid.New()
	}
	return nil
} 