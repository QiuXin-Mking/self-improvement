<template>
  <div class="no-questions-view">
    <div class="card">
      <h2>🎉 太棒了！</h2>
      <p>今天没有需要复习的问题！</p>
      <div class="button-group">
        <button class="btn btn-primary" @click="$emit('refresh')">
          刷新统计
        </button>
        <button class="btn btn-secondary" @click="$emit('reinit')">
          重新初始化
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineEmits<{
  (e: 'refresh'): void
  (e: 'reinit'): void
}>()
</script>

<style lang="scss" scoped>
.no-questions-view {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 300px;
}

.card {
  text-align: center;
  max-width: 500px;
  width: 100%;
}

h2 {
  font-size: 32px;
  margin-bottom: $spacing-md;
}

p {
  font-size: 18px;
  color: $text-secondary;
  margin-bottom: $spacing-lg;
}

.button-group {
  display: flex;
  gap: $spacing-sm;
  justify-content: center;
}

@media (max-width: 480px) {
  .button-group {
    flex-direction: column;
  }
}
</style>
