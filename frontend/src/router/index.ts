import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

// Define route components (lazy loading)
const LoginView = () => import('@/views/LoginView.vue')
const RegisterView = () => import('@/views/RegisterView.vue')
const DashboardView = () => import('@/views/DashboardView.vue')
const LearningView = () => import('@/views/LearningView.vue') // Reusing existing component
const ProtectedLayout = () => import('@/layouts/ProtectedLayout.vue')

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: LoginView
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterView
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: DashboardView
  },
  {
    path: '/learn',
    name: 'Learn',
    component: LearningView
  },
  {
    path: '/',
    redirect: '/dashboard'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Global navigation guard
router.beforeEach(async (to, from, next) => {
  console.log('Router guard: navigating to', to.path, 'name:', to.name)
  const authStore = useAuthStore()
  const token = authStore.token
  console.log('Router guard: token exists?', !!token)

  // Public routes
  const publicRoutes = ['Login', 'Register']

  if (publicRoutes.includes(to.name as string)) {
    // If trying to access public route and already authenticated, redirect to dashboard
    if (token) {
      console.log('Router guard: has token, redirecting to dashboard')
      next('/dashboard')
    } else {
      console.log('Router guard: no token, allowing public route')
      next()
    }
  } else {
    // Protected routes
    if (token) {
      console.log('Router guard: has token, allowing protected route')
      next()
    } else {
      console.log('Router guard: no token, redirecting to login')
      next('/login')
    }
  }
})

// Log router errors
router.onError((error) => {
  console.error('Router error:', error)
})

export default router