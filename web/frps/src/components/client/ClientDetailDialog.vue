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
      <el-button-group class="ml-4">
        <el-popconfirm
          title="确定卸载客户端吗，会导致不可恢复？"
          @confirm="handleUninstall"
        >
          <template #reference>
            <el-button type="danger" plain>卸载</el-button>
          </template>
        </el-popconfirm>
        <el-popconfirm title="确定重启客户端吗？" @confirm="handleReboot">
          <template #reference>
            <el-button type="warning" plain>重启</el-button>
          </template>
        </el-popconfirm>
        <el-button type="info" plain @click="handleTest">测试</el-button>
      </el-button-group>

      <el-row style="margin-top: 10px">
        <el-col :span="10">
          <el-input
            v-model="frpcTomlContent"
            :autosize="{ minRows: 2, maxRows: 23.5 }"
            placeholder="frpc configure file, can not be empty..."
            type="textarea"
          ></el-input>
        </el-col>
        <el-col :span="14">
          <el-card title="日志面板" class="log-container">
            <div>
              <div ref="logContainer" class="log-container">
                <div v-for="(log, index) in logs" :key="index" class="log-item">
                  {{ log }}
                </div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <template #footer>
      <el-button @click="handleClose">取消</el-button>
      <el-button type="danger" @click="handleConfirm">下发配置</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, defineExpose } from 'vue'
import { ElButton } from 'element-plus'
import { Client } from '../../utils/type.ts'
import { EventAwareSSEClient } from '../../utils/sseclient.ts'
import { showLoading, showTips } from '../../utils/utils.ts'

const logContainer = ref<HTMLDivElement | null>(null)
const logs = ref<string[]>([])
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

const handleTest = () => {
  addLog('wahahaha')
}
const addLog = (context: string): void => {
  const newLog = `${new Date().toLocaleString()}: ${context}\r\n`
  logs.value.unshift(newLog)
  // 滚动到顶部
  if (logContainer.value) {
    logContainer.value.scrollTop = 0
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
      addLog(JSON.stringify(data))
    })
    source.value.addEventListener('detail', (data) => {
      addLog(JSON.stringify(data))
    })
    source.value.addEventListener('client-info', (data) => {
      frpcTomlContent.value = data
      addLog(JSON.stringify(data))
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
    console.error('connectSSE err', e)
    addLog(JSON.stringify(e))
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

const handleReboot = () => {
  console.log('handleReboot', showClientDialog.value)
  fetchApi({ cmd: 'reboot' })
}

const handleUninstall = () => {
  console.log('handleUninstall', showClientDialog.value)
  fetchApi({ cmd: 'uninstall' })
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

const fetchApi = (data: any) => {
  data.frpId = client.value?.frpId
  data.secKey = client.value?.secKey
  const loading = showLoading('请求中...')
  fetch('../api/client/cmd', {
    credentials: 'include',
    method: 'POST',
    body: JSON.stringify(data),
  })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      console.log(json)
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

.log-container {
  height: auto;
  max-height: 500px;
  overflow-y: auto;
  margin-left: 20px;
}

.log-item {
  margin-bottom: 5px;
}
</style>
