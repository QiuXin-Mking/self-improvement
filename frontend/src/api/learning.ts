import api from './index'
import type {
  Stats,
  Question,
  DueQuestionsData,
  InitData,
  FeedbackLevel,
  ApiResponse
} from './types'

export const learningApi = {
  // 获取统计信息
  getStats(): Promise<ApiResponse<{ stats: Stats }>> {
    return api.get('/stats')
  },

  // 获取待复习问题
  getDueQuestions(): Promise<ApiResponse<DueQuestionsData>> {
    return api.get('/due-questions')
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
  }
}
