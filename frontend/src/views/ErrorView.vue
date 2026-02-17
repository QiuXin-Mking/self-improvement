<template>
  <div class="error-view">
    <div class="card error-card">
      <h2>❌ 出错了</h2>
      <p>{{ errorMessage || '发生了未知错误' }}</p>
      <button class="btn btn-primary" @click="handleReload">
        刷新页面
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  errorMessage?: string
}

defineProps<Props>()

function handleReload() {
  window.location.reload()
}
</script>

<style lang="scss" scoped>
.error-view {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 300px;
}

.error-card {
  text-align: center;
  max-width: 500px;
  width: 100%;
  border-left: 4px solid $error-color;
}

h2 {
  font-size: 24px;
  color: $error-color;
  margin-bottom: $spacing-md;
}

p {
  font-size: 16px;
  color: $text-secondary;
  margin-bottom: $spacing-lg;
}
</style>
