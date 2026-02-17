package models

import (
	"time"

	"gorm.io/gorm"
)

// Question represents a question-answer pair with learning data
type Question struct {
	ID           string         `json:"id" gorm:"primaryKey"`
	UserID       uint           `json:"user_id" gorm:"not null;index"`
	QuestionText string         `json:"question" gorm:"not null"`
	AnswerText   string         `json:"answer" gorm:"not null"`
	Source       string         `json:"source"`
	Level        int            `json:"level"`             // 1-4: 1=proficient, 2=fair, 3=forgotten, 4=completely forgotten
	NextReview   time.Time      `json:"next_review"`       // Next review scheduled time
	ReviewCount  int            `json:"review_count"`      // Total number of reviews
	CorrectCount int            `json:"correct_count"`     // Number of correct answers
	CreatedAt    time.Time      `json:"created_at"`        // When question was added
	LastReviewed *time.Time     `json:"last_reviewed"`     // When last reviewed (nil if never)
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Associations
	User User `json:"user" gorm:"foreignKey:UserID"`
}

// TableName sets the table name for Question model
func (Question) TableName() string {
	return "questions"
}