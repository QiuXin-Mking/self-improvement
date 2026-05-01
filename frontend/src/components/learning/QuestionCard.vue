<template>
  <div class="question-card card">
    <div class="question-section">
      <div class="question-label">问题</div>
      <div class="question-content">{{ question }}</div>
    </div>

    <div class="action-buttons">
      <button class="btn btn-show-answer" @click="$emit('show-answer')">
        查看答案
      </button>
      <button class="btn btn-delete" @click="$emit('delete')">
        删除问题
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  question: string
}

defineProps<Props>()
defineEmits<{
  (e: 'show-answer'): void
  (e: 'delete'): void
}>()
</script>

<style lang="scss" scoped>
.question-card {
  border-left: 4px solid $ink-blue;
  background: $card-bg;
  border: 1px solid $card-border;
  border-left: 4px solid $ink-blue;
}

.question-section {
  margin-bottom: $spacing-lg;
}

.question-label {
  font-size: 12px;
  font-weight: 600;
  color: $ink-mid;
  text-transform: uppercase;
  letter-spacing: 1px;
  margin-bottom: $spacing-sm;
}

.question-content {
  font-size: 17px;
  line-height: 1.7;
  color: $text-primary;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.action-buttons {
  display: flex;
  gap: $spacing-sm;
  justify-content: center;
}

.btn-show-answer {
  flex: 1;
  background: $ink-blue;
  color: $text-white;
  border: none;
  border-radius: $border-radius-md;
  padding: 12px 22px;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
  transition: all $transition-normal;

  &:hover {
    background: color.scale($ink-blue, $lightness: -8%);
  }
  &:active {
    transform: scale(0.97);
  }
}

.btn-delete {
  background: transparent;
  color: $text-muted;
  border: 1px solid $card-border;
  border-radius: $border-radius-md;
  padding: 12px 18px;
  font-size: 14px;
  cursor: pointer;
  transition: all $transition-normal;

  &:hover {
    color: $danger;
    border-color: rgba($danger, 0.4);
  }
  &:active {
    transform: scale(0.97);
  }
}

@media (max-width: 480px) {
  .action-buttons {
    flex-direction: column;
  }
}
</style>
