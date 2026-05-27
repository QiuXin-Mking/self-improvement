package spacedrepetition

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"self-improvement/internal/models"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.Question{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func TestAddQuestion(t *testing.T) {
	db := setupTestDB(t)
	sr := NewSpacedRepetition(db)

	err := sr.AddQuestion(1, "q_1_abc", "什么是闭包？", "闭包是指有权访问另一个函数作用域中变量的函数。", "test.md", "js")
	if err != nil {
		t.Fatalf("AddQuestion failed: %v", err)
	}

	q, err := sr.GetQuestion(1, "q_1_abc")
	if err != nil {
		t.Fatalf("GetQuestion failed: %v", err)
	}

	if q.QuestionText != "什么是闭包？" {
		t.Errorf("question text mismatch: %s", q.QuestionText)
	}
	if q.Level != 4 {
		t.Errorf("expected level 4, got %d", q.Level)
	}
	if q.ReviewCount != 0 {
		t.Errorf("expected 0 reviews, got %d", q.ReviewCount)
	}
	if q.Category != "js" {
		t.Errorf("expected category 'js', got '%s'", q.Category)
	}
}

func TestGetDueQuestions(t *testing.T) {
	db := setupTestDB(t)
	sr := NewSpacedRepetition(db)

	// New question should be immediately due (NextReview = now)
	sr.AddQuestion(1, "q1", "Q1", "A1", "test.md", "go")

	questions, err := sr.GetDueQuestions(1)
	if err != nil {
		t.Fatalf("GetDueQuestions failed: %v", err)
	}
	if len(questions) != 1 {
		t.Fatalf("expected 1 due question, got %d", len(questions))
	}
}

func TestGetDueQuestionsByCategory(t *testing.T) {
	db := setupTestDB(t)
	sr := NewSpacedRepetition(db)

	sr.AddQuestion(1, "q1", "Q1", "A1", "test.md", "go")
	sr.AddQuestion(1, "q2", "Q2", "A2", "test.md", "python")

	questions, err := sr.GetDueQuestionsByCategory(1, "go")
	if err != nil {
		t.Fatalf("GetDueQuestionsByCategory failed: %v", err)
	}
	if len(questions) != 1 {
		t.Fatalf("expected 1 due question, got %d", len(questions))
	}
	if questions[0].Category != "go" {
		t.Errorf("expected category 'go', got '%s'", questions[0].Category)
	}
}

func TestUpdateReview_Proficient(t *testing.T) {
	db := setupTestDB(t)
	sr := NewSpacedRepetition(db)

	sr.AddQuestion(1, "q1", "Q1", "A1", "test.md", "go")

	err := sr.UpdateReview(1, "q1", 1) // feedback=1 (proficient)
	if err != nil {
		t.Fatalf("UpdateReview failed: %v", err)
	}

	q, _ := sr.GetQuestion(1, "q1")
	if q.ReviewCount != 1 {
		t.Errorf("expected 1 review, got %d", q.ReviewCount)
	}
	if q.CorrectCount != 1 {
		t.Errorf("expected 1 correct, got %d", q.CorrectCount)
	}

	// Base = 7d * 2.5 * 1.2 (accuracy=100% > 80%) = 21 days
	expectedMin := time.Now().Add(20 * 24 * time.Hour)
	expectedMax := time.Now().Add(22 * 24 * time.Hour)
	if q.NextReview.Before(expectedMin) || q.NextReview.After(expectedMax) {
		t.Errorf("next_review out of range: %v", q.NextReview)
	}
}

func TestUpdateReview_CompletelyForgotten(t *testing.T) {
	db := setupTestDB(t)
	sr := NewSpacedRepetition(db)

	sr.AddQuestion(1, "q1", "Q1", "A1", "test.md", "go")

	err := sr.UpdateReview(1, "q1", 4) // feedback=4 (completely forgotten)
	if err != nil {
		t.Fatalf("UpdateReview failed: %v", err)
	}

	q, _ := sr.GetQuestion(1, "q1")
	if q.CorrectCount != 0 {
		t.Errorf("expected 0 correct, got %d", q.CorrectCount)
	}

	// Base = 2h * 1.0 = 2h
	expected := time.Now().Add(2 * time.Hour)
	if q.NextReview.Sub(expected).Abs() > time.Minute {
		t.Errorf("next_review around +2h, got %v", q.NextReview)
	}
}

func TestUpdateReview_LevelDecrease(t *testing.T) {
	db := setupTestDB(t)
	sr := NewSpacedRepetition(db)

	sr.AddQuestion(1, "q1", "Q1", "A1", "test.md", "go")
	// Level decreases by 1 per correct review after 3+ correct.
	// Level 4 -> 3 -> 2 -> 1 takes 5 correct reviews
	for i := 0; i < 5; i++ {
		sr.UpdateReview(1, "q1", 1)
	}

	q, _ := sr.GetQuestion(1, "q1")
	if q.Level != 1 {
		t.Errorf("expected level 1 after 5 correct, got %d", q.Level)
	}
}

func TestUpdateReview_LevelIncrease(t *testing.T) {
	db := setupTestDB(t)
	sr := NewSpacedRepetition(db)

	sr.AddQuestion(1, "q1", "Q1", "A1", "test.md", "go")

	// After forgetting, level should stay at 4 (already max)
	sr.UpdateReview(1, "q1", 4)
	q, _ := sr.GetQuestion(1, "q1")
	if q.Level != 4 {
		t.Errorf("expected level 4 (max clamped), got %d", q.Level)
	}

	// Get 5 correct answers to reach level 1
	for i := 0; i < 5; i++ {
		sr.UpdateReview(1, "q1", 1)
	}
	q, _ = sr.GetQuestion(1, "q1")
	if q.Level != 1 {
		t.Errorf("expected level 1, got %d", q.Level)
	}

	// Now forget it — level should increase by 1
	sr.UpdateReview(1, "q1", 3)
	q, _ = sr.GetQuestion(1, "q1")
	if q.Level != 2 {
		t.Errorf("expected level 2 after forgetting, got %d", q.Level)
	}
}

func TestUpdateReview_AccuracyBoost(t *testing.T) {
	db := setupTestDB(t)
	sr := NewSpacedRepetition(db)

	sr.AddQuestion(1, "q1", "Q1", "A1", "test.md", "go")
	// First review: accuracy = 100% > 80%
	sr.UpdateReview(1, "q1", 1)

	q, _ := sr.GetQuestion(1, "q1")
	// Base = 7d * 2.5 * 1.2 = 21 days
	expectedMin := time.Now().Add(20 * 24 * time.Hour)
	if q.NextReview.Before(expectedMin) {
		t.Errorf("next_review should have accuracy boost: %v", q.NextReview)
	}
}

func TestDeleteQuestion(t *testing.T) {
	db := setupTestDB(t)
	sr := NewSpacedRepetition(db)

	sr.AddQuestion(1, "q1", "Q1", "A1", "test.md", "go")

	err := sr.DeleteQuestion(1, "q1")
	if err != nil {
		t.Fatalf("DeleteQuestion failed: %v", err)
	}

	_, err = sr.GetQuestion(1, "q1")
	if err == nil {
		t.Error("expected error getting deleted question")
	}
}

func TestGetStats(t *testing.T) {
	db := setupTestDB(t)
	sr := NewSpacedRepetition(db)

	sr.AddQuestion(1, "q1", "Q1", "A1", "test.md", "go")
	sr.AddQuestion(1, "q2", "Q2", "A2", "test.md", "python")

	stats, err := sr.GetStats(1)
	if err != nil {
		t.Fatalf("GetStats failed: %v", err)
	}

	if stats["total_questions"].(int) != 2 {
		t.Errorf("expected 2 total, got %v", stats["total_questions"])
	}
	if stats["due_questions"].(int) != 2 {
		t.Errorf("expected 2 due, got %v", stats["due_questions"])
	}
	if stats["total_reviews"].(int) != 0 {
		t.Errorf("expected 0 reviews, got %v", stats["total_reviews"])
	}
	if stats["accuracy"].(string) != "0.00" {
		t.Errorf("expected accuracy 0.00, got %s", stats["accuracy"])
	}

	// After one review with correct answer
	sr.UpdateReview(1, "q1", 1)
	stats, _ = sr.GetStats(1)
	if stats["total_reviews"].(int) != 1 {
		t.Errorf("expected 1 review, got %v", stats["total_reviews"])
	}
	if stats["accuracy"].(string) != "100.00" {
		t.Errorf("expected accuracy 100.00, got %s", stats["accuracy"])
	}
}

func TestGetCategories(t *testing.T) {
	db := setupTestDB(t)
	sr := NewSpacedRepetition(db)

	sr.AddQuestion(1, "q1", "Q1", "A1", "test.md", "go")
	sr.AddQuestion(1, "q2", "Q2", "A2", "test.md", "go")
	sr.AddQuestion(1, "q3", "Q3", "A3", "test.md", "python")

	categories, err := sr.GetCategories(1)
	if err != nil {
		t.Fatalf("GetCategories failed: %v", err)
	}
	if len(categories) != 2 {
		t.Fatalf("expected 2 categories, got %d", len(categories))
	}

	goCat := categories[0]
	if goCat["name"].(string) != "go" {
		t.Errorf("expected category 'go', got '%s'", goCat["name"])
	}
	if goCat["total"].(int64) != 2 {
		t.Errorf("expected 2 questions in go, got %d", goCat["total"])
	}
	if goCat["due"].(int64) != 2 {
		t.Errorf("expected 2 due in go, got %d", goCat["due"])
	}
}

func TestHash(t *testing.T) {
	h1 := Hash("什么是闭包？")
	h2 := Hash("什么是闭包？")
	h3 := Hash(" 什么是闭包？ ") // trimmed
	h4 := Hash("不同的问题")

	if h1 != h2 {
		t.Error("same input should produce same hash")
	}
	if h1 != h3 {
		t.Error("trimmed input should produce same hash")
	}
	if h1 == h4 {
		t.Error("different inputs should produce different hashes")
	}
}

func TestMultiUserIsolation(t *testing.T) {
	db := setupTestDB(t)
	sr := NewSpacedRepetition(db)

	sr.AddQuestion(1, "q1", "Q1", "A1", "test.md", "go")
	sr.AddQuestion(2, "q2", "Q2", "A2", "test.md", "go")

	// User 1 should only see their question
	q1, _ := sr.GetDueQuestions(1)
	if len(q1) != 1 {
		t.Errorf("user 1 should have 1 question, got %d", len(q1))
	}

	// User 2 should only see their question
	q2, _ := sr.GetDueQuestions(2)
	if len(q2) != 1 {
		t.Errorf("user 2 should have 1 question, got %d", len(q2))
	}

	// User 1 can't update user 2's question
	err := sr.UpdateReview(1, "q2", 1)
	if err == nil {
		t.Error("should fail updating other user's question")
	}
}

// Ensure min/max helpers work correctly
func TestMinMax(t *testing.T) {
	if min(1, 2) != 1 {
		t.Error("min(1,2) should be 1")
	}
	if min(2, 1) != 1 {
		t.Error("min(2,1) should be 1")
	}
	if max(1, 2) != 2 {
		t.Error("max(1,2) should be 2")
	}
	if max(2, 1) != 2 {
		t.Error("max(2,1) should be 2")
	}
}
