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
import { showSuccessToast, showFailToast } from 'vant'

const router = useRouter()
const authStore = useAuthStore()

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
      showSuccessToast({
        message: result.message || '登录成功',
        position: 'top'
      })
      // Redirect to dashboard after successful login
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
}

.loading-overlay {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>