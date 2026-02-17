<template>
  <div class="dashboard-container">
    <div class="dashboard-header">
      <h1 class="dashboard-title">欢迎，{{ user?.username }}！</h1>
      <van-button type="default" size="small" @click="logout">退出登录</van-button>
    </div>

    <div class="dashboard-stats">
      <div class="stat-card">
        <div class="stat-value">{{ stats.total_questions }}</div>
        <div class="stat-label">总问题数</div>
      </div>
      <div class="stat-card">
        <div class="stat-value">{{ stats.due_questions }}</div>
        <div class="stat-label">今日待复习</div>
      </div>
      <div class="stat-card">
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
        style="margin-top: 16px;"
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
  background: linear-gradient(135deg, $primary-gradient-start 0%, $primary-gradient-end 100%);
}

.dashboard-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: $spacing-xl;
  color: white;

  .dashboard-title {
    margin: 0;
    font-size: $font-size-lg;
  }
}

.dashboard-stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: $spacing-md;
  margin-bottom: $spacing-xl;
}

.stat-card {
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
  border-radius: $border-radius-lg;
  padding: $spacing-md;
  text-align: center;
  color: white;

  .stat-value {
    font-size: $font-size-xl;
    font-weight: bold;
    margin-bottom: $spacing-xs;
  }

  .stat-label {
    font-size: $font-size-sm;
    opacity: 0.8;
  }
}

.dashboard-actions {
  background: white;
  border-radius: $border-radius-lg;
  padding: $spacing-lg;
  box-shadow: $shadow-md;
}
</style>