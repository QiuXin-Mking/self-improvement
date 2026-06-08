package spacedrepetition

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"self-improvement/internal/models"
)

// SpacedRepetition manages the spaced repetition algorithm with multi-tenancy
type SpacedRepetition struct {
	DB *gorm.DB
}

// NewSpacedRepetition creates a new spaced repetition instance
func NewSpacedRepetition(db *gorm.DB) *SpacedRepetition {
	return &SpacedRepetition{
		DB: db,
	}
}

// AddQuestion adds a new question to the user's knowledge base
func (sr *SpacedRepetition) AddQuestion(userID uint, id, question, answer, source, category string) error {
	now := time.Now()
	q := models.Question{
		ID:           id,
		UserID:       userID,
		QuestionText: question,
		AnswerText:   answer,
		Source:       source,
		Category:     category,
		Level:        4, // Start as completely forgotten
		NextReview:   now,
		ReviewCount:  0,
		CorrectCount: 0,
		CreatedAt:    now,
		LastReviewed: nil,
	}

	return sr.DB.Create(&q).Error
}

// GetDueQuestions returns questions that are due for review for a specific user
func (sr *SpacedRepetition) GetDueQuestions(userID uint) ([]*models.Question, error) {
	now := time.Now()
	var questions []*models.Question

	err := sr.DB.Where("user_id = ? AND next_review <= ?", userID, now).
		Order("next_review ASC").
		Limit(100).
		Find(&questions).Error

	if err != nil {
		return nil, err
	}

	return questions, nil
}

// GetDueQuestionsByCategory returns due questions filtered by category
func (sr *SpacedRepetition) GetDueQuestionsByCategory(userID uint, category string) ([]*models.Question, error) {
	now := time.Now()
	var questions []*models.Question

	err := sr.DB.Where("user_id = ? AND next_review <= ? AND category = ?", userID, now, category).
		Order("next_review ASC").
		Find(&questions).Error

	if err != nil {
		return nil, err
	}

	return questions, nil
}

// GetDueQuestionsByCategories returns due questions filtered by multiple categories
func (sr *SpacedRepetition) GetDueQuestionsByCategories(userID uint, categories []string) ([]*models.Question, error) {
	now := time.Now()
	var questions []*models.Question

	err := sr.DB.Where("user_id = ? AND next_review <= ? AND category IN ?", userID, now, categories).
		Order("next_review ASC").
		Limit(100).
		Find(&questions).Error

	if err != nil {
		return nil, err
	}

	return questions, nil
}

// GetCategories returns all distinct categories with stats for a user
func (sr *SpacedRepetition) GetCategories(userID uint) ([]map[string]interface{}, error) {
	type CategoryStats struct {
		Category string
		Total    int64
	}

	var results []CategoryStats
	err := sr.DB.Model(&models.Question{}).
		Select("category, COUNT(*) as total").
		Where("user_id = ? AND category != ''", userID).
		Group("category").
		Order("category ASC").
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	now := time.Now()
	var categories []map[string]interface{}
	for _, r := range results {
		var due int64
		sr.DB.Model(&models.Question{}).
			Where("user_id = ? AND category = ? AND next_review <= ?", userID, r.Category, now).
			Count(&due)

		categories = append(categories, map[string]interface{}{
			"name":  r.Category,
			"total": r.Total,
			"due":   due,
		})
	}

	return categories, nil
}

// GetQuestion returns a specific question by ID for the user
func (sr *SpacedRepetition) GetQuestion(userID uint, id string) (*models.Question, error) {
	var q models.Question

	err := sr.DB.Where("user_id = ? AND id = ?", userID, id).First(&q).Error
	if err != nil {
		return nil, err
	}

	return &q, nil
}

// UpdateReview updates review results for a question
func (sr *SpacedRepetition) UpdateReview(userID uint, id string, feedback int) error {
	question, err := sr.GetQuestion(userID, id)
	if err != nil {
		return err
	}

	now := time.Now()
	question.ReviewCount++
	question.LastReviewed = &now

	// Update statistics
	if feedback <= 2 { // Proficient or fair counts as correct
		question.CorrectCount++
	}

	// Calculate next review time based on feedback
	intervals := map[int]time.Duration{
		1: 7 * 24 * time.Hour, // Proficient: 7 days
		2: 3 * 24 * time.Hour, // Fair: 3 days
		3: 24 * time.Hour,     // Forgotten: 1 day
		4: 2 * time.Hour,      // Completely forgotten: 2 hours
	}

	baseInterval := intervals[feedback]

	// Apply multipliers based on feedback level and historical accuracy
	multipliers := map[int]float64{
		1: 2.5, // Proficient gets 2.5x multiplier
		2: 1.8, // Fair gets 1.8x multiplier
		3: 1.3, // Forgotten gets 1.3x multiplier
		4: 1.0, // Completely forgotten stays at 1.0x
	}

	multiplier := multipliers[feedback]

	// Adjust multiplier based on historical accuracy
	if question.ReviewCount > 0 {
		accuracy := float64(question.CorrectCount) / float64(question.ReviewCount)
		if accuracy > 0.8 { // High accuracy gets additional boost
			multiplier *= 1.2
		}
	}

	// Calculate final interval
	intervalHours := float64(baseInterval.Hours()) * multiplier
	nextReview := now.Add(time.Duration(intervalHours * float64(time.Hour)))

	question.NextReview = nextReview

	// Update memory level
	if feedback <= 2 {
		// If answered correctly, potentially upgrade level (decrease number)
		if question.CorrectCount >= 3 && question.Level > 1 {
			question.Level = max(1, question.Level-1)
		}
	} else {
		// If forgotten, downgrade level (increase number)
		question.Level = min(4, question.Level+1)
	}

	return sr.DB.Save(question).Error
}

// ResetUserQuestions resets all questions for a user to fresh state.
// All questions become due immediately with review counts zeroed.
func (sr *SpacedRepetition) ResetUserQuestions(userID uint) error {
	now := time.Now()
	return sr.DB.Model(&models.Question{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"level":         4,
			"next_review":   now,
			"review_count":  0,
			"correct_count": 0,
			"last_reviewed": nil,
		}).Error
}

// GetForecast returns review counts for the next N days
func (sr *SpacedRepetition) GetForecast(userID uint, days int) ([]map[string]interface{}, error) {
	now := time.Now()
	var forecast []map[string]interface{}

	for i := 0; i < days; i++ {
		dayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, i)
		dayEnd := dayStart.AddDate(0, 0, 1).Add(-time.Second)

		var count int64
		sr.DB.Model(&models.Question{}).
			Where("user_id = ? AND next_review BETWEEN ? AND ?", userID, dayStart, dayEnd).
			Count(&count)

		forecast = append(forecast, map[string]interface{}{
			"date":  dayStart.Format("2006-01-02"),
			"count": count,
		})
	}

	// Also include overdue items (next_review <= now) in today's count
	if len(forecast) > 0 {
		var overdue int64
		sr.DB.Model(&models.Question{}).
			Where("user_id = ? AND next_review <= ?", userID, now).
			Count(&overdue)
		forecast[0]["count"] = overdue
	}

	return forecast, nil
}

// DeleteQuestion removes a question from the user's knowledge base
func (sr *SpacedRepetition) DeleteQuestion(userID uint, id string) error {
	result := sr.DB.Where("user_id = ? AND id = ?", userID, id).Delete(&models.Question{})
	return result.Error
}

// GetStats returns learning statistics for a specific user
func (sr *SpacedRepetition) GetStats(userID uint) (map[string]interface{}, error) {
	var questions []*models.Question
	err := sr.DB.Where("user_id = ?", userID).Find(&questions).Error
	if err != nil {
		return nil, err
	}

	total := len(questions)

	var due int
	now := time.Now()
	for _, q := range questions {
		if q.NextReview.Before(now) || q.NextReview.Equal(now) {
			due++
		}
	}

	totalReviews := 0
	totalCorrect := 0

	for _, q := range questions {
		totalReviews += q.ReviewCount
		totalCorrect += q.CorrectCount
	}

	accuracy := 0.0
	if totalReviews > 0 {
		accuracy = float64(totalCorrect) / float64(totalReviews) * 100
	}

	return map[string]interface{}{
		"total_questions": total,
		"due_questions":   due,
		"total_reviews":   totalReviews,
		"total_correct":   totalCorrect,
		"accuracy":        fmt.Sprintf("%.2f", accuracy),
	}, nil
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Hash generates a hash for a string (used for question IDs)
func Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(strings.TrimSpace(s)))
	return hex.EncodeToString(h.Sum(nil))
}
