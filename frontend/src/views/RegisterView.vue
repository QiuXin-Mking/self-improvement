<template>
  <div class="register-container">
    <div class="register-card">
      <div class="brand-mark">◆</div>
      <h1 class="app-title">KnowLoop</h1>
      <p class="app-subtitle">开启你的知识旅程</p>
      <h2 class="register-title">注册</h2>

      <van-form @submit="onSubmit">
        <van-cell-group inset>
          <van-field
            v-model="username"
            name="username"
            label="用户名"
            placeholder="请输入用户名"
            :rules="[{ required: true, message: '请填写用户名' }, { validator: usernameValidator, message: '用户名长度3-32位' }]"
          />
          <van-field
            v-model="password"
            type="password"
            name="password"
            label="密码"
            placeholder="请输入密码"
            :rules="[{ required: true, message: '请填写密码' }, { validator: passwordValidator, message: '密码长度至少6位' }]"
          />
          <van-field
            v-model="confirmPassword"
            type="password"
            name="confirmPassword"
            label="确认密码"
            placeholder="请再次输入密码"
            :rules="[{ required: true, message: '请确认密码' }, { validator: confirmPasswordValidator, message: '两次输入密码不一致' }]"
          />
        </van-cell-group>
        <div class="form-actions">
          <van-button round block type="primary" native-type="submit">注册</van-button>
        </div>
      </van-form>

      <div class="login-link">
        <span>已有账号？</span>
        <router-link to="/login">立即登录</router-link>
      </div>
    </div>

    <van-loading v-if="isLoading" class="loading-overlay" size="24px" color="#2c3e6b" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { showSuccessToast, showFailToast } from 'vant'

const router = useRouter()
const authStore = useAuthStore()

const username = ref('')
const password = ref('')
const confirmPassword = ref('')
const isLoading = ref(false)

// Validation functions
const usernameValidator = (val: string) => val.length >= 3 && val.length <= 32
const passwordValidator = (val: string) => val.length >= 6
const confirmPasswordValidator = (val: string) => val === password.value

const onSubmit = async () => {
  isLoading.value = true

  try {
    const result = await authStore.register(username.value, password.value)

    if (result.success) {
      showSuccessToast({
        message: result.message || '注册成功',
        position: 'top'
      })
      // Redirect to dashboard after successful registration
      router.push('/dashboard')
    } else {
      showFailToast({
        message: result.message || '注册失败',
        position: 'top'
      })
    }
  } catch (error) {
    console.error('Registration error:', error)
    showFailToast({
      message: '注册时发生错误',
      position: 'top'
    })
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped lang="scss">
@use '@/styles/global.scss' as *;

.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: $spacing-lg;
  background: $bg-paper;
}

.register-card {
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

  .register-title {
    text-align: center;
    margin-bottom: $spacing-xl;
    color: $text-primary;
    font-size: $font-size-lg;
    font-weight: 500;
  }

  .form-actions {
    margin: $spacing-md;
  }

  .login-link {
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
  }

  :deep(.van-cell-group--inset .van-cell) {
    margin-bottom: $spacing-md;
    border-radius: $border-radius-sm;
  }

  :deep(.van-cell-group--inset .van-cell:last-child) {
    margin-bottom: 0;
  }
}

.loading-overlay {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>