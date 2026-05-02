# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Spaced repetition learning system — user knowledge points with Ebbinghaus-forgetting-curve-based review scheduling. Two interfaces: CLI (`cli_server.go`) and web (Go backend + Vue 3 frontend). Multi-tenant with JWT auth and SQLite storage.

## Commands

```bash
# Backend
go run web_server.go                          # Start web API (port 5000)
PORT=8080 JWT_SECRET=xxx go run web_server.go # Custom port/secret
go run cli_server.go --init                   # CLI: init knowledge base from .md files
go run cli_server.go                          # CLI: start training session
go run cli_server.go --stats                  # CLI: view stats

# Frontend
cd frontend && npm run dev                    # Vite dev server (port 3000, proxies /api → :8000)
cd frontend && npm run build                  # Production build → frontend/dist

# Full stack (start.sh)
./scripts/start.sh                            # Start backend (8000) + frontend (3000)
./scripts/stop.sh                             # Stop both services

# Build
./scripts/compile.sh                          # Build Go binary + frontend → bin/ + frontend/dist
BUILD_TARGET=linux ./scripts/compile.sh       # Cross-compile for Linux

# Deploy
./scripts/deploy_to_huoshan.sh                # Deploy to huoshan server

# Makefile targets
make run-web                                  # go run web_server.go with JWT_SECRET
make run-cli                                  # go run cli_server.go
make build                                    # Build CLI + web binaries to bin/
make test                                     # go test -v ./...
make deps                                     # go mod tidy && cd frontend && npm install
```

## Architecture

### Backend (`web_server.go`)
- **Framework**: Gin with CORS, grouped routes (`/api` public, `/api` protected via JWT middleware)
- **ORM**: GORM with SQLite (`data/app.db`), auto-migrates `User` and `Question` models
- **Auth**: JWT (HS256, 24h expiry), `internal/middleware/auth.go` — `GenerateToken()` + `AuthMiddleware()`
- **API routes**: register, login (public); profile, stats, due-questions, update-review, delete-question, init, upload-zip (protected)

### Spaced Repetition (`internal/spacedrepetition/spaced_repetition.go`)
- Feedback levels: 1=proficient (7d), 2=fair (3d), 3=forgotten (1d), 4=completely forgotten (2h)
- Interval multipliers: 2.5x / 1.8x / 1.3x / 1.0x, with extra 1.2x when accuracy > 80%
- Level improves (decreases) after 3+ correct reviews, degrades on forgotten
- Question ID: `q_` + SHA-256 hash of trimmed question text

### Parser (`internal/parser/parser.go`)
- Recursively scans configured directories for `.md` files
- Parses `# q` / `# a` markers — question text between markers, answer text after `# a`
- Directories configured via `question_input` file (one path per line, defaults to `questions/`)

### Models (`internal/models/`)
- `User`: ID, Username (unique), Password (bcrypt hash, `json:"-"`), has-many Questions
- `Question`: ID (hash-based string PK), UserID (FK), QuestionText, AnswerText, Source, Level (1-4), NextReview, ReviewCount, CorrectCount, soft-delete

### Frontend (Vue 3 + TypeScript + Vite)
- **Stack**: Vue 3, Vant UI (mobile-first), Pinia, Vue Router, Axios
- **Routes**: `/login`, `/register`, `/dashboard`, `/learn` (all lazy-loaded), `/` → dashboard
- **Stores**: `auth` (token in localStorage, user profile), `learning` (stats, questions queue, review flow)
- **API layer**: Axios instance with Bearer token interceptor and error toast notifications
- **Vite config**: port 3000, `/api` proxied to `localhost:8000`, SCSS with `@use` variables injection, Vant auto-import

### Data flow
1. **Init**: Parser scans .md files → deduplication by question hash → bulk insert into SQLite
2. **Review**: `GET /due-questions` (next_review <= now, sorted ASC) → user gives feedback 1-4 → `POST /update-review` recalculates next_review with interval multipliers
3. **Stats**: `GET /stats` returns total, due, total_reviews, accuracy per user

### Key conventions
- Question format in .md: `# q` line, then question body, then `# a` line, then answer body
- JWT_SECRET env var is required; defaults to `my-secret-key` in dev scripts
- Frontend dev on :3000, backend on :8000 (start.sh) or :5000 (standalone `go run`)
- Database auto-created at `data/app.db` on first run
