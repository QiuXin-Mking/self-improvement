package server

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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

// AddQuestionRequest represents the request body for adding a question manually
type AddQuestionRequest struct {
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

var db *gorm.DB
var sr *spacedrepetition.SpacedRepetition

// SetupRouter initializes the Gin engine with all routes configured.
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin,Content-Type,Content-Length,Accept,Authorization,X-Requested-With")
		c.Header("Access-Control-Max-Age", "43200")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	public := r.Group("/api")
	{
		public.POST("/register", registerHandler)
		public.POST("/login", loginHandler)
	}

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", profileHandler)
		protected.GET("/stats", getStatsHandler)
		protected.GET("/categories", getCategoriesHandler)
		protected.GET("/due-questions", getDueQuestionsHandler)
		protected.POST("/update-review", updateReviewHandler)
		protected.POST("/delete-question", deleteQuestionHandler)
		protected.POST("/init", initDatabaseHandler)
		protected.POST("/upload-zip", uploadZipHandler)
		protected.POST("/upload-md", uploadMdHandler)
		protected.POST("/add-question", addQuestionHandler)
		protected.POST("/reset-demo", resetDemoHandler)
		protected.GET("/forecast", getForecastHandler)
	}

	staticDir := os.Getenv("STATIC_DIR")
	if staticDir == "" {
		staticDir = "dist"
	}
	r.Static("/assets", filepath.Join(staticDir, "assets"))
	r.StaticFile("/vite.svg", filepath.Join(staticDir, "vite.svg"))

	r.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(staticDir, "index.html"))
	})

	return r
}

// Run initializes the database and starts the web server.
func Run() {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "data/app.db"
	}

	os.MkdirAll("data", 0755)

	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.User{}, &models.Question{})
	if err != nil {
		panic("failed to migrate database")
	}

	// Composite index for due-questions query: WHERE user_id + ORDER BY next_review
	db.Exec("CREATE INDEX IF NOT EXISTS idx_questions_user_next_review ON questions(user_id, next_review)")

	db.Model(&models.Question{}).Where("category = '' OR category IS NULL").Update("category", "未分类")
	var questions []models.Question
	db.Where("source LIKE ? AND category = ?", "questions/%", "未分类").Find(&questions)
	for _, q := range questions {
		cat := extractCategory(q.Source)
		if cat != "未分类" {
			db.Model(&q).Update("category", cat)
		}
	}

	sr = spacedrepetition.NewSpacedRepetition(db)

	// 自动创建 demo 体验账户
	seedDemoUser(db, sr)

	r := SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "4430"
	}

	r.Run("0.0.0.0:" + port)
}

// InitTestDB initializes an in-memory database for testing.
func InitTestDB(database *gorm.DB) {
	db = database
	sr = spacedrepetition.NewSpacedRepetition(db)
}

func registerHandler(c *gin.Context) {
	type RegisterRequest struct {
		Username string `json:"username" binding:"required,min=3,max=32"`
		Password string `json:"password" binding:"required,min=6,max=128"`
	}

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "Invalid request format"})
		return
	}

	var existingUser models.User
	if err := db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, Response{Success: false, Error: "Username already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "Failed to hash password"})
		return
	}

	user := models.User{Username: req.Username, Password: string(hashedPassword)}
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "Failed to create user"})
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "Failed to generate token"})
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
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "Invalid request format"})
		return
	}

	var user models.User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, Response{Success: false, Error: "用户名或密码错误"})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "Database error occurred"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "用户名或密码错误"})
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "Failed to generate token"})
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
		Data:    map[string]interface{}{"user_id": userId, "username": username},
	})
}

func getStatsHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	stats, err := sr.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "Failed to get stats"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: map[string]interface{}{"stats": stats}})
}

func getCategoriesHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	categories, err := sr.GetCategories(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "Failed to get categories"})
		return
	}

	for _, cat := range categories {
		cat["label"] = categoryLabel(cat["name"].(string))
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: map[string]interface{}{"categories": categories}})
}

func getDueQuestionsHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	category := c.Query("category")
	if category == "" {
		category = c.Query("categories")
	}

	var dueQuestions []*models.Question
	var err error
	if category != "" {
		categories := splitCategories(category)
		dueQuestions, err = sr.GetDueQuestionsByCategories(userID, categories)
	} else {
		dueQuestions, err = sr.GetDueQuestions(userID)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "Failed to get due questions"})
		return
	}

	if len(dueQuestions) == 0 {
		if category != "" {
			c.JSON(http.StatusOK, Response{
				Success: true,
				Data:    map[string]interface{}{"questions": []interface{}{}, "message": "该分类下没有待复习的问题！"},
			})
			return
		}

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
			Data:    map[string]interface{}{"questions": []interface{}{}, "message": "太棒了！今天没有需要复习的问题！"},
		})
		return
	}

	var questionsData []map[string]interface{}
	for _, q := range dueQuestions {
		questionsData = append(questionsData, map[string]interface{}{
			"id": q.ID, "question": q.QuestionText, "answer": q.AnswerText,
			"review_count": q.ReviewCount, "correct_count": q.CorrectCount,
			"source": q.Source, "category": q.Category,
		})
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    map[string]interface{}{"questions": questionsData, "total": len(questionsData)},
	})
}

func updateReviewHandler(c *gin.Context) {
	var req UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "Invalid request format"})
		return
	}

	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	if err := sr.UpdateReview(userID, req.QuestionID, req.Feedback); err != nil {
		c.JSON(http.StatusNotFound, Response{Success: false, Error: "Question not found or update failed"})
		return
	}

	stats, err := sr.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "Failed to get updated stats"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: map[string]interface{}{"stats": stats}})
}

func deleteQuestionHandler(c *gin.Context) {
	var req DeleteQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "Invalid request format"})
		return
	}

	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	if err := sr.DeleteQuestion(userID, req.QuestionID); err != nil {
		c.JSON(http.StatusNotFound, Response{Success: false, Error: "Question not found"})
		return
	}

	stats, err := sr.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "Failed to get updated stats"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: map[string]interface{}{"stats": stats}})
}

func extractCategory(source string) string {
	if source == "" || source == "手动输入" {
		return "未分类"
	}
	source = filepath.ToSlash(source)
	parts := strings.Split(source, "/")
	for i, p := range parts {
		if p == "questions" && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	if len(parts) >= 2 {
		return parts[0]
	}
	return "未分类"
}

var categoryLabelMap = map[string]string{
	"00_summaries": "总结",
	"01_storage":   "存储",
	"01_调试命令集合.md": "调试命令集",
	"03_languages": "编程语言",
	"05_problems":  "问题案例",
	"06_career":    "职场修炼",
	"08_tools":     "工具",
	"09_ai":        "AI",
	"10_billing":   "计费",
}

func splitCategories(raw string) []string {
	parts := strings.Split(raw, ",")
	var result []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}

func categoryLabel(name string) string {
	if label, ok := categoryLabelMap[name]; ok {
		return label
	}
	return name
}

func importQuestions(userID uint, questions []*parser.Question) (imported, skipped, duplicates int, err error) {
	seenQuestions := make(map[string]bool)
	var uniqueQuestions []*parser.Question

	for _, q := range questions {
		if !seenQuestions[q.QuestionText] {
			seenQuestions[q.QuestionText] = true
			uniqueQuestions = append(uniqueQuestions, q)
		} else {
			duplicates++
		}
	}

	for _, q := range uniqueQuestions {
		qID := fmt.Sprintf("q_%d_%s", userID, spacedrepetition.Hash(q.QuestionText))

		var existingQuestion models.Question
		err := db.Where("user_id = ? AND question_text = ?", userID, q.QuestionText).First(&existingQuestion).Error
		if err == nil {
			skipped++
			continue
		}

		if err := sr.AddQuestion(userID, qID, q.QuestionText, q.AnswerText, q.SourceFile, extractCategory(q.SourceFile)); err != nil {
			continue
		}
		imported++
	}

	return imported, skipped, duplicates, nil
}

func initDatabaseHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID := userId.(uint)
	username, _ := c.Get("username")

	// 每个用户有自己独立的题目目录，避免扫描到其他用户的内容
	userDir := filepath.Join("questions", username.(string))
	p, err := parser.NewQuestionParser(userDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: err.Error()})
		return
	}

	questions, err := p.ParseAllFiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: err.Error()})
		return
	}

	if len(questions) == 0 {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "没有找到任何问题！请先将 .md 文件放入 questions/" + username.(string) + "/ 目录，或使用上传功能添加题目。"})
		return
	}

	imported, skipped, duplicates, err := importQuestions(userID, questions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "导入问题失败"})
		return
	}

	stats, err := sr.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "Failed to get updated stats"})
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

func uploadZipHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "请上传文件"})
		return
	}

	if filepath.Ext(file.Filename) != ".zip" {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "只支持 .zip 格式的压缩文件"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "无法打开上传的文件"})
		return
	}
	defer src.Close()

	zipBytes, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "读取文件失败"})
		return
	}

	zipReader, err := zip.NewReader(bytes.NewReader(zipBytes), file.Size)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "无法解析 zip 文件，请确保文件未损坏"})
		return
	}

	tempDir, err := os.MkdirTemp("", "spaced-repetition-zip-*")
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "服务器内部错误"})
		return
	}
	defer os.RemoveAll(tempDir)

	for _, f := range zipReader.File {
		if f.FileInfo().IsDir() {
			continue
		}
		dstPath := filepath.Join(tempDir, f.Name)
		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			continue
		}
		dst, err := os.Create(dstPath)
		if err != nil {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			dst.Close()
			continue
		}
		io.Copy(dst, rc)
		rc.Close()
		dst.Close()
	}

	p, err := parser.NewQuestionParser(tempDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "处理文件失败"})
		return
	}

	questions, err := p.ParseAllFiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "解析文件失败"})
		return
	}

	if len(questions) == 0 {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "zip 文件中没有找到有效的问题！请确保包含格式正确的 .md 文件。"})
		return
	}

	imported, skipped, duplicates, err := importQuestions(userID, questions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "导入问题失败"})
		return
	}

	stats, err := sr.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "获取统计失败"})
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

func uploadMdHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "请上传文件"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".md" && ext != ".markdown" {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "只支持 .md 或 .markdown 格式的文件"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "无法打开上传的文件"})
		return
	}
	defer src.Close()

	content, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "读取文件失败"})
		return
	}

	p, err := parser.NewQuestionParser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "处理文件失败"})
		return
	}

	questions := p.ParseContent(string(content), file.Filename)
	if len(questions) == 0 {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "文件中没有找到有效的问题！请使用 # q 和 # a 标记格式。"})
		return
	}

	imported, skipped, duplicates, err := importQuestions(userID, questions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "导入问题失败"})
		return
	}

	stats, err := sr.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "获取统计失败"})
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

func addQuestionHandler(c *gin.Context) {
	var req AddQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "请提供问题和答案"})
		return
	}

	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	questionText := strings.TrimSpace(req.Question)
	answerText := strings.TrimSpace(req.Answer)

	if questionText == "" || answerText == "" {
		c.JSON(http.StatusBadRequest, Response{Success: false, Error: "问题和答案不能为空"})
		return
	}

	var existingQuestion models.Question
	if err := db.Where("user_id = ? AND question_text = ?", userID, questionText).First(&existingQuestion).Error; err == nil {
		c.JSON(http.StatusConflict, Response{Success: false, Error: "该问题已存在"})
		return
	}

	qID := fmt.Sprintf("q_%d_%s", userID, spacedrepetition.Hash(questionText))
	if err := sr.AddQuestion(userID, qID, questionText, answerText, "手动输入", "未分类"); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "添加问题失败"})
		return
	}

	stats, err := sr.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "获取统计失败"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Message: "问题添加成功", Data: map[string]interface{}{"stats": stats}})
}

func resetDemoHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	if err := sr.ResetUserQuestions(userID); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "重置体验数据失败"})
		return
	}

	stats, err := sr.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "获取统计失败"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Message: "体验数据已重置", Data: map[string]interface{}{"stats": stats}})
}

func getForecastHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID := userId.(uint)

	daysStr := c.DefaultQuery("days", "7")
	days, err := strconv.Atoi(daysStr)
	if err != nil || days < 1 || days > 30 {
		days = 7
	}

	forecast, err := sr.GetForecast(userID, days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Error: "获取复习预告失败"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: map[string]interface{}{"forecast": forecast}})
}

// seedDemoUser creates a demo account with sample questions so first-time
// visitors can explore KnowLoop instantly without registering.
func seedDemoUser(database *gorm.DB, srInstance *spacedrepetition.SpacedRepetition) {
	// Check if demo user already exists
	var demoUser models.User
	result := database.Where("username = ?", "demo").First(&demoUser)

	if result.Error != nil {
		// Create demo user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("demo123"), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		demoUser = models.User{
			Username: "demo",
			Password: string(hashedPassword),
		}
		if err := database.Create(&demoUser).Error; err != nil {
			return
		}
	}

	// Check if demo questions are already seeded
	var count int64
	database.Model(&models.Question{}).Where("user_id = ?", demoUser.ID).Count(&count)
	if count > 0 {
		return
	}

	// Sample questions about learning methods to showcase the system
	demoQuestions := []struct{ question, answer string }{
		{
			"什么是艾宾浩斯遗忘曲线？",
			"由德国心理学家赫尔曼·艾宾浩斯于1885年提出，描述了人类记忆随时间的遗忘规律：学习后20分钟遗忘约42%，1小时后遗忘约56%，1天后遗忘约74%，一周后遗忘约77%。遗忘速度呈现「先快后慢」的特点，这为间隔重复学习法提供了科学依据。",
		},
		{
			"什么是间隔重复（Spaced Repetition）？",
			"一种科学的学习技巧，通过在即将遗忘的关键时间点进行复习，最大化记忆保留效果。相比传统的死记硬背，间隔重复的效率可提升2-5倍。Anki、SuperMemo以及本系统KnowLoop都基于这一原理。",
		},
		{
			"番茄工作法的核心是什么？",
			"由Francesco Cirillo在1980年代创立：25分钟专注工作 + 5分钟短暂休息 = 1个番茄钟。每完成4个番茄钟后休息15-30分钟。核心理念是通过短时间的高强度专注来克服拖延、提升效率。",
		},
		{
			"费曼学习法是什么？",
			"诺贝尔物理学奖得主理查德·费曼提出：用教别人的方式来检验自己是否真正理解。具体步骤：①选择一个概念、②尝试用最简单的语言解释它、③发现解释不通的地方回去学习、④简化并类比。如果你不能简单解释清楚，说明你还没真正学会。",
		},
		{
			"什么是主动回忆（Active Recall）？",
			"不看书本和笔记，主动从大脑中提取信息的学习方法。研究表明，主动回忆比被动重读的效率高50%以上。具体做法：读完一页书后合上书本，尝试回忆主要内容、关键概念和逻辑关系。这正是KnowLoop的核心学习方式。",
		},
		{
			"SQ3R阅读法包含哪五个步骤？",
			"Survey（浏览）：快速浏览标题、摘要、图表，建立整体印象；Question（提问）：将标题转化为问题；Read（阅读）：带着问题精读内容；Recite（复述）：用自己的话总结要点；Review（复习）：间隔时间后回顾巩固。这是一种系统性的主动阅读方法。",
		},
		{
			"刻意练习（Deliberate Practice）的核心要素是什么？",
			"由心理学家Anders Ericsson提出：①明确具体的目标、②高度的专注和投入、③获得即时反馈、④持续走出舒适区、⑤建立高质量的心理表征。刻意练习不是简单的10000小时重复，而是有目的、有反馈的针对性训练。",
		},
		{
			"Dunning-Kruger效应（达克效应）是什么？",
			"1999年由David Dunning和Justin Kruger提出：能力较低的人往往高估自己的能力水平，而真正有能力的人反而容易低估自己。原因在于：缺乏能力的人也缺乏识别自己不足所需的元认知能力。真正的智慧始于「知道自己不知道」。",
		},
	}

	for _, q := range demoQuestions {
		qID := fmt.Sprintf("q_%d_%s", demoUser.ID, spacedrepetition.Hash(q.question))
		_ = srInstance.AddQuestion(demoUser.ID, qID, q.question, q.answer, "demo", "学习方法")
	}
}
