<template>
  <div class="progress-bar">
    <div class="progress-fill" :style="{ width: `${progress}%` }"></div>
    <span class="progress-text">{{ text }}</span>
  </div>
</template>

<script setup lang="ts">
interface Props {
  progress: number
  text: string
}

defineProps<Props>()
</script>

<style lang="scss" scoped>
.progress-bar {
  position: relative;
  height: 44px;
  background: $card-bg;
  border: 1px solid $card-border;
  border-radius: $border-radius-md;
  overflow: hidden;
  margin-bottom: $spacing-md;
  box-shadow: $shadow-xs;
}

.progress-fill {
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  background: linear-gradient(90deg, $ink-blue 0%, $ink-mid 100%);
  transition: width $transition-normal;
  border-radius: $border-radius-md;
}

.progress-text {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  font-size: 14px;
  font-weight: 600;
  color: $text-white;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.25);
  z-index: 1;
}
</style>
