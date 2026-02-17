package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"self-improvement/internal/middleware"
	"self-improvement/internal/models"
	"self-improvement/internal/parser"
	"self-improvement/internal/spacedrepetition"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// UpdateReviewRequest represents the request body for updating review
type UpdateReviewRequest struct {
	QuestionID string `json:"question_id" binding:"required"`
	Feedback   int    `json:"feedback" binding:"required,min=1,max=4"`
}

// DeleteQuestionRequest represents the request body for deleting a question
type DeleteQuestionRequest struct {
	QuestionID string `json:"question_id" binding:"required"`
}

var db *gorm.DB
var sr *spacedrepetition.SpacedRepetition

func main() {
	// Initialize database
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "data/app.db"
	}

	// Create data directory if it doesn't exist
	os.MkdirAll("data", 0755)

	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate schema
	err = db.AutoMigrate(&models.User{}, &models.Question{})
	if err != nil {
		panic("failed to migrate database")
	}

	// Initialize spaced repetition with database connection
	sr = spacedrepetition.NewSpacedRepetition(db)

	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// Public routes
	public := r.Group("/api")
	{
		public.POST("/register", registerHandler)
		public.POST("/login", loginHandler)
	}

	// Protected routes (require authentication)
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", profileHandler)
		protected.GET("/stats", getStatsHandler)
		protected.GET("/due-questions", getDueQuestionsHandler)
		protected.POST("/update-review", updateReviewHandler)
		protected.POST("/delete-question", deleteQuestionHandler)
		protected.POST("/init", initDatabaseHandler)
	}

	// Get port from environment, default to 5000
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r.Run(":" + port)
}

func registerHandler(c *gin.Context) {
	type RegisterRequest struct {
		Username string `json:"username" binding:"required,min=3,max=32"`
		Password string `json:"password" binding:"required,min=6,max=128"`
	}

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request format",
		})
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, Response{
			Success: false,
			Error:   "Username already exists",
		})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to hash password",
		})
		return
	}

	// Create the user
	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to create user",
		})
		return
	}

	// Generate token
	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    map[string]interface{}{"token": token, "user_id": user.ID, "username": user.Username},
		Message: "Registration successful",
	})
}

func loginHandler(c *gin.Context) {
	type LoginRequest struct {
		Username string `json:"username" binding:"required,min=3,max=32"`
		Password string `json:"password" binding:"required,min=6,max=128"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request format",
		})
		return
	}

	// Find user by username
	var user models.User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, Response{
				Success: false,
				Error:   "Invalid username or password",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Database error occurred",
		})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, Response{
			Success: false,
			Error:   "Invalid username or password",
		})
		return
	}

	// Generate token
	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    map[string]interface{}{"token": token, "user_id": user.ID, "username": user.Username},
		Message: "Login successful",
	})
}

func profileHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data: map[string]interface{}{
			"user_id":  userId,
			"username": username,
		},
	})
}

func getStatsHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	stats, err := sr.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get stats",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    map[string]interface{}{"stats": stats},
	})
}

func getDueQuestionsHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	dueQuestions, err := sr.GetDueQuestions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get due questions",
		})
		return
	}

	totalQuestions := len(dueQuestions)
	if totalQuestions == 0 {
		// Check if user has any questions at all
		var count int64
		db.Model(&models.Question{}).Where("user_id = ?", userID).Count(&count)

		if count == 0 {
			c.JSON(http.StatusOK, Response{
				Success: false,
				Error:   "知识库为空！请先初始化知识库。",
				Data:    map[string]bool{"needs_init": true},
			})
			return
		}
	}

	if len(dueQuestions) == 0 {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data: map[string]interface{}{
				"questions": []interface{}{},
				"message":   "太棒了！今天没有需要复习的问题！",
			},
		})
		return
	}

	// Format question data
	var questionsData []map[string]interface{}
	for _, q := range dueQuestions {
		questionsData = append(questionsData, map[string]interface{}{
			"id":            q.ID,
			"question":      q.QuestionText,
			"answer":        q.AnswerText,
			"review_count":  q.ReviewCount,
			"correct_count": q.CorrectCount,
			"source":        q.Source,
		})
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data: map[string]interface{}{
			"questions": questionsData,
			"total":     len(questionsData),
		},
	})
}

func updateReviewHandler(c *gin.Context) {
	var req UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request format",
		})
		return
	}

	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	err := sr.UpdateReview(userID, req.QuestionID, req.Feedback)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   "Question not found or update failed",
		})
		return
	}

	stats, err := sr.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get updated stats",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    map[string]interface{}{"stats": stats},
	})
}

func deleteQuestionHandler(c *gin.Context) {
	var req DeleteQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request format",
		})
		return
	}

	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	err := sr.DeleteQuestion(userID, req.QuestionID)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   "Question not found",
		})
		return
	}

	stats, err := sr.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get updated stats",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    map[string]interface{}{"stats": stats},
	})
}

func initDatabaseHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	p, err := parser.NewQuestionParser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	questions, err := p.ParseAllFiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	if len(questions) == 0 {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "没有找到任何问题！请确保配置的目录下有 .md 文件。",
		})
		return
	}

	// Deduplicate based on question text
	seenQuestions := make(map[string]bool)
	var uniqueQuestions []*parser.Question
	duplicates := 0

	for _, q := range questions {
		normalizedQ := q.QuestionText
		if !seenQuestions[normalizedQ] {
			seenQuestions[normalizedQ] = true
			uniqueQuestions = append(uniqueQuestions, q)
		} else {
			duplicates++
		}
	}

	// Process questions and associate with current user
	imported := 0
	skipped := 0

	for _, q := range uniqueQuestions {
		// Include user ID in the hash to ensure uniqueness across users
		qID := fmt.Sprintf("q_%d_%s", userID, spacedrepetition.Hash(q.QuestionText))

		// Check if this question already exists for this user
		var existingQuestion models.Question
		err := db.Where("user_id = ? AND question_text = ?", userID, q.QuestionText).First(&existingQuestion).Error

		if err == nil {
			// Question already exists for this user
			skipped++
			continue
		}

		// Add the question for this user
		err = sr.AddQuestion(userID, qID, q.QuestionText, q.AnswerText, q.SourceFile)
		if err != nil {
			// Log error but continue processing
			continue
		}
		imported++
	}

	stats, err := sr.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get updated stats",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data: map[string]interface{}{
			"message":    "成功导入 " + strconv.Itoa(imported) + " 个新问题到知识库！",
			"imported":   imported,
			"skipped":    skipped,
			"duplicates": duplicates,
			"stats":      stats,
		},
	})
}