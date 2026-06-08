import api from './index'
import type {
  Stats,
  DueQuestionsData,
  CategoriesData,
  InitData,
  ImportResult,
  FeedbackLevel,
  ForecastDay,
  ApiResponse
} from './types'

export const learningApi = {
  // 获取统计信息
  getStats(): Promise<ApiResponse<{ stats: Stats }>> {
    return api.get('/stats')
  },

  // 获取分类列表
  getCategories(): Promise<ApiResponse<CategoriesData>> {
    return api.get('/categories')
  },

  // 获取待复习问题
  getDueQuestions(category?: string): Promise<ApiResponse<DueQuestionsData>> {
    const params = category ? { category } : {}
    return api.get('/due-questions', { params })
  },

  // 提交复习反馈
  updateReview(questionId: string, feedback: FeedbackLevel): Promise<ApiResponse<{ stats: Stats }>> {
    return api.post('/update-review', {
      question_id: questionId,
      feedback
    })
  },

  // 删除问题
  deleteQuestion(questionId: string): Promise<ApiResponse<{ stats: Stats }>> {
    return api.post('/delete-question', {
      question_id: questionId
    })
  },

  // 初始化知识库
  initDatabase(): Promise<ApiResponse<InitData>> {
    return api.post('/init')
  },

  // 上传 zip 文件导入知识库
  uploadZip(file: File): Promise<ApiResponse<ImportResult>> {
    const formData = new FormData()
    formData.append('file', file)
    return api.post('/upload-zip', formData, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
  },

  // 上传单个 .md 文件导入
  uploadMd(file: File): Promise<ApiResponse<ImportResult>> {
    const formData = new FormData()
    formData.append('file', file)
    return api.post('/upload-md', formData, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
  },

  // 重置体验数据（重置所有复习状态）
  resetDemo(): Promise<ApiResponse<{ stats: Stats }>> {
    return api.post('/reset-demo')
  },

  // 获取未来复习预告
  getForecast(days: number = 7): Promise<ApiResponse<{ forecast: ForecastDay[] }>> {
    return api.get('/forecast', { params: { days } })
  },

  // 手动添加问题
  addQuestion(question: string, answer: string): Promise<ApiResponse<{ stats: Stats }>> {
    return api.post('/add-question', { question, answer })
  }
}
