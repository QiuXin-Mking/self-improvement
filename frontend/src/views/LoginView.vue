<template>
  <div class="login-container">
    <div class="login-card">
      <div class="brand-mark">◆</div>
      <h1 class="app-title">KnowLoop</h1>
      <p class="app-subtitle">知识的温度，记忆的节奏</p>
      <h2 class="login-title">登录</h2>

      <van-form @submit="onSubmit">
        <van-cell-group inset>
          <van-field
            v-model="username"
            name="username"
            label="用户名"
            placeholder="请输入用户名"
            :rules="[{ required: true, message: '请填写用户名' }, { validator: usernameValidator, message: '用户名长度至少3位' }]"
          />
          <van-field
            v-model="password"
            type="password"
            name="password"
            label="密码"
            placeholder="请输入密码"
            :rules="[{ required: true, message: '请填写密码' }, { validator: passwordValidator, message: '密码长度至少6位' }]"
          />
        </van-cell-group>
        <div class="form-actions">
          <van-button round block type="primary" native-type="submit">登录</van-button>
        </div>

        <div class="demo-section">
          <div class="demo-divider"><span>or</span></div>
          <van-button round block plain type="warning" class="btn-demo" @click="demoLogin">
            🎯 一键体验
          </van-button>
          <p class="demo-hint">无需注册，点击即可体验完整功能</p>
        </div>
      </van-form>

      <div class="register-link">
        <span>还没有账号？</span>
        <router-link to="/register">立即注册</router-link>
      </div>
    </div>

    <van-loading v-if="isLoading" class="loading-overlay" size="24px" color="#2c3e6b" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useLearningStore } from '@/stores/learning'
import { showFailToast } from 'vant'

const router = useRouter()
const authStore = useAuthStore()
const learningStore = useLearningStore()

const username = ref('')
const password = ref('')
const isLoading = ref(false)

// Validation functions
const usernameValidator = (val: string) => val.length >= 3
const passwordValidator = (val: string) => val.length >= 6

const onSubmit = async () => {
  isLoading.value = true

  try {
    const result = await authStore.login(username.value, password.value)

    if (result.success) {
      // 直接跳转，无需 toast（跳转本身即是成功反馈，避免白框闪烁）
      router.push('/dashboard')
    } else {
      showFailToast({
        message: result.message || '登录失败',
        position: 'top'
      })
    }
  } catch (error) {
    console.error('Login error:', error)
    showFailToast({
      message: '登录时发生错误',
      position: 'top'
    })
  } finally {
    isLoading.value = false
  }
}

const demoLogin = async () => {
  isLoading.value = true

  try {
    const result = await authStore.login('demo', 'demo123')

    if (result.success) {
      // 重置体验数据，确保每次一键体验都是全新状态
      try {
        await learningStore.resetDemo()
      } catch (e) {
        // 重置失败不阻塞流程，继续进入体验
        console.error('重置体验数据失败:', e)
      }
      // 导航到分类选择页，让体验流程使用分类功能
      router.push('/categories')
    } else {
      showFailToast({
        message: result.message || '体验登录失败',
        position: 'top'
      })
    }
  } catch (error) {
    console.error('Demo login error:', error)
    showFailToast({
      message: '登录时发生错误',
      position: 'top'
    })
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped lang="scss">
@use '@/styles/global.scss' as *;

.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: $spacing-lg;
  background: $bg-paper;
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: $spacing-2xl $spacing-xl;
  background: $card-bg;
  border: 1px solid $card-border;
  border-radius: $border-radius-xl;
  box-shadow: $shadow-card;

  .brand-mark {
    text-align: center;
    font-size: 32px;
    color: $amber;
    margin-bottom: $spacing-sm;
  }

  .app-title {
    text-align: center;
    margin-bottom: $spacing-xs;
    color: $ink-deep;
    font-size: 28px;
    font-weight: 700;
    letter-spacing: 2px;
  }

  .app-subtitle {
    text-align: center;
    color: $text-muted;
    font-size: $font-size-sm;
    margin-bottom: $spacing-xl;
    letter-spacing: 1px;
  }

  .login-title {
    text-align: center;
    margin-bottom: $spacing-xl;
    color: $text-primary;
    font-size: $font-size-lg;
    font-weight: 500;
  }

  .form-actions {
    margin: $spacing-md;
  }

  .demo-section {
    margin: $spacing-md;

    .demo-divider {
      display: flex;
      align-items: center;
      margin-bottom: $spacing-md;

      &::before,
      &::after {
        content: '';
        flex: 1;
        height: 1px;
        background: $card-border;
      }

      span {
        padding: 0 $spacing-sm;
        color: $text-muted;
        font-size: $font-size-xs;
        text-transform: uppercase;
        letter-spacing: 1px;
      }
    }

    .btn-demo {
      font-size: $font-size-md;
      font-weight: 600;
      border-width: 2px;
    }

    .demo-hint {
      text-align: center;
      margin-top: $spacing-sm;
      color: $text-muted;
      font-size: $font-size-xs;
      margin-bottom: 0;
    }
  }

  .register-link {
    text-align: center;
    margin-top: $spacing-md;
    color: $text-secondary;
    font-size: $font-size-sm;

    a {
      color: $ink-blue;
      text-decoration: none;
      margin-left: $spacing-xs;
      font-weight: 500;
    }
  }

  :deep(.van-cell-group--inset) {
    margin: 0;
    background: transparent;
  }

  :deep(.van-cell-group--inset .van-cell) {
    margin-bottom: $spacing-md;
    border-radius: $border-radius-sm;
    background: $bg-cream;
    border: 1px solid $card-border;
    padding: $spacing-sm $spacing-md;
    transition: border-color $transition-fast;

    &:focus-within {
      border-color: $ink-blue;
    }
  }

  :deep(.van-cell-group--inset .van-cell:last-child) {
    margin-bottom: 0;
  }

  :deep(.van-field__label) {
    color: $ink-mid;
    font-weight: 500;
  }

  :deep(.van-field__control::placeholder) {
    color: $text-muted;
  }

  :deep(.van-field__body) {
    color: $text-primary;
  }

  :deep(.van-cell:focus-within::after) {
    border-color: $amber;
  }
}

.loading-overlay {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>