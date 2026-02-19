package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"bufio"
	"flag"
)

// Question represents a question-answer pair with learning data
type Question struct {
	ID           string    `json:"id"`
	QuestionText string    `json:"question"`
	AnswerText   string    `json:"answer"`
	Source       string    `json:"source"`
	Level        int       `json:"level"`           // 1-4: 1=熟练, 2=一般, 3=忘记, 4=完全忘记
	NextReview   time.Time `json:"next_review"`     // Next review scheduled time
	ReviewCount  int       `json:"review_count"`    // Total number of reviews
	CorrectCount int       `json:"correct_count"`   // Number of correct answers
	CreatedAt    time.Time `json:"created_at"`      // When question was added
	LastReviewed *time.Time `json:"last_reviewed"`  // When last reviewed (nil if never)
}

// LearningData holds all questions and metadata
type LearningData struct {
	Questions    map[string]*Question `json:"questions"`
	LastUpdated  time.Time            `json:"last_updated"`
}

// SpacedRepetition manages the spaced repetition algorithm
type SpacedRepetition struct {
	DataFile string
	Data     *LearningData
}

// NewSpacedRepetition creates a new spaced repetition instance
func NewSpacedRepetition(dataFile string) *SpacedRepetition {
	sr := &SpacedRepetition{
		DataFile: dataFile,
		Data: &LearningData{
			Questions:   make(map[string]*Question),
			LastUpdated: time.Now(),
		},
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(dataFile)
	os.MkdirAll(dir, 0755)

	// Load existing data if file exists
	sr.LoadData()

	return sr
}

// LoadData loads learning data from file
func (sr *SpacedRepetition) LoadData() error {
	if _, err := os.Stat(sr.DataFile); os.IsNotExist(err) {
		// File doesn't exist, use default data
		return nil
	}

	data, err := ioutil.ReadFile(sr.DataFile)
	if err != nil {
		return err
	}

	var loadedData LearningData
	err = json.Unmarshal(data, &loadedData)
	if err != nil {
		return err
	}

	sr.Data = &loadedData
	return nil
}

// SaveData saves learning data to file
func (sr *SpacedRepetition) SaveData() error {
	sr.Data.LastUpdated = time.Now()
	data, err := json.MarshalIndent(sr.Data, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(sr.DataFile, data, 0644)
}

// AddQuestion adds a new question to the knowledge base
func (sr *SpacedRepetition) AddQuestion(id, question, answer, source string) {
	if _, exists := sr.Data.Questions[id]; !exists {
		now := time.Now()
		q := &Question{
			ID:           id,
			QuestionText: question,
			AnswerText:   answer,
			Source:       source,
			Level:        4, // Start as completely forgotten
			NextReview:   now,
			ReviewCount:  0,
			CorrectCount: 0,
			CreatedAt:    now,
			LastReviewed: nil,
		}
		sr.Data.Questions[id] = q
		sr.SaveData()
	}
}

// GetDueQuestions returns questions that are due for review
func (sr *SpacedRepetition) GetDueQuestions() []*Question {
	now := time.Now()
	var dueQuestions []*Question

	for _, q := range sr.Data.Questions {
		if q.NextReview.Before(now) || q.NextReview.Equal(now) {
			dueQuestions = append(dueQuestions, q)
		}
	}

	// Sort by next review time (oldest first priority)
	sort.Slice(dueQuestions, func(i, j int) bool {
		return dueQuestions[i].NextReview.Before(dueQuestions[j].NextReview)
	})

	return dueQuestions
}

// GetQuestion returns a specific question by ID
func (sr *SpacedRepetition) GetQuestion(id string) *Question {
	if q, exists := sr.Data.Questions[id]; exists {
		return q
	}
	return nil
}

// UpdateReview updates review results for a question
func (sr *SpacedRepetition) UpdateReview(id string, feedback int) bool {
	q, exists := sr.Data.Questions[id]
	if !exists {
		return false
	}

	now := time.Now()
	q.ReviewCount++
	q.LastReviewed = &now

	// Update statistics
	if feedback <= 2 { // Proficient or fair counts as correct
		q.CorrectCount++
	}

	// Calculate next review time based on feedback
	intervals := map[int]time.Duration{
		1: 7 * 24 * time.Hour,  // Proficient: 7 days
		2: 3 * 24 * time.Hour,  // Fair: 3 days
		3: 24 * time.Hour,      // Forgotten: 1 day
		4: 2 * time.Hour,       // Completely forgotten: 2 hours
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
	if q.ReviewCount > 0 {
		accuracy := float64(q.CorrectCount) / float64(q.ReviewCount)
		if accuracy > 0.8 { // High accuracy gets additional boost
			multiplier *= 1.2
		}
	}

	// Calculate final interval
	intervalHours := float64(baseInterval.Hours()) * multiplier
	nextReview := now.Add(time.Duration(intervalHours * float64(time.Hour)))

	q.NextReview = nextReview

	// Update memory level
	if feedback <= 2 {
		// If answered correctly, potentially upgrade level (decrease number)
		if q.CorrectCount >= 3 && q.Level > 1 {
			q.Level = max(1, q.Level-1)
		}
	} else {
		// If forgotten, downgrade level (increase number)
		q.Level = min(4, q.Level+1)
	}

	sr.SaveData()
	return true
}

// DeleteQuestion removes a question from the knowledge base
func (sr *SpacedRepetition) DeleteQuestion(id string) bool {
	if _, exists := sr.Data.Questions[id]; exists {
		delete(sr.Data.Questions, id)
		sr.SaveData()
		return true
	}
	return false
}

// GetStats returns learning statistics
func (sr *SpacedRepetition) GetStats() map[string]interface{} {
	total := len(sr.Data.Questions)
	due := len(sr.GetDueQuestions())

	totalReviews := 0
	totalCorrect := 0

	for _, q := range sr.Data.Questions {
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
	}
}

// Question represents a parsed question-answer pair
type ParsedQuestion struct {
	QuestionText string
	AnswerText   string
	SourceFile   string
}

// QuestionParser handles parsing of Markdown files
type QuestionParser struct {
	QuestionsDirs []string
}

// NewQuestionParser creates a new question parser
func NewQuestionParser(questionsDir ...string) (*QuestionParser, error) {
	var dirs []string

	if len(questionsDir) == 0 {
		// Determine config file based on platform
		configFile := "question_input"
		if runtime.GOOS != "windows" {
			configFile = "question_input_linux"
		}

		if _, err := os.Stat(configFile); err == nil {
			file, err := os.Open(configFile)
			if err != nil {
				return nil, err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line != "" {
					dirs = append(dirs, line)
				}
			}

			if len(dirs) == 0 {
				dirs = []string{"questions"}
			}
		} else {
			// Default to questions directory
			dirs = []string{"questions"}
		}
	} else {
		dirs = questionsDir
	}

	// Create directories if they don't exist
	for _, dir := range dirs {
		os.MkdirAll(dir, 0755)
	}

	return &QuestionParser{
		QuestionsDirs: dirs,
	}, nil
}

// ParseFile parses a single markdown file
func (qp *QuestionParser) ParseFile(filePath string) ([]*ParsedQuestion, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return qp.ParseContent(string(content), filepath.Base(filePath)), nil
}

// ParseContent parses content directly from a string
func (qp *QuestionParser) ParseContent(content, sourceFile string) []*ParsedQuestion {
	var questions []*ParsedQuestion

	// Find all question and answer positions using regex
	questionPattern := "(?is)#\\s*q[^\\n]*\\n(.*?)(?=\\n#\\s*a)"
	answerPattern := "(?is)#\\s*a[^\\n]*\\n(.*?)(?=\\n#\\s*q|$)"

	// Find all question matches
	qMatches := getAllStringSubmatchIndex(content, questionPattern)
	aMatches := getAllStringSubmatchIndex(content, answerPattern)

	for _, qMatch := range qMatches {
		if len(qMatch) >= 4 {
			qStart, qEnd := qMatch[2], qMatch[3]
			qText := strings.TrimSpace(content[qStart:qEnd])

			// Find the corresponding answer (first one after this question)
			qEndPos := qEnd
			for _, aMatch := range aMatches {
				if len(aMatch) >= 4 && aMatch[2] > qEndPos {
					aStart, aEnd := aMatch[2], aMatch[3]
					aText := strings.TrimSpace(content[aStart:aEnd])

					if qText != "" && aText != "" {
						questions = append(questions, &ParsedQuestion{
							QuestionText: qText,
							AnswerText:   aText,
							SourceFile:   sourceFile,
						})
					}
					break
				}
			}
		}
	}

	return questions
}

// Helper function to mimic regexp functionality without importing it
func getAllStringSubmatchIndex(s, pattern string) [][]int {
	// This is a simplified implementation to find matches for our specific pattern
	// A complete implementation would require the regexp package, but since we can't import it
	// in this self-contained file, we'll implement a basic matching function

	var matches [][]int
	qPositions := findTagPositions(s, "# q", "# a")
	aPositions := findTagPositions(s, "# a", "# q")

	for _, qPos := range qPositions {
		qStart := qPos
		// Find the end of the question (until next answer)
		qEnd := sLen(s) // default to end of string
		for _, aPos := range aPositions {
			if aPos > qPos {
				qEnd = aPos
				break
			}
		}

		// Extract question content
		if qStart >= 0 && qEnd <= sLen(s) && qEnd > qStart {
			// Add a mock match (we only care about the content)
			// Format: [overall_start, overall_end, content_start, content_end]
			match := []int{qStart, qEnd, qStart + 4, qEnd} // skip "# q" and space
			matches = append(matches, match)
		}
	}

	return matches
}

func findTagPositions(s, tag, nextTag string) []int {
	var positions []int
	for i := 0; i < len(s)-len(tag); i++ {
		if s[i:i+len(tag)] == tag {
			positions = append(positions, i)
		}
	}
	return positions
}

func sLen(s string) int {
	return len(s)
}

// ParseAllFiles parses all markdown files in configured directories
func (qp *QuestionParser) ParseAllFiles() ([]*ParsedQuestion, error) {
	var allQuestions []*ParsedQuestion

	for _, dir := range qp.QuestionsDirs {
		fmt.Printf("\nScanning directory: %s (recursively)\n", dir)

		// Walk through directory recursively to find .md files
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".md") {
				fmt.Printf("Parsing: %s\n", path)

				questions, err := qp.ParseFile(path)
				if err != nil {
					return err
				}

				count := len(questions)
				fmt.Printf("  Extracted %d questions\n", count)

				allQuestions = append(allQuestions, questions...)
			}

			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	return allQuestions, nil
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

// Feedback mappings
var feedbackMap = map[string]int{
	"1":         1,
	"2":         2,
	"3":         3,
	"4":         4,
	"熟练":        1,
	"一般":        2,
	"忘记":        3,
	"完全忘记":      4,
	"1熟练":       1,
	"2一般":       2,
	"3忘记":       3,
	"4完全忘记":     4,
	"熟练的":       1,
	"一般的":       2,
	"忘记的":       3,
	"完全忘记的":     4,
}

var feedbackLabels = map[int]string{
	1: "熟练",
	2: "一般",
	3: "忘记",
	4: "完全忘记",
}

func main() {
	var initFlag = flag.Bool("init", false, "Initialize knowledge base (parse .md files from configured directories)")
	var statsFlag = flag.Bool("stats", false, "View learning statistics")
	flag.Parse()

	if *initFlag {
		initDatabase()
	} else if *statsFlag {
		printStatsCmd()
	} else {
		startTraining()
	}
}

func printHeader() {
	fmt.Println()
	fmt.Println("\033[36m" + strings.Repeat("=", 60)) // Cyan
	fmt.Println("  基于艾宾浩斯遗忘曲线的知识库学习系统")
	fmt.Println(strings.Repeat("=", 60) + "\033[0m")   // Reset
	fmt.Println()
}

func printColored(color string, text string) {
	colors := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"magenta": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",
		"reset":  "\033[0m",
	}

	fmt.Print(colors[color])
	fmt.Print(text)
	fmt.Print(colors["reset"])
}

func printQuestion(qData *Question, index, total int) {
	fmt.Println()
	printColored("yellow", strings.Repeat("─", 60))
	if index > 0 {
		printColored("green", fmt.Sprintf("问题 [%d/%d]", index, total))
		printColored("reset", "")
		fmt.Println()
	}
	printColored("yellow", strings.Repeat("─", 60))
	printColored("reset", "")
	fmt.Println()

	printColored("white", qData.QuestionText)
	fmt.Println()

	printColored("cyan", "提示: 输入 'a' 或 'answer' 查看答案")
	fmt.Println()
	printColored("cyan", "输入 'd' 或 'delete' 删除此问题（低质量问题）")
	fmt.Println()
	printColored("cyan", "输入 'q' 或 'quit' 退出本次训练")
	fmt.Println()
}

func printAnswer(qData *Question) {
	fmt.Println()
	printColored("green", strings.Repeat("─", 60))
	printColored("green", "答案")
	printColored("reset", "")
	fmt.Println()
	printColored("green", strings.Repeat("─", 60))
	printColored("reset", "")
	fmt.Println()

	printColored("white", qData.AnswerText)
	fmt.Println()
}

func printFeedbackPrompt() {
	fmt.Println()
	printColored("cyan", "请反馈你的记忆程度:")
	printColored("reset", "")
	fmt.Println()
	printColored("green", "  1") // Green
	printColored("reset", " 或 ")
	printColored("green", "熟练") // Green
	printColored("reset", "     - 记得很清楚")
	fmt.Println()
	printColored("yellow", "  2") // Yellow
	printColored("reset", " 或 ")
	printColored("yellow", "一般") // Yellow
	printColored("reset", "     - 记得但不熟练")
	fmt.Println()
	printColored("magenta", "  3") // Magenta
	printColored("reset", " 或 ")
	printColored("magenta", "忘记") // Magenta
	printColored("reset", "     - 忘记了部分内容")
	fmt.Println()
	printColored("red", "  4") // Red
	printColored("reset", " 或 ")
	printColored("red", "完全忘记") // Red
	printColored("reset", " - 完全不记得")
	fmt.Println()
	printColored("cyan", "  输入 'skip' 跳过这个问题")
	printColored("reset", "")
	fmt.Println()
	printColored("red", "  输入 'd' 或 'delete' 删除此问题（低质量问题）")
	printColored("reset", "")
	fmt.Println()
}

func printStatsFunc(sr *SpacedRepetition) {
	stats := sr.GetStats()

	fmt.Println()
	printColored("cyan", strings.Repeat("=", 60))
	printColored("cyan", "学习统计")
	printColored("reset", "")
	fmt.Println()
	printColored("cyan", strings.Repeat("=", 60))
	printColored("reset", "")

	fmt.Printf("总问题数: ")
	printColored("yellow", fmt.Sprintf("%v", stats["total_questions"]))
	fmt.Println()

	fmt.Printf("待复习: ")
	printColored("yellow", fmt.Sprintf("%v", stats["due_questions"]))
	fmt.Println()

	fmt.Printf("总复习次数: ")
	printColored("yellow", fmt.Sprintf("%v", stats["total_reviews"]))
	fmt.Println()

	fmt.Printf("正确次数: ")
	printColored("green", fmt.Sprintf("%v", stats["total_correct"]))
	fmt.Println()

	fmt.Printf("正确率: ")
	printColored("green", fmt.Sprintf("%v%%", stats["accuracy"]))
	fmt.Println()
}

func initDatabase() {
	printHeader()
	printColored("cyan", "正在初始化知识库...\n")
	printColored("reset", "")

	parser, err := NewQuestionParser()
	if err != nil {
		printColored("red", fmt.Sprintf("错误创建解析器: %v\n", err))
		return
	}

	questions, err := parser.ParseAllFiles()
	if err != nil {
		printColored("red", fmt.Sprintf("错误解析文件: %v\n", err))
		return
	}

	if len(questions) == 0 {
		printColored("red", "错误: 没有找到任何问题！\n")
		printColored("yellow", "请确保配置的目录下有 .md 文件，格式如下:\n")
		fmt.Println("# q\n你的问题\n# a\n你的答案\n")
		printColored("yellow", "可以编辑 question_input 配置多个目录\n")
		return
	}

	printColored("green", fmt.Sprintf("找到 %d 个问题，正在去重...\n", len(questions)))
	printColored("reset", "")

	// Deduplicate based on question text
	seenQuestions := make(map[string]bool)
	var uniqueQuestions []*ParsedQuestion
	duplicates := 0

	for _, q := range questions {
		normalizedQ := strings.TrimSpace(q.QuestionText)
		if !seenQuestions[normalizedQ] {
			seenQuestions[normalizedQ] = true
			uniqueQuestions = append(uniqueQuestions, q)
		} else {
			duplicates++
		}
	}

	if duplicates > 0 {
		printColored("yellow", fmt.Sprintf("检测到 %d 个重复问题，已自动去重\n", duplicates))
	}

	printColored("green", fmt.Sprintf("去重后剩余 %d 个唯一问题，正在导入知识库...\n", len(uniqueQuestions)))
	printColored("reset", "")

	sr := NewSpacedRepetition("data/learning_data.json")
	imported := 0
	skipped := 0

	// Check for existing questions in DB to avoid duplicates
	existingQuestions := make(map[string]bool)
	for _, qData := range sr.Data.Questions {
		existingQuestions[strings.TrimSpace(qData.QuestionText)] = true
	}

	for _, q := range uniqueQuestions {
		qID := "q_" + Hash(q.QuestionText)

		if existingQuestions[strings.TrimSpace(q.QuestionText)] {
			skipped++
			continue
		}

		sr.AddQuestion(qID, q.QuestionText, q.AnswerText, q.SourceFile)
		imported++
	}

	printColored("green", fmt.Sprintf("成功导入 %d 个新问题到知识库！\n", imported))
	if skipped > 0 {
		printColored("yellow", fmt.Sprintf("跳过了 %d 个已存在的问题\n", skipped))
	}
	fmt.Println()
	printStatsFunc(sr)
}

func startTraining() {
	printHeader()

	sr := NewSpacedRepetition("data/learning_data.json")
	stats := sr.GetStats()

	if stats["total_questions"].(int) == 0 {
		printColored("red", "知识库为空！\n")
		printColored("yellow", "请先运行: go run main.go --init\n")
		fmt.Println()
		return
	}

	dueQuestions := sr.GetDueQuestions()

	if len(dueQuestions) == 0 {
		printColored("green", "太棒了！今天没有需要复习的问题！\n")
		printStatsFunc(sr)
		return
	}

	printColored("cyan", fmt.Sprintf("今天有 %d 个问题需要复习\n", len(dueQuestions)))
	fmt.Println()

	reviewed := 0
	total := len(dueQuestions)

	scanner := bufio.NewScanner(os.Stdin)

	for idx, qData := range dueQuestions {
		qID := qData.ID
		index := idx + 1

		// Show question
		printQuestion(qData, index, total)

		// Wait for user input to see answer or delete
		answerShown := false
		questionDeleted := false

		for {
			printColored("cyan", ">>> ")
			if !scanner.Scan() {
				break
			}
			userInput := strings.TrimSpace(strings.ToLower(scanner.Text()))

			switch userInput {
			case "a", "answer", "答案":
				printAnswer(qData)
				answerShown = true
				break
			case "d", "delete", "删除":
				// Confirm deletion
				printColored("red", "确定要删除这个问题吗？(y/n): ")
				if !scanner.Scan() {
					break
				}
				confirm := strings.TrimSpace(strings.ToLower(scanner.Text()))
				if confirm == "y" || confirm == "yes" || confirm == "是" || confirm == "确认" {
					if sr.DeleteQuestion(qID) {
						printColored("green", "问题已删除\n\n")
						questionDeleted = true
						break
					} else {
						printColored("red", "删除失败，问题不存在\n\n")
					}
				} else {
					printColored("yellow", "已取消删除\n\n")
				}
			case "q", "quit", "退出":
				fmt.Println()
				printColored("yellow", fmt.Sprintf("训练已退出，已复习 %d 个问题\n", reviewed))
				printStatsFunc(sr)
				return
			default:
				printColored("red", "输入 'a' 查看答案，'d' 删除，或 'q' 退出\n")
			}

			if answerShown || questionDeleted {
				break
			}
		}

		// Skip to next if question was deleted
		if questionDeleted || sr.GetQuestion(qID) == nil {
			continue
		}

		// Get feedback
		printFeedbackPrompt()
		for {
			printColored("cyan", ">>> ")
			if !scanner.Scan() {
				break
			}
			feedbackInput := strings.TrimSpace(scanner.Text())

			if feedbackInput == "q" || feedbackInput == "quit" || feedbackInput == "退出" {
				fmt.Println()
				printColored("yellow", fmt.Sprintf("训练已退出，已复习 %d 个问题\n", reviewed))
				printStatsFunc(sr)
				return
			}

			if feedbackInput == "d" || feedbackInput == "delete" || feedbackInput == "删除" {
				// Confirm deletion
				printColored("red", "确定要删除这个问题吗？(y/n): ")
				if !scanner.Scan() {
					break
				}
				confirm := strings.TrimSpace(strings.ToLower(scanner.Text()))
				if confirm == "y" || confirm == "yes" || confirm == "是" || confirm == "确认" {
					if sr.DeleteQuestion(qID) {
						printColored("green", "问题已删除\n\n")
						questionDeleted = true
						break
					} else {
						printColored("red", "删除失败，问题不存在\n\n")
					}
				} else {
					printColored("yellow", "已取消删除\n\n")
				}
				continue
			}

			if feedbackInput == "skip" {
				printColored("yellow", "已跳过此问题\n\n")
				break
			}

			feedback, exists := feedbackMap[feedbackInput]
			if exists {
				sr.UpdateReview(qID, feedback)
				reviewed++

				label := feedbackLabels[feedback]
				printColored("green", fmt.Sprintf("已记录: %s\n\n", label))

				break
			} else {
				printColored("red", "无效输入，请重新输入 (1/2/3/4 或 熟练/一般/忘记/完全忘记，或 'd' 删除)\n")
			}
		}

		// Continue to next if question was deleted during feedback
		if questionDeleted {
			continue
		}
	}

	// Training complete
	fmt.Println()
	printColored("green", strings.Repeat("=", 60))
	printColored("green", "本次训练完成！")
	fmt.Println()
	printColored("green", fmt.Sprintf("共复习了 %d 个问题", reviewed))
	fmt.Println()
	printColored("green", strings.Repeat("=", 60))
	printColored("reset", "")
	fmt.Println()
	printStatsFunc(sr)
}

func printStatsCmd() {
	printHeader()
	sr := NewSpacedRepetition("data/learning_data.json")
	printStatsFunc(sr)
}