// API 类型定义

export interface Question {
  id: string
  question: string
  answer: string
  review_count: number
  correct_count: number
  source: string
}

export interface Stats {
  total_questions: number
  due_questions: number
  total_reviews: number
  total_correct: number
  accuracy: number
}

export interface ApiResponse<T = any> {
  success: boolean
  data?: T
  error?: string
  message?: string
  needs_init?: boolean
}

export interface DueQuestionsData {
  questions: Question[]
  total: number
}

export interface InitData {
  message: string
  imported: number
  skipped: number
  duplicates: number
  stats: Stats
}

export type FeedbackLevel = 1 | 2 | 3 | 4

// Auth types
export interface LoginRequest {
  username: string
  password: string
}

export interface RegisterRequest {
  username: string
  password: string
}

export interface LoginResponse {
  token: string
  user_id: number
  username: string
}

export interface ProfileResponse {
  user_id: number
  username: string
}
