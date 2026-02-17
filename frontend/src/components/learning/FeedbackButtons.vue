<template>
  <div class="feedback-section">
    <div class="feedback-label">请反馈你的记忆程度:</div>
    <div class="feedback-buttons">
      <button
        v-for="item in feedbackOptions"
        :key="item.value"
        class="btn-feedback"
        :class="`btn-feedback-${item.value}`"
        @click="handleFeedback(item.value)"
      >
        <span class="feedback-emoji">{{ item.emoji }}</span>
        <span class="feedback-text">{{ item.text }}</span>
        <span class="feedback-desc">{{ item.desc }}</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
interface FeedbackOption {
  value: 1 | 2 | 3 | 4
  emoji: string
  text: string
  desc: string
}

const emit = defineEmits<{
  (e: 'feedback', value: 1 | 2 | 3 | 4): void
}>()

const feedbackOptions: FeedbackOption[] = [
  { value: 1, emoji: '✅', text: '熟练', desc: '记得很清楚' },
  { value: 2, emoji: '👍', text: '一般', desc: '记得但不熟练' },
  { value: 3, emoji: '😐', text: '忘记', desc: '忘记了部分内容' },
  { value: 4, emoji: '❌', text: '完全忘记', desc: '完全不记得' }
]

function handleFeedback(value: 1 | 2 | 3 | 4) {
  emit('feedback', value)
}
</script>

<style lang="scss" scoped>
.feedback-section {
  margin-top: $spacing-xl;
}

.feedback-label {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: $spacing-md;
  color: $text-primary;
  text-align: center;
}

.feedback-buttons {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: $spacing-sm;
}

.btn-feedback {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: $spacing-md $spacing-sm;
  border-radius: $border-radius-sm;
  border: 2px solid #ddd;
  background: $bg-white;
  cursor: pointer;
  transition: all $transition-normal;
  user-select: none;
  touch-action: manipulation;
  -webkit-tap-highlight-color: transparent;

  &:active {
    transform: scale(0.95);
  }

  &.btn-feedback-1 {
    border-color: $success-color;

    &:active {
      background: $success-color;
      color: $text-white;
    }
  }

  &.btn-feedback-2 {
    border-color: $warning-color;

    &:active {
      background: $warning-color;
      color: $text-white;
    }
  }

  &.btn-feedback-3 {
    border-color: #fd7e14;

    &:active {
      background: #fd7e14;
      color: $text-white;
    }
  }

  &.btn-feedback-4 {
    border-color: $error-color;

    &:active {
      background: $error-color;
      color: $text-white;
    }
  }
}

.feedback-emoji {
  font-size: 32px;
  margin-bottom: 5px;
}

.feedback-text {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 3px;
}

.feedback-desc {
  font-size: 12px;
  opacity: 0.8;
}

@media (max-width: 480px) {
  .feedback-buttons {
    grid-template-columns: 1fr;
  }

  .feedback-emoji {
    font-size: 28px;
  }
}
</style>
