# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a spaced repetition learning system based on the Ebbinghaus Forgetting Curve implemented in Go. It helps users review knowledge scientifically through both CLI and web interfaces. The system parses Markdown question files, tracks learning progress, and schedules reviews based on memory retention algorithms.

## Commands

### CLI Application (main.go)

```bash
# Initialize knowledge base (parse .md files from configured directories)
go run main.go --init

# Start training session (default command)
go run main.go

# Build and run
go build -o train main.go
./train

# View learning statistics
go run main.go --stats
```

### Web Application (web_server.go)

```bash
# Run web server (default port 5000)
go run web_server.go

# Build and run with custom port
PORT=8080 go run web_server.go
```

### Building and Deployment

```bash
# Install dependencies
go mod tidy

# Build standalone executables
go build -o bin/train main.go
go build -o bin/web_app web_server.go

# Using Makefile
make build          # Build all binaries
make run-cli        # Run CLI app
make run-web        # Run web server
make init           # Initialize knowledge base
```

## Architecture

### Core Components

**1. Question Parser (internal/parser/parser.go)**
- Parses Markdown files with `# q` (question) and `# a` (answer) markers
- Supports multiple question directories configured via `question_input` (Windows) or `question_input_linux` (Linux)
- Recursively scans directories with filepath.Walk
- Returns slice of structs with fields: `QuestionText`, `AnswerText`, `SourceFile`

**2. Spaced Repetition Algorithm (internal/spacedrepetition/spaced_repetition.go)**
- Implements Ebbinghaus forgetting curve with 4 memory levels
- Data stored in `data/learning_data.json` with structure:
  ```
  {
    "questions": {
      "q_<hash>": {
        "id": string,
        "question": string,
        "answer": string,
        "level": 1-4,
        "next_review": ISO datetime,
        "review_count": int,
        "correct_count": int,
        "created_at": ISO datetime,
        "last_reviewed": ISO datetime
      }
    }
  }
  ```
- Review intervals by feedback level (hours):
  - Level 1 (熟练/Proficient): 168h (7 days)
  - Level 2 (一般/Fair): 72h (3 days)
  - Level 3 (忘记/Forgotten): 24h (1 day)
  - Level 4 (完全忘记/Completely Forgotten): 2h
- Interval multipliers for correct answers: 2.5x, 1.8x, 1.3x, 1.0x respectively
- Additional 1.2x multiplier when accuracy > 80%

**3. CLI Interface (main.go)**
- Uses ANSI escape codes for colored terminal output
- Interactive training loop with question display, answer reveal, and feedback collection
- Supports question deletion during training (low-quality questions)
- Question ID format: `q_{hash(question.strip())}` for deduplication

**4. Web Interface (web_server.go)**
- Gin REST API with endpoints:
  - `GET /api/stats` - learning statistics
  - `GET /api/due-questions` - questions needing review
  - `POST /api/update-review` - submit feedback (1-4)
  - `POST /api/delete-question` - remove question
  - `POST /api/init` - initialize database
- Mobile-friendly HTML template in `templates/index.html`
- Static assets in `static/css/` and `static/js/`

### Question File Format

Markdown files must follow this structure:
```markdown
# q
What is the question?

# a
This is the answer.

# q
Another question?

# a
Another answer.
```

### Configuration Files

**question_input (Linux)** / **question_input_windows (Windows)**
- Plain text file with one directory path per line
- Defaults to `questions/` if file doesn't exist or is empty

### Data Flow

1. **Initialization**: `parser.ParseAllFiles()` → deduplication → `sr.AddQuestion()` → save to JSON
2. **Training**: `sr.GetDueQuestions()` → filter by `next_review <= now` → sort by oldest first
3. **Review Update**: feedback (1-4) → calculate interval with multipliers → update `next_review` and stats
4. **Question ID**: Hash-based IDs ensure same question always gets same ID (prevents duplicates)

### Important Implementation Details

- **Deduplication**: Questions are compared using `strings.TrimSpace()` to handle whitespace variations
- **Review Scheduling**: Questions sorted by `next_review` ASC (oldest first priority)
- **Level Adjustment**: Level decreases (improves) after 3+ correct reviews, increases on forgetting
- **Path Handling**: Platform-specific path separators handled in QuestionParser initialization
- **Encoding**: All files use UTF-8 encoding explicitly

### Deployment Options

1. **Go Build**: Creates standalone executables in `bin/`
2. **Direct**: Run Go programs with proper modules setup

## Testing

No automated test suite currently exists. Manual testing involves:
- Creating .md files in configured directories
- Running `--init` to verify parsing
- Testing training workflow with various feedback inputs
- Verifying JSON data structure in `data/learning_data.json`

## Development Reminders

**IMPORTANT**: During every agent call/development session, ensure to verify that `http://localhost:3000/login` is functioning correctly. This is a critical endpoint that must remain operational after any changes.