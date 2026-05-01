<template>
  <div class="app-container">
    <router-view />
  </div>
</template>

<script setup lang="ts">
// The routing system now handles which view to display based on authentication state
</script>

<style lang="scss">
@use '@/styles/global.scss' as *;

.app-container {
  min-height: 100vh;
  background: $bg-paper;
  padding: $spacing-sm;
}
</style>