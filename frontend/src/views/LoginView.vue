<template>
  <div class="login-container">
    <div class="login-card">
      <h1 class="app-title">KnowLoop</h1>
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
        <div style="margin: 16px;">
          <van-button round block type="primary" native-type="submit">登录</van-button>
        </div>
      </van-form>

      <div class="register-link">
        <span>还没有账号？</span>
        <router-link to="/register">立即注册</router-link>
      </div>
    </div>

    <!-- Loading -->
    <van-loading v-if="isLoading" class="loading-overlay" size="24px" color="#1989fa" />
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
  background: linear-gradient(135deg, $primary-gradient-start 0%, $primary-gradient-end 100%);
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: $spacing-xl;
  background: white;
  border-radius: $border-radius-lg;
  box-shadow: $shadow-lg;

  .app-title {
    text-align: center;
    margin-bottom: $spacing-lg;
    color: $text-primary;
    font-size: 20px;
    font-weight: 600;
    line-height: 1.4;
  }

  .login-title {
    text-align: center;
    margin-bottom: $spacing-xl;
    color: $text-primary;
    font-size: $font-size-xl;
  }

  .register-link {
    text-align: center;
    margin-top: $spacing-md;
    color: $text-secondary;

    a {
      color: $color-primary;
      text-decoration: none;
      margin-left: $spacing-xs;
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