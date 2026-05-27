import axios from 'axios'
import type { AxiosInstance, AxiosError } from 'axios'
import { showToast } from 'vant'

const api: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Request interceptor - add auth token if available
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor
api.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error: AxiosError) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      showToast({
        type: 'fail',
        message: '登录已过期，请重新登录',
        duration: 1500
      })
      setTimeout(() => {
        window.location.href = '/login'
      }, 800)
      return Promise.reject(error)
    }

    // For other errors, let the caller handle the toast
    return Promise.reject(error)
  }
)

export default api
