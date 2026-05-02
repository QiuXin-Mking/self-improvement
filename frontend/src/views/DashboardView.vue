<template>
  <div class="dashboard-container">
    <div class="dashboard-header">
      <div class="greeting">
        <h1 class="dashboard-title">欢迎，{{ user?.username }}</h1>
        <p class="dashboard-subtitle">持之以恒，积微成著</p>
      </div>
      <button class="btn-logout" @click="logout">退出</button>
    </div>

    <div class="dashboard-stats">
      <div class="stat-card stat-card-total">
        <div class="stat-icon">📚</div>
        <div class="stat-value">{{ stats.total_questions }}</div>
        <div class="stat-label">总问题数</div>
      </div>
      <div class="stat-card stat-card-due">
        <div class="stat-icon">📝</div>
        <div class="stat-value">{{ stats.due_questions }}</div>
        <div class="stat-label">今日待复习</div>
      </div>
      <div class="stat-card stat-card-accuracy">
        <div class="stat-icon">🎯</div>
        <div class="stat-value">{{ stats.accuracy }}%</div>
        <div class="stat-label">正确率</div>
      </div>
    </div>

    <div class="dashboard-actions">
      <van-button
        type="primary"
        size="large"
        block
        :disabled="stats.due_questions === 0"
        @click="startLearning"
      >
        {{ stats.due_questions > 0 ? `开始复习 (${stats.due_questions})` : '暂无待复习问题' }}
      </van-button>

      <van-button
        type="default"
        size="large"
        block
        class="btn-init"
        @click="initDatabase"
      >
        重新初始化知识库
      </van-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useLearningStore } from '@/stores/learning'
import { storeToRefs } from 'pinia'
import { showToast } from 'vant'

const router = useRouter()
const authStore = useAuthStore()
const learningStore = useLearningStore()

const { user } = storeToRefs(authStore)
const { stats } = storeToRefs(learningStore)

const logout = () => {
  authStore.logout()
  router.push('/login')
}

const startLearning = () => {
  router.push('/learn')
}

const initDatabase = async () => {
  try {
    const response = await learningStore.initDatabase()
    if (response) {
      await learningStore.fetchStats()
    }
  } catch (error) {
    console.error('Init error:', error)
    showToast({ message: '初始化失败', type: 'fail' })
  }
}

onMounted(async () => {
  console.log('DashboardView onMounted')
  try {
    await learningStore.fetchStats()
    console.log('Stats loaded:', stats.value)
  } catch (error) {
    console.error('Failed to load stats:', error)
  }
})
</script>

<style scoped lang="scss">
@use '@/styles/global.scss' as *;

.dashboard-container {
  padding: $spacing-lg;
  min-height: 100vh;
  background: $bg-paper;
}

.dashboard-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: $spacing-xl;

  .greeting {
    .dashboard-title {
      margin: 0;
      font-size: $font-size-xl;
      color: $ink-deep;
    }
    .dashboard-subtitle {
      margin: $spacing-xs 0 0;
      font-size: $font-size-sm;
      color: $text-muted;
    }
  }

  .btn-logout {
    background: none;
    border: none;
    color: $text-muted;
    font-size: $font-size-sm;
    cursor: pointer;
    padding: 4px 0;
    font-family: inherit;
    transition: color $transition-fast;

    &:hover {
      color: $danger;
    }
  }
}

.dashboard-stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: $spacing-sm;
  margin-bottom: $spacing-xl;
}

.stat-card {
  border-radius: $border-radius-lg;
  padding: $spacing-lg $spacing-sm;
  text-align: center;
  box-shadow: $shadow-sm;

  .stat-icon {
    font-size: 24px;
    margin-bottom: $spacing-xs;
  }

  .stat-value {
    font-size: 26px;
    font-weight: 700;
    margin-bottom: 2px;
  }

  .stat-label {
    font-size: $font-size-xs;
    opacity: 0.7;
  }

  &.stat-card-total {
    background: $bg-cream;
    border: 1px solid $card-border;
    .stat-value { color: $ink-deep; }
    .stat-label { color: $text-secondary; }
  }

  &.stat-card-due {
    background: rgba($amber, 0.12);
    border: 1px solid rgba($amber, 0.3);
    .stat-value { color: $amber; }
    .stat-label { color: $amber; opacity: 0.8; }
  }

  &.stat-card-accuracy {
    background: rgba($success, 0.1);
    border: 1px solid rgba($success, 0.25);
    .stat-value { color: $success; }
    .stat-label { color: $success; opacity: 0.8; }
  }
}

.dashboard-actions {
  background: $card-bg;
  border: 1px solid $card-border;
  border-radius: $border-radius-lg;
  padding: $spacing-lg;
  box-shadow: $shadow-card;

  .btn-init {
    margin-top: $spacing-md;
  }
}
</style>