<template>
  <div class="client-config">
    <div class="page-head">
      <h2 class="page-title">客户端配置</h2>
      <p class="page-subtitle">管理客户端用户、凭证与下发配置</p>
    </div>
    <UserConfig />
  </div>
</template>

<script lang="ts" setup>
import UserConfig from '../components/UserConfig.vue'
</script>

<style scoped>
.page-head {
  margin-bottom: 20px;
}
</style>
