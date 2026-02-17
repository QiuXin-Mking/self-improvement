<template>
  <div class="init-prompt-view">
    <div class="card">
      <h2>欢迎使用知识库学习系统！</h2>
      <p>首次使用需要初始化知识库</p>
      <button class="btn btn-primary" :disabled="loading" @click="handleInit">
        {{ loading ? '正在初始化...' : '初始化知识库' }}
      </button>
      <div v-if="statusMessage" class="status-message" :class="statusType">
        {{ statusMessage }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { showToast } from 'vant'
import { useLearningStore } from '@/stores/learning'

const emit = defineEmits<{
  (e: 'init-success'): void
}>()

const store = useLearningStore()
const loading = ref(false)
const statusMessage = ref('')
const statusType = ref<'success' | 'error'>('success')

async function handleInit() {
  loading.value = true
  statusMessage.value = '正在初始化，请稍候...'
  statusType.value = 'success'

  try {
    const result = await store.initDatabase()

    if (result) {
      statusMessage.value = `✅ 成功导入 ${result.imported} 个新问题到知识库！`
      statusType.value = 'success'

      showToast({
        type: 'success',
        message: '初始化成功！',
        duration: 2000
      })

      // 2秒后通知父组件
      setTimeout(() => {
        emit('init-success')
      }, 2000)
    }
  } catch (error) {
    statusMessage.value = '❌ 初始化失败，请检查配置'
    statusType.value = 'error'
  } finally {
    loading.value = false
  }
}
</script>

<style lang="scss" scoped>
.init-prompt-view {
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
  font-size: 22px;
  color: $text-primary;
  margin-bottom: $spacing-md;
}

p {
  font-size: 16px;
  color: $text-secondary;
  margin-bottom: $spacing-lg;
}

.btn-primary {
  min-width: 200px;
}

.status-message {
  margin-top: $spacing-lg;
  padding: $spacing-md;
  border-radius: $border-radius-sm;
  font-size: 14px;

  &.success {
    background: lighten($success-color, 40%);
    color: darken($success-color, 10%);
  }

  &.error {
    background: lighten($error-color, 40%);
    color: darken($error-color, 10%);
  }
}
</style>
