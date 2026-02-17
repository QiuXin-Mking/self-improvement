import { defineStore } from 'pinia'
import { ref } from 'vue'
import { authApi } from '@/api/auth'

export const useAuthStore = defineStore('auth', () => {
  // State
  const token = ref<string | null>(localStorage.getItem('token'))
  const user = ref<{ id: number; username: string } | null>(null)
  const isAuthenticated = ref(false)

  // Actions
  async function login(username: string, password: string) {
    try {
      const response = await authApi.login({ username, password })

      if (response.success && response.data) {
        const { token: authToken, user_id, username: userName } = response.data

        // Store token in localStorage
        localStorage.setItem('token', authToken)
        token.value = authToken
        user.value = { id: user_id, username: userName }
        isAuthenticated.value = true

        return { success: true, message: response.message }
      } else {
        return { success: false, message: response.error || 'Login failed' }
      }
    } catch (error) {
      console.error('Login error:', error)
      return { success: false, message: 'Network error occurred' }
    }
  }

  async function register(username: string, password: string) {
    try {
      const response = await authApi.register({ username, password })

      if (response.success && response.data) {
        const { token: authToken, user_id, username: userName } = response.data

        // Store token in localStorage
        localStorage.setItem('token', authToken)
        token.value = authToken
        user.value = { id: user_id, username: userName }
        isAuthenticated.value = true

        return { success: true, message: response.message }
      } else {
        return { success: false, message: response.error || 'Registration failed' }
      }
    } catch (error) {
      console.error('Registration error:', error)
      return { success: false, message: 'Network error occurred' }
    }
  }

  function logout() {
    localStorage.removeItem('token')
    token.value = null
    user.value = null
    isAuthenticated.value = false
  }

  async function loadUserProfile() {
    if (!token.value) {
      console.log('loadUserProfile: no token')
      return false
    }

    try {
      const response = await authApi.profile()
      console.log('loadUserProfile response:', response)

      if (response.success && response.data) {
        const { user_id, username } = response.data
        user.value = { id: user_id, username }
        isAuthenticated.value = true
        console.log('loadUserProfile: success, user:', user.value)
        return true
      } else {
        // Token might be invalid/expired, logout
        console.log('loadUserProfile: failed, logging out')
        logout()
        return false
      }
    } catch (error) {
      console.error('Profile loading error:', error)
      logout()
      return false
    }
  }

  return {
    // State
    token,
    user,
    isAuthenticated,

    // Actions
    login,
    register,
    logout,
    loadUserProfile
  }
})