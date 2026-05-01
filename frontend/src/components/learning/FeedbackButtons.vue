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
  font-size: 15px;
  font-weight: 500;
  margin-bottom: $spacing-md;
  color: $text-secondary;
  text-align: center;
}

.feedback-buttons {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
}

.btn-feedback {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: $spacing-md $spacing-sm;
  border-radius: $border-radius-md;
  border: 1.5px solid $card-border;
  background: $card-bg;
  cursor: pointer;
  transition: all $transition-normal;
  user-select: none;
  touch-action: manipulation;
  -webkit-tap-highlight-color: transparent;

  &:active {
    transform: scale(0.95);
  }

  &.btn-feedback-1 {
    border-color: rgba($success, 0.4);
    background: rgba($success, 0.05);
    &:active { background: $success; color: $text-white; border-color: $success; }
  }

  &.btn-feedback-2 {
    border-color: rgba($ink-blue, 0.35);
    background: rgba($ink-blue, 0.04);
    &:active { background: $ink-blue; color: $text-white; border-color: $ink-blue; }
  }

  &.btn-feedback-3 {
    border-color: rgba($warning, 0.45);
    background: rgba($warning, 0.06);
    &:active { background: $warning; color: $text-white; border-color: $warning; }
  }

  &.btn-feedback-4 {
    border-color: rgba($danger, 0.35);
    background: rgba($danger, 0.04);
    &:active { background: $danger; color: $text-white; border-color: $danger; }
  }
}

.feedback-emoji {
  font-size: 30px;
  margin-bottom: 4px;
}

.feedback-text {
  font-size: 15px;
  font-weight: 600;
  margin-bottom: 2px;
}

.feedback-desc {
  font-size: 11px;
  opacity: 0.6;
}

@media (max-width: 480px) {
  .feedback-buttons {
    grid-template-columns: 1fr;
  }
  .feedback-emoji {
    font-size: 26px;
  }
}
</style>
