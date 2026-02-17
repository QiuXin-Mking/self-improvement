<template>
  <header class="app-header">
    <h1 class="header-title">📚 知识库学习系统</h1>
    <div class="stats-bar">
      <div class="stat-item">
        <span class="stat-label">待复习:</span>
        <span class="stat-value">{{ store.stats.due_questions }}</span>
      </div>
      <div class="stat-item">
        <span class="stat-label">总问题:</span>
        <span class="stat-value">{{ store.stats.total_questions }}</span>
      </div>
      <div class="stat-item">
        <span class="stat-label">正确率:</span>
        <span class="stat-value">{{ store.stats.accuracy }}%</span>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { useLearningStore } from '@/stores/learning'

const store = useLearningStore()
</script>

<style lang="scss" scoped>
.app-header {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: $border-radius-md;
  padding: $spacing-lg;
  margin-bottom: $spacing-md;
  box-shadow: $box-shadow-md;
}

.header-title {
  font-size: 24px;
  font-weight: bold;
  color: $text-primary;
  margin-bottom: $spacing-md;
  text-align: center;
}

.stats-bar {
  display: flex;
  justify-content: space-around;
  align-items: center;
  flex-wrap: wrap;
  gap: $spacing-sm;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  min-width: 80px;
}

.stat-label {
  font-size: 12px;
  color: $text-secondary;
  margin-bottom: 5px;
}

.stat-value {
  font-size: 20px;
  font-weight: bold;
  color: $primary-gradient-start;
}

@media (max-width: 480px) {
  .header-title {
    font-size: 20px;
  }

  .stat-value {
    font-size: 18px;
  }
}
</style>
