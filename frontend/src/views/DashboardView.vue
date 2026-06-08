<template>
  <div class="dashboard-container">
    <div class="dashboard-header">
      <div class="greeting">
        <h1 class="dashboard-title">欢迎，{{ user?.username }}</h1>
        <p class="dashboard-subtitle">持之以恒，积微成著</p>
      </div>
      <button class="btn-logout" @click="logout">退出</button>
    </div>

    <!-- 复习完成总结横幅 -->
    <div v-if="sessionSummary" class="session-summary">
      <div class="session-badge">🎉 复习完成！</div>
      <div class="session-stats">
        <span>本次正确率 {{ sessionSummary.accuracy }}%</span>
        <span class="session-divider">|</span>
        <span>完成 {{ sessionSummary.completed }} 题</span>
      </div>
      <div v-if="nextDueDays > 0" class="session-next">
        📅 下次复习：{{ nextDueDays === 1 ? '明天' : nextDueDays + '天后' }}预计有 {{ forecastDueSoon }} 题待复习
      </div>
      <button class="session-dismiss" @click="sessionSummary = null">✕</button>
    </div>

    <!-- 统计卡片 -->
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
      <div class="stat-card stat-card-reviews">
        <div class="stat-icon">🔄</div>
        <div class="stat-value">{{ stats.total_reviews }}</div>
        <div class="stat-label">总复习次数</div>
      </div>
    </div>

    <!-- 近期复习预告 -->
    <div v-if="forecast.length > 0" class="forecast-section">
      <h2 class="section-title">📅 近期复习预告</h2>
      <div class="forecast-list">
        <div
          v-for="(day, index) in forecast"
          :key="day.date"
          class="forecast-item"
          :class="{ 'forecast-today': index === 0, 'forecast-empty': day.count === 0 }"
        >
          <div class="forecast-label">
            {{ index === 0 ? '今天' : index === 1 ? '明天' : day.date.slice(5) }}
          </div>
          <div class="forecast-bar-wrap">
            <div
              class="forecast-bar"
              :style="{ width: forecastBarWidth(day.count) + '%' }"
            ></div>
          </div>
          <div class="forecast-count">{{ day.count }}题</div>
        </div>
      </div>
    </div>

    <!-- 分类进度 -->
    <div v-if="categories.length > 0" class="category-progress-section">
      <h2 class="section-title">📊 分类进度</h2>
      <div class="category-progress-list">
        <div v-for="cat in categories" :key="cat.name" class="category-progress-item">
          <div class="category-progress-header">
            <span class="category-progress-label">{{ cat.label || cat.name }}</span>
            <span class="category-progress-meta">{{ cat.total - cat.due }}/{{ cat.total }} 已掌握</span>
          </div>
          <div class="category-progress-bar-wrap">
            <div
              class="category-progress-bar"
              :style="{ width: categoryProgress(cat) + '%' }"
            ></div>
          </div>
        </div>
      </div>
    </div>

    <!-- 操作按钮 -->
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
import { onMounted, ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useLearningStore } from '@/stores/learning'
import { storeToRefs } from 'pinia'
import { showToast } from 'vant'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()
const learningStore = useLearningStore()

const { user } = storeToRefs(authStore)
const { stats, forecast, categories } = storeToRefs(learningStore)

interface SessionSummary {
  accuracy: string
  completed: number
}

const sessionSummary = ref<SessionSummary | null>(null)

// 从路由参数读取复习完成总结
if (route.query.accuracy && route.query.completed) {
  sessionSummary.value = {
    accuracy: route.query.accuracy as string,
    completed: Number(route.query.completed)
  }
  // 清除 query 参数，防止刷新重复显示
  router.replace({ path: '/dashboard', query: {} })
}

// 计算下次复习还有几天
const nextDueDays = computed(() => {
  if (forecast.value.length <= 1) return 0
  for (let i = 1; i < forecast.value.length; i++) {
    if ((forecast.value[i]?.count ?? 0) > 0) return i
  }
  return 0
})

const forecastDueSoon = computed(() => {
  if (nextDueDays.value === 0 || nextDueDays.value >= forecast.value.length) return 0
  return forecast.value[nextDueDays.value]?.count ?? 0
})

const maxForecastCount = computed(() => {
  if (forecast.value.length === 0) return 1
  return Math.max(...forecast.value.map(d => d.count), 1)
})

function forecastBarWidth(count: number): number {
  return Math.round((count / maxForecastCount.value) * 100)
}

function categoryProgress(cat: { total: number; due: number }): number {
  if (cat.total === 0) return 0
  return Math.round(((cat.total - cat.due) / cat.total) * 100)
}

const logout = () => {
  authStore.logout()
  router.push('/login')
}

const startLearning = () => {
  router.push('/categories')
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
  try {
    await Promise.all([
      learningStore.fetchStats(),
      learningStore.fetchForecast(),
      learningStore.fetchCategories()
    ])
  } catch (error) {
    console.error('Failed to load dashboard data:', error)
  }
})
</script>

<style scoped lang="scss">
@use '@/styles/global.scss' as *;

.dashboard-container {
  padding: $spacing-lg;
  padding-top: calc($spacing-lg + var(--safe-area-top));
  padding-bottom: calc($spacing-lg + var(--safe-area-bottom));
  min-height: 100vh;
  background: $bg-paper;
  max-width: 600px;
  margin: 0 auto;
}

.dashboard-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: $spacing-lg;

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

// 复习完成总结横幅
.session-summary {
  position: relative;
  background: linear-gradient(135deg, rgba($color-primary, 0.08), rgba($amber, 0.12));
  border: 1px solid rgba($amber, 0.3);
  border-radius: $border-radius-lg;
  padding: $spacing-lg;
  margin-bottom: $spacing-lg;
  text-align: center;

  .session-badge {
    font-size: $font-size-lg;
    font-weight: 700;
    color: $ink-deep;
    margin-bottom: $spacing-xs;
  }

  .session-stats {
    font-size: $font-size-sm;
    color: $text-secondary;
    margin-bottom: $spacing-xs;

    .session-divider {
      margin: 0 $spacing-sm;
      color: $text-muted;
    }
  }

  .session-next {
    font-size: $font-size-sm;
    color: $amber;
    font-weight: 500;
  }

  .session-dismiss {
    position: absolute;
    top: 8px;
    right: 12px;
    background: none;
    border: none;
    font-size: 18px;
    color: $text-muted;
    cursor: pointer;
    padding: 4px;
  }
}

.dashboard-stats {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: $spacing-sm;
  margin-bottom: $spacing-lg;
}

.stat-card {
  border-radius: $border-radius-lg;
  padding: $spacing-md $spacing-sm;
  text-align: center;
  box-shadow: $shadow-sm;

  .stat-icon {
    font-size: 20px;
    margin-bottom: $spacing-xs;
  }

  .stat-value {
    font-size: 24px;
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
  }

  &.stat-card-due {
    background: rgba($amber, 0.12);
    border: 1px solid rgba($amber, 0.3);
    .stat-value { color: $amber; }
  }

  &.stat-card-accuracy {
    background: rgba($success, 0.1);
    border: 1px solid rgba($success, 0.25);
    .stat-value { color: $success; }
  }

  &.stat-card-reviews {
    background: rgba($ink-blue, 0.08);
    border: 1px solid rgba($ink-blue, 0.2);
    .stat-value { color: $ink-blue; }
  }
}

// 近期复习预告
.forecast-section {
  background: $card-bg;
  border: 1px solid $card-border;
  border-radius: $border-radius-lg;
  padding: $spacing-lg;
  margin-bottom: $spacing-lg;
  box-shadow: $shadow-card;

  .section-title {
    font-size: $font-size-md;
    font-weight: 600;
    color: $ink-deep;
    margin: 0 0 $spacing-md;
  }

  .forecast-list {
    display: flex;
    flex-direction: column;
    gap: $spacing-sm;
  }

  .forecast-item {
    display: flex;
    align-items: center;
    gap: $spacing-sm;
    padding: 6px 0;

    .forecast-label {
      width: 48px;
      font-size: $font-size-xs;
      color: $text-secondary;
      text-align: right;
      flex-shrink: 0;
    }

    .forecast-bar-wrap {
      flex: 1;
      height: 8px;
      background: $bg-cream;
      border-radius: 4px;
      overflow: hidden;
    }

    .forecast-bar {
      height: 100%;
      background: $color-primary;
      border-radius: 4px;
      transition: width 0.3s ease;
      min-width: 4px;
    }

    .forecast-count {
      width: 36px;
      font-size: $font-size-xs;
      color: $text-secondary;
      text-align: left;
      flex-shrink: 0;
    }

    &.forecast-today {
      .forecast-label { color: $amber; font-weight: 600; }
      .forecast-bar { background: $amber; }
      .forecast-count { color: $amber; font-weight: 600; }
    }

    &.forecast-empty {
      .forecast-bar { opacity: 0.3; }
      .forecast-count { color: $text-muted; }
    }
  }
}

// 分类进度
.category-progress-section {
  background: $card-bg;
  border: 1px solid $card-border;
  border-radius: $border-radius-lg;
  padding: $spacing-lg;
  margin-bottom: $spacing-lg;
  box-shadow: $shadow-card;

  .section-title {
    font-size: $font-size-md;
    font-weight: 600;
    color: $ink-deep;
    margin: 0 0 $spacing-md;
  }

  .category-progress-list {
    display: flex;
    flex-direction: column;
    gap: $spacing-md;
  }

  .category-progress-item {
    .category-progress-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 4px;

      .category-progress-label {
        font-size: $font-size-sm;
        font-weight: 500;
        color: $ink-mid;
      }

      .category-progress-meta {
        font-size: $font-size-xs;
        color: $text-muted;
      }
    }

    .category-progress-bar-wrap {
      height: 6px;
      background: $bg-cream;
      border-radius: 3px;
      overflow: hidden;
    }

    .category-progress-bar {
      height: 100%;
      background: linear-gradient(90deg, $color-primary, $success);
      border-radius: 3px;
      transition: width 0.4s ease;
      min-width: 4px;
    }
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

// 超窄屏 stat cards 单列
@media (max-width: 340px) {
  .dashboard-stats {
    grid-template-columns: 1fr;
  }
  .dashboard-container {
    padding: $spacing-md;
    padding-top: calc($spacing-md + var(--safe-area-top));
    padding-bottom: calc($spacing-md + var(--safe-area-bottom));
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
