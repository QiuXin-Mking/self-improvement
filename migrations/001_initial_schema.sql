-- Migration 001: Create users and questions tables
-- Table: users
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME
);

-- Index for soft deletes
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at);

-- Index for username lookups
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);

-- Table: questions
CREATE TABLE IF NOT EXISTS questions (
    id TEXT PRIMARY KEY,
    user_id INTEGER NOT NULL,
    question_text TEXT NOT NULL,
    answer_text TEXT NOT NULL,
    source TEXT,
    level INTEGER DEFAULT 4, -- 1-4: 1=proficient, 2=fair, 3=forgotten, 4=completely forgotten
    next_review DATETIME NOT NULL,
    review_count INTEGER DEFAULT 0,
    correct_count INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    last_reviewed DATETIME,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,

    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Index for user questions lookup
CREATE INDEX IF NOT EXISTS idx_questions_user_id ON questions(user_id);

-- Index for next review date lookups (critical for scheduling)
CREATE INDEX IF NOT EXISTS idx_questions_next_review ON questions(next_review);

-- Index for soft deletes
CREATE INDEX IF NOT EXISTS idx_questions_deleted_at ON questions(deleted_at);