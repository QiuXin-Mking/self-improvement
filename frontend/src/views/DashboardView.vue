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
        type="primary"
        size="large"
        block
        class="btn-import"
        @click="showImportSheet = true"
      >
        导入知识
      </van-button>

      <van-action-sheet
        v-model:show="showImportSheet"
        title="选择导入方式"
        :actions="importActions"
        @select="onImportSelect"
      />

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

    <input
      ref="fileInput"
      type="file"
      accept=".zip,.md,.markdown"
      style="display: none"
      @change="onFileSelected"
    />

    <van-dialog
      v-model:show="showManualDialog"
      title="添加问题"
      show-cancel-button
      class="import-dialog"
      @confirm="onManualConfirm"
    >
      <div class="manual-entry-form">
        <van-field
          v-model="manualQuestion"
          label="问题"
          placeholder="输入问题内容"
        />
        <van-field
          v-model="manualAnswer"
          label="答案"
          placeholder="输入答案内容"
        />
      </div>
    </van-dialog>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
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
      showToast({ message: response.message || '知识库初始化成功', type: 'success' })
    } else {
      showToast({ message: '没有找到新的问题', type: 'fail' })
    }
  } catch (error: any) {
    const errMsg = error?.response?.data?.error || '初始化失败'
    showToast({ message: errMsg, type: 'fail' })
  }
}

const showImportSheet = ref(false)
const fileInput = ref<HTMLInputElement>()
const pendingImportType = ref<'zip' | 'md'>('zip')

const importActions = [
  { name: '上传 .md 文件', value: 'md', subname: '导入单个 Markdown 题库文件' },
  { name: '上传 zip 压缩包', value: 'zip', subname: '批量导入多个 .md 文件' },
  { name: '手动输入问题', value: 'manual', subname: '逐条添加问题' },
]

const onImportSelect = (action: { value: string }) => {
  showImportSheet.value = false

  if (action.value === 'zip' || action.value === 'md') {
    pendingImportType.value = action.value as 'zip' | 'md'
    fileInput.value?.click()
  } else if (action.value === 'manual') {
    showManualEntryDialog()
  }
}

const onFileSelected = async () => {
  const file = fileInput.value?.files?.[0]
  if (!file) return

  try {
    let result
    if (pendingImportType.value === 'zip') {
      result = await learningStore.uploadZip(file)
    } else {
      result = await learningStore.uploadMd(file)
    }

    if (result) {
      showToast({ message: result.message || '导入成功', type: 'success' })
    }
  } catch (error: any) {
    const errMsg = error?.response?.data?.error || '导入失败'
    showToast({ message: errMsg, type: 'fail' })
  }

  if (fileInput.value) {
    fileInput.value.value = ''
  }
}

const showManualDialog = ref(false)
const manualQuestion = ref('')
const manualAnswer = ref('')

const showManualEntryDialog = () => {
  manualQuestion.value = ''
  manualAnswer.value = ''
  showManualDialog.value = true
}

const onManualConfirm = async () => {
  const q = manualQuestion.value.trim()
  const a = manualAnswer.value.trim()

  if (!q || !a) {
    showToast({ message: '问题和答案不能为空', type: 'fail' })
    return
  }

  try {
    await learningStore.addQuestion(q, a)
    showToast({ message: '问题添加成功', type: 'success' })
    showManualDialog.value = false
  } catch (error: any) {
    const errMsg = error?.response?.data?.error || '添加失败'
    showToast({ message: errMsg, type: 'fail' })
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
  display: flex;
  flex-direction: column;
  gap: $spacing-md;

  .btn-import {
    margin-bottom: 0;
  }

  .btn-init {
    margin-top: 0;
  }
}

</style>

<style lang="scss">
@use '@/styles/global.scss' as *;

.manual-entry-form {
  padding: $spacing-lg $spacing-md;

  .van-cell {
    background: $bg-cream;
    border: 1px solid $card-border;
    border-radius: $border-radius-sm;
    margin-bottom: $spacing-md;
    padding: $spacing-sm $spacing-md;
  }

  .van-field__label {
    color: $text-secondary;
    font-weight: 500;
    margin-right: $spacing-sm;
  }
}

.import-dialog {
  .van-dialog__header {
    padding: $spacing-lg $spacing-lg $spacing-sm;
    font-size: $font-size-lg;
    font-weight: 600;
    color: $ink-deep;
  }

  .van-dialog__content {
    padding: 0;
  }

  .van-dialog__footer {
    padding: $spacing-sm $spacing-lg $spacing-lg;
  }
}
</style>