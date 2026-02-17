<template>
  <div v-if="visible" class="answer-card card">
    <div class="answer-section">
      <div class="answer-label">答案</div>
      <div class="answer-content">{{ answer }}</div>
    </div>
    <slot></slot>
  </div>
</template>

<script setup lang="ts">
interface Props {
  answer: string
  visible: boolean
}

defineProps<Props>()
</script>

<style lang="scss" scoped>
.answer-card {
  border-left: 4px solid $success-color;
}

.answer-section {
  margin-bottom: $spacing-lg;
}

.answer-label {
  font-size: 14px;
  font-weight: bold;
  color: $text-secondary;
  margin-bottom: $spacing-sm;
}

.answer-content {
  font-size: 16px;
  line-height: 1.6;
  color: $text-primary;
  white-space: pre-wrap;
  word-wrap: break-word;
}
</style>
