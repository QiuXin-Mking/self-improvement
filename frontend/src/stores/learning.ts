import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { learningApi } from '@/api/learning'
import type { Question, Stats } from '@/api/types'

export const useLearningStore = defineStore('learning', () => {
  // 状态
  const stats = ref<Stats>({
    total_questions: 0,
    due_questions: 0,
    total_reviews: 0,
    total_correct: 0,
    accuracy: 0
  })

  const questions = ref<Question[]>([])
  const currentQuestionIndex = ref(0)
  const isAnswerVisible = ref(false)

  // 计算属性
  const currentQuestion = computed(() => {
    return questions.value[currentQuestionIndex.value] || null
  })

  const progress = computed(() => {
    if (questions.value.length === 0) return 0
    return Math.min((currentQuestionIndex.value + 1) / questions.value.length, 1) * 100
  })

  const progressText = computed(() => {
    if (questions.value.length === 0) return '0/0'
    return `${Math.min(currentQuestionIndex.value + 1, questions.value.length)}/${questions.value.length}`
  })

  // Actions
  async function fetchStats() {
    try {
      const response = await learningApi.getStats()
      if (response.success && response.data && response.data.stats) {
        stats.value = response.data.stats
      }
    } catch (error) {
      console.error('获取统计失败:', error)
    }
  }

  async function fetchDueQuestions() {
    try {
      const response = await learningApi.getDueQuestions()
      if (response.success && response.data) {
        questions.value = response.data.questions || []
        currentQuestionIndex.value = 0
        isAnswerVisible.value = false
        return {
          count: questions.value.length,
          needsInit: false
        }
      }
      if (response.needs_init) {
        return {
          count: 0,
          needsInit: true
        }
      }
      return {
        count: 0,
        needsInit: false
      }
    } catch (error) {
      console.error('获取问题失败:', error)
      throw error
    }
  }

  async function submitFeedback(feedback: 1 | 2 | 3 | 4) {
    if (!currentQuestion.value) return false

    try {
      const response = await learningApi.updateReview(
        currentQuestion.value.id,
        feedback
      )

      if (response.success && response.data && response.data.stats) {
        stats.value = response.data.stats
        nextQuestion()
        return true
      }
      return false
    } catch (error) {
      console.error('提交反馈失败:', error)
      return false
    }
  }

  async function deleteCurrentQuestion() {
    if (!currentQuestion.value) return false

    try {
      const response = await learningApi.deleteQuestion(currentQuestion.value.id)

      if (response.success && response.data && response.data.stats) {
        stats.value = response.data.stats
        questions.value.splice(currentQuestionIndex.value, 1)

        // 调整索引
        if (currentQuestionIndex.value >= questions.value.length) {
          currentQuestionIndex.value = Math.max(0, questions.value.length - 1)
        }

        return true
      }
      return false
    } catch (error) {
      console.error('删除问题失败:', error)
      return false
    }
  }

  async function initDatabase() {
    try {
      const response = await learningApi.initDatabase()
      if (response.success && response.data) {
        stats.value = response.data.stats
        return response.data
      }
      return null
    } catch (error) {
      console.error('初始化失败:', error)
      throw error
    }
  }

  function showAnswer() {
    isAnswerVisible.value = true
  }

  function nextQuestion() {
    currentQuestionIndex.value++
    isAnswerVisible.value = false
  }

  function reset() {
    questions.value = []
    currentQuestionIndex.value = 0
    isAnswerVisible.value = false
  }

  return {
    // State
    stats,
    questions,
    currentQuestionIndex,
    isAnswerVisible,

    // Computed
    currentQuestion,
    progress,
    progressText,

    // Actions
    fetchStats,
    fetchDueQuestions,
    submitFeedback,
    deleteCurrentQuestion,
    initDatabase,
    showAnswer,
    nextQuestion,
    reset
  }
})
