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
  border-left: 4px solid $info-color;
}

.question-section {
  margin-bottom: $spacing-lg;
}

.question-label {
  font-size: 14px;
  font-weight: bold;
  color: $text-secondary;
  margin-bottom: $spacing-sm;
}

.question-content {
  font-size: 16px;
  line-height: 1.6;
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
  background: $success-color;
  color: $text-white;

  &:hover {
    background: color.scale($success-color, $lightness: -10%);
  }
}

.btn-delete {
  background: $error-color;
  color: $text-white;

  &:hover {
    background: color.scale($error-color, $lightness: -10%);
  }
}

@media (max-width: 480px) {
  .action-buttons {
    flex-direction: column;
  }
}
</style>
