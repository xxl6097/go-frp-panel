<template>
  <div>
    <el-page-header
      :icon="null"
      style="width: 100%; margin-left: 30px; margin-bottom: 20px"
    >
      <template #title>
        <span>{{ proxyType }}</span>
      </template>
      <template #content></template>
      <template #extra>
        <div class="flex items-center" style="margin-right: 30px">
          <el-popconfirm
            title="您确定清除所有离线代理数据吗？"
            @confirm="clearOfflineProxies"
          >
            <template #reference>
              <el-button>清除离线代理</el-button>
            </template>
          </el-popconfirm>
          <el-button @click="$emit('refresh')">刷新</el-button>
        </div>
      </template>
    </el-page-header>

    <el-table
      :data="proxies"
      :default-sort="{ prop: 'name', order: 'ascending' }"
      style="width: 100%"
    >
      <el-table-column type="expand">
        <template #default="props">
          <!--          <ProxyViewExpand :row="props.row" :proxyType="proxyType" />-->
          <ProxyListView :proxies="props.row.list" proxyType="proxyType" />
        </template>
      </el-table-column>
      <el-table-column label="名称" prop="name" sortable></el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { ProxyConfig } from '../utils/proxy.js'
import { ElMessage } from 'element-plus'
import ProxyListView from './ProxyListView.vue'

defineProps<{
  proxies: ProxyConfig[]
  proxyType: string
}>()

const emit = defineEmits(['refresh'])

const clearOfflineProxies = () => {
  fetch('../api/proxies?status=offline', {
    method: 'DELETE',
    credentials: 'include',
  })
    .then((res) => {
      if (res.ok) {
        ElMessage({
          message: '成功清除离线代理！',
          type: 'success',
        })
        emit('refresh')
      } else {
        ElMessage({
          message: '离线代理清除失败: ' + res.status + ' ' + res.statusText,
          type: 'warning',
        })
      }
    })
    .catch((err) => {
      ElMessage({
        message: '离线代理清除失败: ' + err.message,
        type: 'warning',
      })
    })
}
</script>
