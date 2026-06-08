<template>
  <div class="learning-view">
    <div v-if="isLoading" class="loading-state">
      <van-loading size="48" text-color="#1989fa" />
      <p>正在加载题目...</p>
    </div>

    <template v-else>
      <div class="learning-top-bar">
        <button class="btn-exit" @click="handleExit">✕ 退出</button>
      </div>
      <ProgressBar :progress="progress" :text="progressText" />

      <QuestionCard
        v-if="currentQuestion"
        :question="currentQuestion.question"
        @show-answer="handleShowAnswer"
        @delete="handleDelete"
      />

      <AnswerCard
        v-if="currentQuestion"
        :answer="currentQuestion.answer"
        :visible="isAnswerVisible"
      >
        <FeedbackButtons @feedback="handleFeedback" />
      </AnswerCard>

      <div v-if="!currentQuestion && questions.length === 0" class="empty-state">
        <van-empty description="该分类下没有待复习的问题" />
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import { onMounted, nextTick, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { showDialog, showToast } from 'vant'
import { useLearningStore } from '@/stores/learning'
import { storeToRefs } from 'pinia'
import ProgressBar from '@/components/learning/ProgressBar.vue'
import QuestionCard from '@/components/learning/QuestionCard.vue'
import AnswerCard from '@/components/learning/AnswerCard.vue'
import FeedbackButtons from '@/components/learning/FeedbackButtons.vue'

const router = useRouter()
const route = useRoute()
const isLoading = ref(true)

// 本次复习会话统计
const sessionTotal = ref(0)
const sessionCorrect = ref(0)

const store = useLearningStore()
const { progress, progressText, currentQuestion, isAnswerVisible, questions } = storeToRefs(store)

async function handleExit() {
  const totalQuestions = sessionTotal.value + (currentQuestion.value ? 1 : 0)
  const completed = sessionTotal.value
  const accuracy = completed > 0
    ? Math.round((sessionCorrect.value / completed) * 100)
    : 0

  try {
    await showDialog({
      title: '退出复习',
      message: `已完成 ${completed} 题（正确率 ${accuracy}%），进度已自动保存${
        totalQuestions > completed ? `，剩余 ${totalQuestions - completed} 题下次再复习` : ''
      }`,
      confirmButtonText: '退出',
      cancelButtonText: '继续复习',
    })
  } catch {
    return // 用户点了取消
  }

  router.push(`/dashboard?accuracy=${accuracy}&completed=${completed}`)
}

function handleShowAnswer() {
  store.showAnswer()

  // 延迟滚动到答案位置
  nextTick(() => {
    setTimeout(() => {
      window.scrollTo({
        top: document.documentElement.scrollHeight,
        behavior: 'smooth'
      })
    }, 100)
  })
}

async function handleFeedback(feedback: 1 | 2 | 3 | 4) {
  // 追踪本次会话统计
  sessionTotal.value++
  if (feedback <= 2) {
    sessionCorrect.value++
  }

  const success = await store.submitFeedback(feedback)

  if (success) {
    // 滚动到顶部
    window.scrollTo({
      top: 0,
      behavior: 'smooth'
    })

    // 检查是否所有问题已完成
    if (!currentQuestion.value) {
      const accuracy = sessionTotal.value > 0
        ? Math.round((sessionCorrect.value / sessionTotal.value) * 100)
        : 0
      // 直接跳转看板（看板上的总结横幅即为完成反馈，无需 toast 避免白框闪烁）
      router.push(`/dashboard?accuracy=${accuracy}&completed=${sessionTotal.value}`)
    }
  }
}

async function handleDelete() {
  const confirmed = await showDialog({
    title: '确认删除',
    message: '确定要删除这个问题吗？',
    showCancelButton: true,
    confirmButtonText: '删除',
    cancelButtonText: '取消'
  })
    .then(() => true)
    .catch(() => false)

  if (confirmed) {
    const success = await store.deleteCurrentQuestion()

    if (success) {
      showToast({
        type: 'success',
        message: '删除成功',
        duration: 1500
      })

      // 如果没有剩余问题，返回 dashboard
      if (questions.value.length === 0) {
        setTimeout(() => router.push('/dashboard'), 500)
      }
    }
  }
}

onMounted(async () => {
  window.scrollTo({ top: 0 })

  const categoriesParam = route.query.categories as string | undefined
  store.activeCategories = categoriesParam || ''

  const result = await store.fetchDueQuestions(categoriesParam)
  isLoading.value = false

  // 记录本次会话的初始题目数
  sessionTotal.value = 0
  sessionCorrect.value = 0

  if (result && result.needsInit) {
    showToast({
      message: '知识库为空，请先初始化知识库',
      duration: 2000
    })
    setTimeout(() => router.push('/dashboard'), 500)
  }
})
</script>

<style lang="scss" scoped>
.learning-view {
  width: 100%;
  max-width: 600px;
  margin: 0 auto;
  padding: 0 $spacing-md;
  padding-top: calc($spacing-sm + var(--safe-area-top));
  padding-bottom: calc($spacing-xl + var(--safe-area-bottom));
}

.learning-top-bar {
  display: flex;
  justify-content: flex-end;
  padding: 8px 0;
}

.btn-exit {
  background: none;
  border: 1px solid #e0e0e0;
  color: #999;
  font-size: 13px;
  padding: 4px 12px;
  border-radius: 14px;
  cursor: pointer;
  font-family: inherit;
  transition: all 0.2s;

  &:hover {
    color: #666;
    border-color: #bbb;
  }
}

.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 80px 20px;
  color: #666;
  p {
    margin-top: 16px;
    font-size: 14px;
  }
}

.empty-state {
  padding: 80px 20px;
}
</style>
