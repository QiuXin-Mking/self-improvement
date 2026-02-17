package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// LearningData represents the structure of learning_data.json
type LearningData struct {
	Questions   map[string]Question `json:"questions"`
	LastUpdated string              `json:"last_updated"`
}

// Question represents a question in learning data (JSON format)
type Question struct {
	ID           string `json:"id"`
	Question     string `json:"question"`
	Answer       string `json:"answer"`
	Source       string `json:"source,omitempty"`
	Level        int    `json:"level"`
	NextReview   string `json:"next_review"`
	ReviewCount  int    `json:"review_count"`
	CorrectCount int    `json:"correct_count"`
	CreatedAt    string `json:"created_at"`
	LastReviewed string `json:"last_reviewed"`
}

// DBQuestion represents a question in SQLite database
type DBQuestion struct {
	ID           string     `gorm:"primaryKey"`
	UserID       uint       `gorm:"not null;index"`
	QuestionText string     `gorm:"not null"`
	AnswerText   string     `gorm:"not null"`
	Source       string     `gorm:""`
	Level        int        `gorm:""`
	NextReview   time.Time  `gorm:""`
	ReviewCount  int        `gorm:""`
	CorrectCount int        `gorm:""`
	CreatedAt    time.Time  `gorm:""`
	LastReviewed *time.Time `gorm:""`
	UpdatedAt    time.Time  `gorm:""`
	DeletedAt    *gorm.DeletedAt `gorm:"index"`
}

// TableName sets the table name for DBQuestion model
func (DBQuestion) TableName() string {
	return "questions"
}

// ResetQuestion resets a question's learning progress to initial state (JSON format)
func ResetQuestion(q Question) Question {
	now := time.Now().Format(time.RFC3339Nano)
	return Question{
		ID:           q.ID,
		Question:     q.Question,
		Answer:       q.Answer,
		Source:       q.Source,
		Level:        1,     // Reset to initial level
		NextReview:   now,    // Set next review to now
		ReviewCount:  0,      // Reset review count
		CorrectCount: 0,      // Reset correct count
		CreatedAt:    q.CreatedAt,
		LastReviewed: "",     // Clear last reviewed time
	}
}

func clearJSONData() (int, error) {
	dataFile := "data/learning_data.json"

	// Check if file exists
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return 0, nil // File doesn't exist, skip
	}

	// Read file
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return 0, fmt.Errorf("读取 JSON 文件失败: %w", err)
	}

	// Parse JSON
	var learningData LearningData
	if err := json.Unmarshal(data, &learningData); err != nil {
		return 0, fmt.Errorf("解析 JSON 失败: %w", err)
	}

	// Reset all questions
	questionCount := len(learningData.Questions)
	for id, question := range learningData.Questions {
		learningData.Questions[id] = ResetQuestion(question)
	}

	// Update last_updated
	learningData.LastUpdated = time.Now().Format(time.RFC3339Nano)

	// Write back to file
	outputData, err := json.MarshalIndent(learningData, "", "  ")
	if err != nil {
		return 0, fmt.Errorf("序列化 JSON 失败: %w", err)
	}

	if err := os.WriteFile(dataFile, outputData, 0644); err != nil {
		return 0, fmt.Errorf("写入 JSON 文件失败: %w", err)
	}

	return questionCount, nil
}

func clearSQLiteData() (int, error) {
	dbFile := "data/app.db"

	// Check if file exists
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return 0, nil // File doesn't exist, skip
	}

	// Open database
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return 0, fmt.Errorf("打开数据库失败: %w", err)
	}

	// Reset all questions' learning progress
	now := time.Now()
	result := db.Model(&DBQuestion{}).Where("1 = 1").Updates(map[string]interface{}{
		"level":         1,
		"next_review":   now,
		"review_count":   0,
		"correct_count":  0,
		"last_reviewed": nil,
	})

	if result.Error != nil {
		return 0, fmt.Errorf("更新数据库失败: %w", result.Error)
	}

	return int(result.RowsAffected), nil
}

func main() {
	// Clear JSON data
	jsonCount, err := clearJSONData()
	if err != nil {
		fmt.Printf("清除 JSON 数据失败: %v\n", err)
	} else if jsonCount > 0 {
		fmt.Printf("✓ 成功清除 JSON 数据中 %d 个问题的复习进度\n", jsonCount)
	}

	// Clear SQLite data
	sqliteCount, err := clearSQLiteData()
	if err != nil {
		fmt.Printf("清除 SQLite 数据失败: %v\n", err)
	} else if sqliteCount > 0 {
		fmt.Printf("✓ 成功清除 SQLite 数据库中 %d 个问题的复习进度\n", sqliteCount)
	}

	totalCount := jsonCount + sqliteCount
	if totalCount == 0 {
		fmt.Println("没有找到需要清除的数据")
	} else {
		fmt.Printf("\n总计: 成功清除 %d 个问题的复习进度数据\n", totalCount)
	}
}
