<template>
  <div class="learning-view">
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
  </div>
</template>

<script setup lang="ts">
import { onMounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { showDialog, showToast } from 'vant'
import { useLearningStore } from '@/stores/learning'
import { storeToRefs } from 'pinia'
import ProgressBar from '@/components/learning/ProgressBar.vue'
import QuestionCard from '@/components/learning/QuestionCard.vue'
import AnswerCard from '@/components/learning/AnswerCard.vue'
import FeedbackButtons from '@/components/learning/FeedbackButtons.vue'

const router = useRouter()

const store = useLearningStore()
const { progress, progressText, currentQuestion, isAnswerVisible, questions, currentQuestionIndex } = storeToRefs(store)

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
  const success = await store.submitFeedback(feedback)

  if (success) {
    // 滚动到顶部
    window.scrollTo({
      top: 0,
      behavior: 'smooth'
    })

    // 检查是否所有问题已完成（通过 currentQuestion 是否为 null 判断）
    if (!currentQuestion.value) {
      showToast({
        type: 'success',
        message: '本轮复习完成！',
        duration: 2000
      })
      setTimeout(() => router.push('/dashboard'), 500)
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
  // 滚动到顶部
  window.scrollTo({ top: 0 })

  // 获取待复习问题
  const result = await store.fetchDueQuestions()

  // 如果没有问题且需要初始化
  if (result && result.needsInit) {
    showToast({
      type: 'warning',
      message: '知识库为空，请先初始化知识库',
      duration: 2000
    })
    // 返回 dashboard
    setTimeout(() => router.push('/dashboard'), 500)
  }
})
</script>

<style lang="scss" scoped>
.learning-view {
  width: 100%;
  max-width: 600px;
  margin: 0 auto;
}
</style>
