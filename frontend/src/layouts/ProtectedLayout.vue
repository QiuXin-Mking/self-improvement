<template>
  <div class="protected-layout">
    <slot v-if="!isLoading" />
    <div v-else class="loading-state">加载中...</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()
const isLoading = ref(true)

onMounted(async () => {
  console.log('ProtectedLayout onMounted, token:', authStore.token)
  // Check if user is authenticated
  const isAuthenticated = await authStore.loadUserProfile()

  console.log('ProtectedLayout isAuthenticated:', isAuthenticated, 'user:', authStore.user)

  if (!isAuthenticated) {
    // Redirect to login if not authenticated
    console.log('Redirecting to login')
    router.push('/login')
  } else {
    isLoading.value = false
  }
})
</script>

<style scoped>
.protected-layout {
  min-height: 100vh;
}

.loading-state {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  color: white;
  font-size: 18px;
}
</style>