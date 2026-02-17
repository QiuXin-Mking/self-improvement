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
    const message = (error.response?.data as any)?.error || '网络错误，请重试'

    // If receiving 401, redirect to login (could be implemented here)
    if (error.response?.status === 401) {
      // Optionally redirect to login page
      // router.push('/login')
    }

    showToast({
      type: 'fail',
      message,
      duration: 2000
    })
    return Promise.reject(error)
  }
)

export default api
