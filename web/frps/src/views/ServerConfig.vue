<template>
  <div class="server-config">
    <div class="page-head">
      <div>
        <h2 class="page-title">服务器配置</h2>
        <p class="page-subtitle">编辑 frps 服务端配置文件（TOML）</p>
      </div>
      <div class="page-actions">
        <el-button @click="fetchData">刷新数据</el-button>
        <el-button type="primary" @click="uploadConfig">更新配置</el-button>
      </div>
    </div>

    <el-card shadow="never" class="config-card">
      <el-input
        v-model="textarea"
        :autosize="{ minRows: 18 }"
        placeholder="frps 配置文件内容，不能为空……"
        type="textarea"
      />
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { ElMessageBox } from 'element-plus'
import {
  showErrorTips,
  showLoading,
  showTips,
  showWarmTips,
} from '../utils/utils.ts'

const textarea = ref('')

const fetchData = () => {
  fetch('../api/server/config/get', { credentials: 'include' })
    .then((res) => res.text())
    .then((text) => {
      textarea.value = text
    })
    .catch(() => {
      showErrorTips('获取配置失败')
    })
}

const uploadConfig = () => {
  ElMessageBox.confirm(
    '这个操作将更新frps服务的配置信息，确定要操作吗？',
    '注意',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    },
  )
    .then(() => {
      if (textarea.value == '') {
        showWarmTips('配置内容不能为空')
        return
      }
      const loading = showLoading('配置修改中...')
      fetch('../api/server/config/set', {
        credentials: 'include',
        method: 'PUT',
        body: textarea.value,
      })
        .then((res) => res.json())
        .then((json) => {
          showTips(json.code, json.msg)
        })
        .catch(() => {})
        .finally(() => {
          loading.close()
          setTimeout(function () {
            window.location.reload()
          }, 1000)
        })
    })
    .catch(() => {
      showErrorTips('配置失败')
    })
}

fetchData()
</script>

<style scoped>
.page-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 16px;
  margin-bottom: 20px;
  flex-wrap: wrap;
}

.page-actions {
  display: flex;
  gap: 8px;
}

.config-card :deep(.el-textarea__inner) {
  font-family:
    ui-monospace, SFMono-Regular, Menlo, Consolas, 'Liberation Mono', monospace;
  font-size: 13px;
  line-height: 1.6;
}
</style>
