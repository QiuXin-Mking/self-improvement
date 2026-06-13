<template>
  <div v-if="visible" class="answer-card card">
    <div class="answer-section">
      <div class="answer-label">答案</div>
      <div class="answer-content markdown-body" v-html="rendered"></div>
    </div>
    <slot></slot>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { renderMarkdown } from '@/composables/useMarkdown'

interface Props {
  answer: string
  visible: boolean
}

const props = defineProps<Props>()
const rendered = computed(() => renderMarkdown(props.answer))
</script>

<style lang="scss" scoped>
.answer-card {
  background: $card-bg;
  border: 1px solid $card-border;
  border-left: 4px solid $amber;
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
}
</style>
