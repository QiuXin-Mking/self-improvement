<template>
  <div class="category-select-container">
    <div class="category-header">
      <h1 class="category-title">选择复习分类</h1>
      <p class="category-subtitle">选择一个或多个分类开始复习</p>
    </div>

    <div v-if="categories.length === 0" class="empty-state">
      <p>暂无分类数据</p>
      <van-button type="primary" @click="goBack">返回</van-button>
    </div>

    <div v-else class="category-list">
      <div
        v-for="cat in categories"
        :key="cat.name"
        class="category-card"
        :class="{ selected: selectedCategories.has(cat.name) }"
        @click="toggleCategory(cat.name)"
      >
        <div class="category-info">
          <div class="category-label">{{ cat.label }}</div>
          <div class="category-meta">
            <span class="meta-due">待复习 {{ cat.due }}</span>
            <span class="meta-divider">/</span>
            <span class="meta-total">共 {{ cat.total }}</span>
          </div>
        </div>
        <div class="category-check">
          <div v-if="selectedCategories.has(cat.name)" class="check-icon">✓</div>
        </div>
      </div>
    </div>

    <div class="category-actions">
      <div class="selection-summary">
        <template v-if="selectedCategories.size === 0">请选择至少一个分类</template>
        <template v-else-if="selectedCategories.size === categories.length">已选择全部 {{ categories.length }} 个分类</template>
        <template v-else>已选择 {{ selectedCategories.size }} / {{ categories.length }} 个分类</template>
      </div>

      <van-button
        type="primary"
        size="large"
        block
        :disabled="selectedCategories.size === 0"
        @click="startReview"
      >
        开始复习
      </van-button>
    </div>

    <div class="back-link">
      <van-button type="default" size="small" @click="goBack">返回首页</van-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useLearningStore } from '@/stores/learning'
import { storeToRefs } from 'pinia'

const router = useRouter()
const store = useLearningStore()
const { categories } = storeToRefs(store)
const selectedCategories = ref<Set<string>>(new Set())

function toggleCategory(name: string) {
  if (selectedCategories.value.has(name)) {
    selectedCategories.value.delete(name)
  } else {
    selectedCategories.value.add(name)
  }
}

function startReview() {
  const selected = Array.from(selectedCategories.value).join(',')
  store.activeCategories = selected
  router.push(`/learn?categories=${encodeURIComponent(selected)}`)
}

function goBack() {
  router.push('/dashboard')
}

onMounted(async () => {
  const cats = await store.fetchCategories()
  // 默认全选所有有待复习题目的分类
  for (const cat of cats) {
    if (cat.due > 0) {
      selectedCategories.value.add(cat.name)
    }
  }
})
</script>

<style scoped lang="scss">
@use '@/styles/global.scss' as *;

.category-select-container {
  padding: $spacing-lg;
  min-height: 100vh;
  background: $bg-paper;
  max-width: 600px;
  margin: 0 auto;
}

.category-header {
  margin-bottom: $spacing-xl;

  .category-title {
    margin: 0;
    font-size: $font-size-xl;
    color: $ink-deep;
  }
  .category-subtitle {
    margin: $spacing-xs 0 0;
    font-size: $font-size-sm;
    color: $text-muted;
  }
}

.empty-state {
  text-align: center;
  padding: $spacing-xl;
  color: $text-muted;
}

.category-list {
  display: flex;
  flex-direction: column;
  gap: $spacing-sm;
  margin-bottom: $spacing-xl;
}

.category-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: $spacing-lg;
  background: $card-bg;
  border: 1px solid $card-border;
  border-radius: $border-radius-lg;
  cursor: pointer;
  transition: all $transition-fast;

  &.selected {
    border-color: $color-primary;
    background: rgba($color-primary, 0.04);
  }

  &:active {
    transform: scale(0.98);
  }

  .category-info {
    .category-label {
      font-size: $font-size-md;
      font-weight: 600;
      color: $ink-deep;
      margin-bottom: 4px;
    }
    .category-meta {
      font-size: $font-size-xs;
      .meta-due {
        color: $amber;
        font-weight: 500;
      }
      .meta-divider {
        color: $text-muted;
        margin: 0 4px;
      }
      .meta-total {
        color: $text-muted;
      }
    }
  }

  .category-check {
    .check-icon {
      width: 24px;
      height: 24px;
      background: $color-primary;
      color: white;
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 14px;
      font-weight: 700;
    }
  }
}

.category-actions {
  background: $card-bg;
  border: 1px solid $card-border;
  border-radius: $border-radius-lg;
  padding: $spacing-lg;
  box-shadow: $shadow-card;
  margin-bottom: $spacing-lg;

  .selection-summary {
    text-align: center;
    font-size: $font-size-sm;
    color: $text-secondary;
    margin-bottom: $spacing-md;
  }
}

.back-link {
  text-align: center;
}
</style>
