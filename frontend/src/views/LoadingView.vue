<template>
  <div class="loading-view">
    <van-loading size="48" color="#667eea" vertical>加载中...</van-loading>
  </div>
</template>

<script setup lang="ts">
import { Loading as VanLoading } from 'vant'
</script>

<style lang="scss" scoped>
.loading-view {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 300px;
}
</style>
