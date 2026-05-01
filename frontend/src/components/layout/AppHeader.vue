<template>
  <header class="app-header">
    <h1 class="header-title">KnowLoop</h1>
    <div class="stats-bar">
      <div class="stat-item">
        <span class="stat-label">待复习</span>
        <span class="stat-value">{{ store.stats.due_questions }}</span>
      </div>
      <div class="stat-item">
        <span class="stat-label">总问题</span>
        <span class="stat-value">{{ store.stats.total_questions }}</span>
      </div>
      <div class="stat-item">
        <span class="stat-label">正确率</span>
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
  background: $card-bg;
  border: 1px solid $card-border;
  border-radius: $border-radius-md;
  padding: $spacing-lg;
  margin-bottom: $spacing-md;
  box-shadow: $shadow-sm;
}

.header-title {
  font-size: 22px;
  font-weight: 700;
  color: $ink-deep;
  margin-bottom: $spacing-md;
  text-align: center;
  letter-spacing: 1px;
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
  min-width: 70px;
}

.stat-label {
  font-size: 11px;
  color: $text-muted;
  margin-bottom: 4px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.stat-value {
  font-size: 20px;
  font-weight: 700;
  color: $ink-blue;
}

@media (max-width: 480px) {
  .header-title { font-size: 20px; }
  .stat-value { font-size: 18px; }
}
</style>
