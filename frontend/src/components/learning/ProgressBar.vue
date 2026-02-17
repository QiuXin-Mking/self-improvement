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
  height: 40px;
  background: rgba(255, 255, 255, 0.3);
  border-radius: $border-radius-lg;
  overflow: hidden;
  margin-bottom: $spacing-md;
}

.progress-fill {
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  background: linear-gradient(90deg, $primary-gradient-start 0%, $primary-gradient-end 100%);
  transition: width $transition-normal;
  border-radius: $border-radius-lg;
}

.progress-text {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  font-size: 16px;
  font-weight: bold;
  color: $text-white;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
  z-index: 1;
}
</style>
