package server

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"self-improvement/internal/models"
)

func setupE2E(t *testing.T) *gin.Engine {
	t.Helper()

	os.Setenv("JWT_SECRET", "e2e-test-secret")

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.Question{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	InitTestDB(db)
	return SetupRouter()
}

func registerAndGetToken(t *testing.T, router *gin.Engine, username string) string {
	t.Helper()
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{
		"username": username,
		"password": "testpass123",
	})
	req := httptest.NewRequest("POST", "/api/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if !resp.Success {
		t.Fatalf("registration failed: %s", resp.Error)
	}
	data := resp.Data.(map[string]interface{})
	return data["token"].(string)
}

func addQuestion(t *testing.T, router *gin.Engine, token, question, answer string) *httptest.ResponseRecorder {
	t.Helper()
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{
		"question": question,
		"answer":   answer,
	})
	req := httptest.NewRequest("POST", "/api/add-question", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)
	return w
}

// ═══════════════════════════════════════════
// Auth Tests
// ═══════════════════════════════════════════

func TestE2E_Register(t *testing.T) {
	router := setupE2E(t)

	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{
		"username": "newuser",
		"password": "secret123",
	})
	req := httptest.NewRequest("POST", "/api/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if !resp.Success {
		t.Fatalf("registration failed: %s", resp.Error)
	}

	data := resp.Data.(map[string]interface{})
	if data["token"].(string) == "" {
		t.Error("token should not be empty")
	}
	if data["username"].(string) != "newuser" {
		t.Errorf("expected username 'newuser', got '%s'", data["username"])
	}
}

func TestE2E_Register_Duplicate(t *testing.T) {
	router := setupE2E(t)
	registerAndGetToken(t, router, "dupuser")

	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{
		"username": "dupuser",
		"password": "secret123",
	})
	req := httptest.NewRequest("POST", "/api/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Errorf("expected 409, got %d", w.Code)
	}
}

func TestE2E_Register_ShortUsername(t *testing.T) {
	router := setupE2E(t)

	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{
		"username": "ab",
		"password": "secret123",
	})
	req := httptest.NewRequest("POST", "/api/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestE2E_Register_ShortPassword(t *testing.T) {
	router := setupE2E(t)

	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{
		"username": "validuser",
		"password": "12345",
	})
	req := httptest.NewRequest("POST", "/api/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestE2E_Login(t *testing.T) {
	router := setupE2E(t)
	registerAndGetToken(t, router, "loginuser")

	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{
		"username": "loginuser",
		"password": "testpass123",
	})
	req := httptest.NewRequest("POST", "/api/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if !resp.Success {
		t.Fatalf("login failed: %s", resp.Error)
	}
}

func TestE2E_Login_WrongPassword(t *testing.T) {
	router := setupE2E(t)
	registerAndGetToken(t, router, "pwuser")

	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{
		"username": "pwuser",
		"password": "wrongpassword",
	})
	req := httptest.NewRequest("POST", "/api/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestE2E_Login_NonExistent(t *testing.T) {
	router := setupE2E(t)

	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{
		"username": "nobody",
		"password": "whatever",
	})
	req := httptest.NewRequest("POST", "/api/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

// ═══════════════════════════════════════════
// Profile Tests
// ═══════════════════════════════════════════

func TestE2E_Profile(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "profileuser")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/profile", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestE2E_Profile_NoAuth(t *testing.T) {
	router := setupE2E(t)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/profile", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", w.Code)
	}
}

// ═══════════════════════════════════════════
// Stats Tests
// ═══════════════════════════════════════════

func TestE2E_Stats_Empty(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "statsuser")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/stats", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	stats := resp.Data.(map[string]interface{})["stats"].(map[string]interface{})
	if stats["total_questions"].(float64) != 0 {
		t.Errorf("expected 0 questions, got %v", stats["total_questions"])
	}
}

func TestE2E_Stats_WithQuestions(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "statsuser2")

	addQuestion(t, router, token, "Q1", "A1")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/stats", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	stats := resp.Data.(map[string]interface{})["stats"].(map[string]interface{})
	if stats["total_questions"].(float64) != 1 {
		t.Errorf("expected 1 question, got %v", stats["total_questions"])
	}
}

// ═══════════════════════════════════════════
// Add Question Tests
// ═══════════════════════════════════════════

func TestE2E_AddQuestion(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "adduser")

	w := addQuestion(t, router, token, "什么是闭包？", "闭包是指有权访问另一个函数作用域中变量的函数。")
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

func TestE2E_AddQuestion_Duplicate(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "dupquser")

	addQuestion(t, router, token, "唯一的问题", "唯一的答案")
	w := addQuestion(t, router, token, "唯一的问题", "唯一的答案")

	if w.Code != http.StatusConflict {
		t.Errorf("expected 409, got %d", w.Code)
	}
}

func TestE2E_AddQuestion_Empty(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "emptyquser")

	w := addQuestion(t, router, token, "", "")
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

// ═══════════════════════════════════════════
// Upload MD Tests
// ═══════════════════════════════════════════

func TestE2E_UploadMd(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "uploaduser")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.md")
	part.Write([]byte("# q\n什么是Rust的所有权？\n# a\nRust的所有权系统确保内存安全。\n"))
	writer.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/upload-md", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

func TestE2E_UploadMd_BadExtension(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "badextuser")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.txt")
	part.Write([]byte("not markdown"))
	writer.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/upload-md", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestE2E_UploadMd_NoFile(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "nofileuser")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/upload-md", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

// ═══════════════════════════════════════════
// Upload ZIP Tests
// ═══════════════════════════════════════════

func TestE2E_UploadZip(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "zipuser")

	zipPath := filepath.Join("..", "..", "test", "multi-level-kb.zip")
	zipData, err := os.ReadFile(zipPath)
	if err != nil {
		t.Skipf("test fixture not found: %s", zipPath)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "multi-level-kb.zip")
	part.Write(zipData)
	writer.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/upload-zip", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token)
	req.ContentLength = int64(body.Len())
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp.Data.(map[string]interface{})
	imported := data["imported"].(float64)
	if imported < 1 {
		t.Errorf("expected at least 1 imported question, got %v", imported)
	}
}

func TestE2E_UploadZip_BadExtension(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "badzipuser")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.rar")
	part.Write([]byte("not a zip"))
	writer.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/upload-zip", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

// ═══════════════════════════════════════════
// Due Questions Tests
// ═══════════════════════════════════════════

func TestE2E_DueQuestions_EmptyKB(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "emptykbuser")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/due-questions", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Success != false {
		t.Error("expected success=false for empty KB")
	}
}

func TestE2E_DueQuestions(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "dueuser")

	addQuestion(t, router, token, "问题1", "答案1")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/due-questions", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if !resp.Success {
		t.Fatalf("expected success: %s", resp.Error)
	}
	questions := resp.Data.(map[string]interface{})["questions"].([]interface{})
	if len(questions) != 1 {
		t.Errorf("expected 1 due question, got %d", len(questions))
	}
}

// ═══════════════════════════════════════════
// Update Review Tests
// ═══════════════════════════════════════════

func TestE2E_UpdateReview(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "reviewuser")

	addQuestion(t, router, token, "待复习的问题", "待复习的答案")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/due-questions", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	questions := resp.Data.(map[string]interface{})["questions"].([]interface{})
	q := questions[0].(map[string]interface{})
	qID := q["id"].(string)

	reviewW := httptest.NewRecorder()
	reviewBody, _ := json.Marshal(map[string]interface{}{
		"question_id": qID,
		"feedback":    1,
	})
	reviewReq := httptest.NewRequest("POST", "/api/update-review", bytes.NewReader(reviewBody))
	reviewReq.Header.Set("Content-Type", "application/json")
	reviewReq.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(reviewW, reviewReq)

	if reviewW.Code != http.StatusOK {
		t.Errorf("expected 200, got %d: %s", reviewW.Code, reviewW.Body.String())
	}
}

func TestE2E_UpdateReview_InvalidFeedback(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "badfeedback")

	addQuestion(t, router, token, "测试问题", "测试答案")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/due-questions", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	questions := resp.Data.(map[string]interface{})["questions"].([]interface{})
	qID := questions[0].(map[string]interface{})["id"].(string)

	w2 := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{"question_id": qID, "feedback": 0})
	req2 := httptest.NewRequest("POST", "/api/update-review", bytes.NewReader(body))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w2, req2)

	if w2.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for feedback=0, got %d", w2.Code)
	}
}

func TestE2E_UpdateReview_NonExistent(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "noreview")

	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"question_id": "nonexistent_id",
		"feedback":    1,
	})
	req := httptest.NewRequest("POST", "/api/update-review", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}

// ═══════════════════════════════════════════
// Delete Question Tests
// ═══════════════════════════════════════════

func TestE2E_DeleteQuestion(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "deluser")

	addQuestion(t, router, token, "要删除的问题", "要删除的答案")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/due-questions", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	questions := resp.Data.(map[string]interface{})["questions"].([]interface{})
	qID := questions[0].(map[string]interface{})["id"].(string)

	w2 := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{"question_id": qID})
	req2 := httptest.NewRequest("POST", "/api/delete-question", bytes.NewReader(body))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Errorf("expected 200, got %d: %s", w2.Code, w2.Body.String())
	}
}

func TestE2E_DeleteQuestion_NonExistent(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "nodeluser")

	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{"question_id": "nonexistent"})
	req := httptest.NewRequest("POST", "/api/delete-question", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

// ═══════════════════════════════════════════
// Categories Tests
// ═══════════════════════════════════════════

func TestE2E_Categories(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "catuser")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/categories", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

// ═══════════════════════════════════════════
// Full Learning Flow Test
// ═══════════════════════════════════════════

func TestE2E_FullLearningFlow(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "flowuser")

	// 1. Add 3 questions
	addQuestion(t, router, token, "知识点1", "答案1")
	addQuestion(t, router, token, "知识点2", "答案2")
	addQuestion(t, router, token, "知识点3", "答案3")

	// 2. Verify stats
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/stats", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)
	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	stats := resp.Data.(map[string]interface{})["stats"].(map[string]interface{})
	if stats["total_questions"].(float64) != 3 {
		t.Errorf("expected 3 total, got %v", stats["total_questions"])
	}
	if stats["due_questions"].(float64) != 3 {
		t.Errorf("expected 3 due, got %v", stats["due_questions"])
	}

	// 3. Get due questions
	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/api/due-questions", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &resp)
	questions := resp.Data.(map[string]interface{})["questions"].([]interface{})
	if len(questions) != 3 {
		t.Fatalf("expected 3 due questions, got %d", len(questions))
	}

	// 4. Review each question with different feedback
	feedbackLevels := []int{1, 2, 3}
	for i, q := range questions {
		qID := q.(map[string]interface{})["id"].(string)

		reviewW := httptest.NewRecorder()
		reviewBody, _ := json.Marshal(map[string]interface{}{
			"question_id": qID,
			"feedback":    feedbackLevels[i],
		})
		reviewReq := httptest.NewRequest("POST", "/api/update-review", bytes.NewReader(reviewBody))
		reviewReq.Header.Set("Content-Type", "application/json")
		reviewReq.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(reviewW, reviewReq)

		if reviewW.Code != http.StatusOK {
			t.Errorf("review %d failed with code %d", i, reviewW.Code)
		}
	}

	// 5. Verify final stats
	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/api/stats", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &resp)
	finalStats := resp.Data.(map[string]interface{})["stats"].(map[string]interface{})
	if finalStats["total_reviews"].(float64) != 3 {
		t.Errorf("expected 3 reviews, got %v", finalStats["total_reviews"])
	}
	if finalStats["total_correct"].(float64) != 2 {
		t.Errorf("expected 2 correct, got %v", finalStats["total_correct"])
	}
}

// ═══════════════════════════════════════════
// Init Database Test
// ═══════════════════════════════════════════

func TestE2E_InitDatabase(t *testing.T) {
	router := setupE2E(t)
	token := registerAndGetToken(t, router, "inituser")

	tmpDir := t.TempDir()
	mdContent := "# q\n什么是测试驱动开发？\n# a\nTDD是一种先写测试再写代码的开发方法。\n\n# q\nGo语言的特点是什么？\n# a\n简洁、高效、并发支持好。\n"
	os.WriteFile(filepath.Join(tmpDir, "tdd.md"), []byte(mdContent), 0644)

	// NewQuestionParser reads question_input_linux from CWD.
	// During tests CWD is internal/server/, so create the config there.
	configPath := "question_input_linux"
	origContent, _ := os.ReadFile(configPath)
	os.WriteFile(configPath, []byte(tmpDir), 0644)
	defer func() {
		os.WriteFile(configPath, origContent, 0644)
	}()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/init", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if !resp.Success {
		t.Fatalf("init failed: %s", resp.Error)
	}
	data := resp.Data.(map[string]interface{})
	if data["imported"].(float64) != 2 {
		t.Errorf("expected 2 imported, got %v", data["imported"])
	}
}

// ═══════════════════════════════════════════
// Multi-User Isolation Test
// ═══════════════════════════════════════════

func TestE2E_MultiUserIsolation(t *testing.T) {
	router := setupE2E(t)

	tokenA := registerAndGetToken(t, router, "userA")
	tokenB := registerAndGetToken(t, router, "userB")

	addQuestion(t, router, tokenA, "A的秘密问题", "A的秘密答案")

	// User B should NOT see User A's question
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/due-questions", nil)
	req.Header.Set("Authorization", "Bearer "+tokenB)
	router.ServeHTTP(w, req)

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Success != false {
		t.Error("user B should not see user A's question")
	}

	// User A should see their question
	w2 := httptest.NewRecorder()
	req2 := httptest.NewRequest("GET", "/api/due-questions", nil)
	req2.Header.Set("Authorization", "Bearer "+tokenA)
	router.ServeHTTP(w2, req2)

	var resp2 Response
	json.Unmarshal(w2.Body.Bytes(), &resp2)
	if !resp2.Success {
		t.Fatalf("user A should see their question: %s", resp2.Error)
	}
	questions := resp2.Data.(map[string]interface{})["questions"].([]interface{})
	if len(questions) != 1 {
		t.Errorf("user A should have 1 question, got %d", len(questions))
	}
}

// ═══════════════════════════════════════════
// Helper Function Tests
// ═══════════════════════════════════════════

func TestExtractCategory(t *testing.T) {
	tests := []struct {
		source   string
		expected string
	}{
		{"questions/01_storage/ceph/foo.md", "01_storage"},
		{"questions/03_languages/go/basics.md", "03_languages"},
		{"手动输入", "未分类"},
		{"", "未分类"},
		{"single-file.md", "未分类"},
		{"my_dir/sub/file.md", "my_dir"},
	}

	for _, tt := range tests {
		result := extractCategory(tt.source)
		if result != tt.expected {
			t.Errorf("extractCategory(%q) = %q, want %q", tt.source, result, tt.expected)
		}
	}
}

func TestCategoryLabel(t *testing.T) {
	if label := categoryLabel("01_storage"); label != "存储" {
		t.Errorf("expected '存储', got '%s'", label)
	}
	if label := categoryLabel("nonexistent"); label != "nonexistent" {
		t.Errorf("expected 'nonexistent', got '%s'", label)
	}
}
