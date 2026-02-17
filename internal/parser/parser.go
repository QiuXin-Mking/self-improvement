package parser

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"self-improvement/internal/models"
	"self-improvement/internal/spacedrepetition"
)

// Question represents a parsed question-answer pair
type Question struct {
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
func (qp *QuestionParser) ParseFile(filePath string) ([]*Question, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return qp.ParseContent(string(content), filePath), nil
}

// ParseContent parses content directly from a string
func (qp *QuestionParser) ParseContent(content, sourceFile string) []*Question {
	var questions []*Question

	// Find question and answer markers using simple string splitting
	// This approach is more compatible with Go's regexp limitations

	// Find all positions of # q and # a markers
	lines := strings.Split(content, "\n")

	var qPositions []int
	var aPositions []int

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "# q" || strings.HasPrefix(trimmed, "# q ") {
			qPositions = append(qPositions, i)
		} else if trimmed == "# a" || strings.HasPrefix(trimmed, "# a ") {
			aPositions = append(aPositions, i)
		}
	}

	// Match questions with answers
	for i, qPos := range qPositions {
		if i < len(aPositions) {
			aPos := aPositions[i]

			// Extract question text (from qPos+1 to aPos)
			qText := strings.Join(lines[qPos+1:aPos], "\n")
			qText = strings.TrimSpace(qText)

			// Extract answer text (from aPos+1 to next question or end)
			nextQPos := len(lines)
			if i+1 < len(qPositions) {
				nextQPos = qPositions[i+1]
			}
			aText := strings.Join(lines[aPos+1:nextQPos], "\n")
			aText = strings.TrimSpace(aText)

			if qText != "" && aText != "" {
				questions = append(questions, &Question{
					QuestionText: qText,
					AnswerText:   aText,
					SourceFile:   sourceFile,
				})
			}
		}
	}

	return questions
}

// ParseAllFiles parses all markdown files in configured directories
func (qp *QuestionParser) ParseAllFiles() ([]*Question, error) {
	var allQuestions []*Question

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

// ConvertToModel converts parser questions to model questions for a specific user
func (qp *QuestionParser) ConvertToModel(userID uint, parserQuestion *Question) *models.Question {
	qID := "q_" + spacedrepetition.Hash(parserQuestion.QuestionText)

	modelQuestion := &models.Question{
		ID:           qID,
		UserID:       userID,
		QuestionText: parserQuestion.QuestionText,
		AnswerText:   parserQuestion.AnswerText,
		Source:       parserQuestion.SourceFile,
		Level:        4, // Start as completely forgotten
	}

	return modelQuestion
}

// ConvertToModels converts parser questions to model questions for a specific user
func (qp *QuestionParser) ConvertToModels(userID uint, parserQuestions []*Question) []*models.Question {
	var modelQuestions []*models.Question

	for _, pq := range parserQuestions {
		modelQuestion := qp.ConvertToModel(userID, pq)
		modelQuestions = append(modelQuestions, modelQuestion)
	}

	return modelQuestions
}