<template>
  <div>
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
      <el-table-column label="名称" prop="name" sortable></el-table-column>
      <el-table-column label="端口" prop="port" sortable></el-table-column>
      <el-table-column label="连接数" prop="conns" sortable></el-table-column>
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
          <el-tag v-if="scope.row.status === 'online'" type="success"
            >{{ scope.row.status }}
          </el-tag>
          <el-tag v-else type="danger">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button
            type="primary"
            :name="scope.row.name"
            style="margin-bottom: 10px"
            @click="handleButton(scope.row)"
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
import ProxyViewExpand from './ProxyViewExpand.vue'
import { ref } from 'vue'
import Traffic from './Traffic.vue'

defineProps<{
  proxies: BaseProxy[]
  proxyType: string
}>()

const dialogVisible = ref(false)
const dialogVisibleName = ref('')

const formatTrafficIn = (row: BaseProxy, _: TableColumnCtx<BaseProxy>) => {
  return Humanize.fileSize(row.trafficIn)
}

const formatTrafficOut = (row: BaseProxy, _: TableColumnCtx<BaseProxy>) => {
  return Humanize.fileSize(row.trafficOut)
}

function handleButton(row: BaseProxy) {
  dialogVisibleName.value = row.name
  dialogVisible.value = true
}
</script>
