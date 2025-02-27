<template>
  <div>
    <el-page-header
      :icon="null"
      style="width: 100%; margin-left: 30px; margin-bottom: 20px"
    >
      <template #title>
        <span>{{ proxyType }}</span>
      </template>
      <template #content> </template>
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
          <ProxyViewExpand :row="props.row" :proxyType="proxyType" />
        </template>
      </el-table-column>
      <el-table-column label="名称" prop="name" sortable> </el-table-column>
      <el-table-column label="端口" prop="port" sortable> </el-table-column>
      <el-table-column label="连接数" prop="conns" sortable> </el-table-column>
      <el-table-column
        label="入站流量"
        prop="trafficIn"
        :formatter="formatTrafficIn"
        sortable
      >
      </el-table-column>
      <el-table-column
        label="出站流量"
        prop="trafficOut"
        :formatter="formatTrafficOut"
        sortable
      >
      </el-table-column>
      <el-table-column label="客户端版本" prop="clientVersion" sortable>
      </el-table-column>
      <el-table-column label="状态" prop="status" sortable>
        <template #default="scope">
          <el-tag v-if="scope.row.status === 'online'" type="success">{{
            scope.row.status
          }}</el-tag>
          <el-tag v-else type="danger">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button
            type="primary"
            :name="scope.row.name"
            style="margin-bottom: 10px"
            @click="dialogVisibleName = scope.row.name; dialogVisible = true"
            >流量
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>

  <el-dialog
    v-model="dialogVisible"
    destroy-on-close="true"
    :title="dialogVisibleName"
    width="700px"
  >
    <Traffic :proxyName="dialogVisibleName" />
  </el-dialog>
</template>

<script setup lang="ts">
import * as Humanize from 'humanize-plus'
import type { TableColumnCtx } from 'element-plus'
import type { BaseProxy } from '../utils/proxy.js'
import { ElMessage } from 'element-plus'
import ProxyViewExpand from './ProxyViewExpand.vue'
import { ref } from 'vue'

defineProps<{
  proxies: BaseProxy[]
  proxyType: string
}>()

const emit = defineEmits(['refresh'])

const dialogVisible = ref(false)
const dialogVisibleName = ref('')

const formatTrafficIn = (row: BaseProxy, _: TableColumnCtx<BaseProxy>) => {
  return Humanize.fileSize(row.trafficIn)
}

const formatTrafficOut = (row: BaseProxy, _: TableColumnCtx<BaseProxy>) => {
  return Humanize.fileSize(row.trafficOut)
}

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

<style>
.el-page-header__title {
  font-size: 20px;
}
</style>
