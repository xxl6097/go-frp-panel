<template>
  <el-dialog
    :modal="true"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    width="80%"
    v-model="showClientDialog"
    :title="title"
    @closed="onClosed"
  >
    <div class="upgrade-popup-content">
      <el-input
        v-model="frpcTomlContent"
        autosize
        placeholder="frpc configure file, can not be empty..."
        type="textarea"
      ></el-input>
    </div>
    <template #footer>
      <el-button @click="handleClose">取消</el-button>
      <el-button type="primary" @click="handleConfirm">下发配置</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, defineExpose } from 'vue'
import { ElButton } from 'element-plus'
import { Client } from '../../utils/type.ts'
import { EventAwareSSEClient } from '../../utils/sseclient.ts'
import { showLoading, showTips } from '../../utils/utils.ts'

const showClientDialog = ref(false)
const client = ref<Client>()
const title = ref<string>()
const frpcTomlContent = ref<string>()
const source = ref<EventAwareSSEClient | null>()

const onClosed = () => {
  if (source.value) {
    console.log('close sse')
    source.value.close()
    source.value = null
  }
}

const openClientDetailDialog = (row: Client) => {
  client.value = row
  showClientDialog.value = true
  console.log('openClientDetailDialog', row)
  connectSSE(row)
}

const connectSSE = (row: Client) => {
  try {
    title.value = `${row?.devMac} (${row?.osType})`
    const sseUrl = `../api/client/sse?type=detail&frpId=${row.frpId}&secKey=${row.secKey}`
    console.log('connectSSE', sseUrl)
    source.value = new EventAwareSSEClient(sseUrl)
    source.value.addEventListener('connected', (data) => {
      console.log('connected:', data)
    })
    source.value.addEventListener('detail', (data) => {
      console.log('detail:', data)
      console.log('detail:', data.data)
    })
    source.value.addEventListener('client-info', (data) => {
      console.log('client-info:', data)
      frpcTomlContent.value = data
    })
    source.value.connect()
    //
    // console.log('connectSSE:', sseUrl)
    // source.value = new EventSource(sseUrl)
    // source.value.onmessage = (event) => {
    //   console.log('收到消息:', event.data)
    // }
    // source.value.onopen = (e) => {
    //   console.log('SSE连接已建立', e, source?.value?.readyState) // readyState=1表示连接正常
    // }
    // source.value.onerror = (e) => {
    //   console.log('onerror received a message', e)
    //   source.value = null
    // }
  } catch (e) {
    console.log('connectSSE err', e)
  }
}

// 暴露方法供父组件调用
defineExpose({
  openClientDialog: openClientDetailDialog,
})

const handleConfirm = () => {
  showClientDialog.value = false
  upgradeFrpcToml()
}

const handleClose = () => {
  showClientDialog.value = false
  console.log('handleClose', showClientDialog.value)
}

const upgradeFrpcToml = () => {
  const loading = showLoading('配置修改中...')
  const data = {
    toml: `${frpcTomlContent.value}`,
    frpId: client.value?.frpId,
    secKey: client.value?.secKey,
  }
  fetch('../api/client/config/upgrade', {
    credentials: 'include',
    method: 'POST',
    body: JSON.stringify(data),
  })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      showTips(json.code, json.msg)
    })
    .catch(() => {
      //showErrorTips('配置失败')
    })
    .finally(() => {
      loading.close()
    })
}

// checkVersion()
</script>
<style scoped>
.client-detail-dialog {
  width: 98%;
  height: 90%;
}

.upgrade-popup-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999; /* 设置较高的 z-index 值，确保在最顶部 */
}

.upgrade-popup {
  border-radius: 4px;
  width: 30%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.upgrade-popup-header {
  padding: 5px 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #e4e7ed;
}

.upgrade-popup-header h3 {
  line-height: 2.5;
  margin: 0;
}

.close-button {
  background: none;
  border: none;
  font-size: 30px;
  cursor: pointer;
}

.upgrade-popup-content {
  padding-left: 20px;
  padding-right: 20px;
}

.upgrade-popup-footer {
  padding: 10px 20px;
  text-align: right;
  border-top: 1px solid #e4e7ed;
}

.upgrade-popup-footer button {
  margin-left: 10px;
}

/* 亮色模式 */
@media (prefers-color-scheme: light) {
  .upgrade-popup-overlay {
    background-color: rgba(0, 0, 0, 0.5);
  }

  .upgrade-popup {
    background-color: white;
  }
}

/* 暗色模式 */
@media (prefers-color-scheme: dark) {
  .upgrade-popup-overlay {
    background-color: rgba(255, 255, 255, 0.1);
  }

  .upgrade-popup {
    background-color: #333;
    color: white;
  }

  .upgrade-popup-header {
    border-bottom: 1px solid #555;
  }

  .upgrade-popup-footer {
    border-top: 1px solid #555;
  }
}
</style>
