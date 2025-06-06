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
      <el-page-header :icon="null" style="width: 100%; margin-bottom: 20px">
        <template #title>
          <el-select
            v-model="selectValue"
            @change="handleSelectChange"
            placeholder="多客户端配置"
            clearable
            :fit-input-width="true"
            size="default"
            class="autoWidth"
          >
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item"
            />
          </el-select>
        </template>
        <template #content>
          <div style="display: flex">
            <el-button-group class="ml-4">
              <el-button type="warning" plain @click="handleRefrsh"
                >刷新
              </el-button>
              <el-popconfirm title="确定保存配置吗？" @confirm="handleChange">
                <template #reference>
                  <el-button type="warning" plain>保存</el-button>
                </template>
              </el-popconfirm>

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
              <el-button type="info" plain @click="handleTest"
                >清空日志
              </el-button>
              <el-popconfirm
                title="确定删除客户端吗，会导致不可恢复？"
                @confirm="handleDelete"
                v-if="
                  selectValue.label !== '' && selectValue.label !== 'frpc.toml'
                "
              >
                <template #reference>
                  <el-button type="danger" plain>删除客户端</el-button>
                </template>
              </el-popconfirm>
              <el-button
                type="primary"
                plain
                @click="newClientForm.showClientDialog = true"
                >新建客户端
              </el-button>
              <el-button type="warning" plain @click="handleNetwork"
                >网络信息
              </el-button>
              <el-button type="warning" plain @click="handleCheckVersion"
                >版本检测
              </el-button>
            </el-button-group>
          </div>
        </template>
        <template #extra></template>
      </el-page-header>

      <div style="margin-left: 10px" @click="handleDevelopment">
        <div>
          <span
            style="color: green; margin-right: 8px"
            v-if="profile?.ports !== ''"
            >允许范围：{{ profile?.ports }}
          </span>
          <span
            style="color: green; margin-right: 8px"
            v-if="profile?.domains !== ''"
            >允许域名：{{ profile?.domains }}
          </span>
          <span style="color: green" v-if="profile?.domains !== ''"
            >允许子域名：{{ profile?.subdomains }}</span
          >
          <span style="margin-left: 10px" v-if="isDevelopment">
            <el-input
              style="width: 200px"
              v-model="cmdstring"
              placeholder="请输入命令"
              @change="handleCMD"
            ></el-input>
          </span>
        </div>

        <el-row style="margin-top: 10px">
          <el-col :span="10">
            <el-input
              v-model="selectValue.content"
              :autosize="{ minRows: 2, maxRows: 23.5 }"
              placeholder="frpc configure file, can not be empty..."
              type="textarea"
            ></el-input>
          </el-col>
          <el-col :span="14">
            <el-card title="日志面板" class="log-container">
              <div>
                <div ref="logContainer" class="log-container">
                  <div
                    v-for="(log, index) in logs"
                    :key="index"
                    class="log-item"
                  >
                    <!--                    {{ log }}-->
                    <pre v-html="log"></pre>
                  </div>
                </div>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>
    </div>
  </el-dialog>

  <!--新建客户端-->
  <el-dialog v-model="newClientForm.showClientDialog" width="700">
    <template #header><span>创建客户端</span></template>
    <template #default>
      <el-form :model="newClientForm">
        <el-form-item label="配置文件名：" prop="label">
          <el-input
            v-model="newClientForm.data.label"
            placeholder="请输入toml配置文件名"
          />
        </el-form-item>

        <el-form-item prop="content">
          <el-input
            type="textarea"
            v-model="newClientForm.data.content"
            rows="13"
            placeholder="请在此输入toml格式配置内容"
          />
        </el-form-item>
      </el-form>
    </template>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="newClientForm.showClientDialog = false"
          >取消
        </el-button>
        <el-button type="primary" @click="handleNew">确定</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, defineExpose } from 'vue'
import { ElButton } from 'element-plus'
import { Client, FrpcConfiguration } from '../../utils/type.ts'
import { EventAwareSSEClient } from '../../utils/sseclient.ts'
import {
  showLoading,
  showMessageDialog,
  showMessageDialogWithCancel,
  showSucessTips,
  showTips,
  showWarmTips,
  syntaxHighlight,
} from '../../utils/utils.ts'

export interface Option {
  label: string
  value: string
  content: string
}

interface NewOption {
  showClientDialog: boolean
  data: Option
}

const logContainer = ref<HTMLDivElement | null>(null)
const logs = ref<string[]>([])
const showClientDialog = ref(false)
const client = ref<Client>()
const profile = ref<FrpcConfiguration>()
const title = ref<string>()
const cmdstring = ref<string>()
const source = ref<EventAwareSSEClient | null>()

const newClientForm = ref<NewOption>({
  showClientDialog: false,
  data: { label: '', value: '', content: '' },
})

const selectValue = ref<Option>({
  label: '',
  value: '',
  content: '',
})
const options = ref<Option[]>([])

const isDevelopment = ref(false)
const clickCount = ref(0)
let timer: number | null = null
const handleDevelopment = () => {
  // 首次点击启动定时器（1秒内有效）
  if (clickCount.value === 0) {
    timer = window.setTimeout(() => {
      clickCount.value = 0
      timer = null
    }, 1000)
  }

  clickCount.value++

  // 触发条件：5次点击
  if (clickCount.value === 5) {
    console.log('连续点击了5次！')
    // 执行目标操作（例如提交表单、跳转页面等）
    executeTargetAction()
    // 重置状态
    clickCount.value = 0
    if (timer) {
      clearTimeout(timer)
      timer = null
    }
  }
}
const executeTargetAction = () => {
  // 这里编写业务逻辑，例如调用接口或跳转页面
  showWarmTips('进入开发者模式')
  isDevelopment.value = true
}

const connectSSE = (row: Client) => {
  try {
    title.value = `${row?.devName} ${row?.devMac} ${row?.osType} ${row?.appVersion} `
    const sseUrl = `../api/client/sse?type=detail&frpId=${row.frpId}&secKey=${row.secKey}`
    console.log('connectSSE', sseUrl)
    source.value = new EventAwareSSEClient(sseUrl)
    source.value.addEventListener('connected', (data) => {
      addLog(JSON.stringify(data))
    })
    source.value.addEventListener('sse-connect', (data) => {
      addLog(JSON.stringify(data))
      console.log('sse-connect', data)
      if (data && client && client.value) {
        client.value.sseId = data.sseId
        console.log('sse-connect client', client.value)
      }
    })
    source.value.addEventListener('disconnect', (data) => {
      addLog(JSON.stringify(data))
      console.log('disconnect', data)
      if (data && data.frpId === client.value?.frpId) {
        const message = `
<font color="red">设备已经断开，信息如下：</font>
frpc连接ID：${client.value?.frpId}<br>
设备MAC：${client.value?.devMac}<br>
操作系统：${client.value?.osType}<br>
websocketID：${client.value?.secKey}<br>
设备IP：${client.value?.devIp}<br>
`
        showMessageDialog('设备警告⚠️', '确定', message)
        showClientDialog.value = false
      }
    })

    source.value.addEventListener('client-refresh', (data) => {
      console.log('config-refresh', data)
      if (data) {
        options.value = data
        if (options.value && options.value.length > 0) {
          const target = options.value.find(
            (item) => item.label === 'frpc.toml',
          )
          if (target) {
            selectValue.value = target
          }
        }

        const rawJson = JSON.stringify(data, null, 2)
        const highlightedJSON = syntaxHighlight(rawJson)
        console.log('config-refresh', data)
        addLog(highlightedJSON)
      }
    })

    source.value.addEventListener('client-version-check', (data) => {
      console.log('client-version-check', data)
      addLog(JSON.stringify(data))
      const type = Object.prototype.toString.call(data)
      if (type === '[object String]') {
        console.log('字符串', data)
        showSucessTips(data)
      } else if (type === '[object Object]') {
        console.log('对象', data)
        const complexVersionRegex =
          /(\d+(?:\.\d+){1,3})(?:-[a-zA-Z0-9.]+)?(?:\+[a-zA-Z0-9.]+)?/
        const text = data.releaseNotes
        const match = text.match(complexVersionRegex)
        const newVersionText = `发现新版本：${match?.[1] || ''} 需要升级吗？`

        if (data.patchUrl && data.patchUrl != '') {
          showMessageDialogWithCancel(
            '版本升级',
            newVersionText,
            '差量升级',
            '全量升级',
          )
            .then(() => {
              console.log('差量升级', data.patchUrl)
              handleConfirmUpgrade(data.patchUrl)
            })
            .catch(() => {
              console.log('全量升级', data.fullUrl)
              handleConfirmUpgrade(data.fullUrl)
            })
        } else {
          showMessageDialog('版本升级', '升级', newVersionText).then(() => {
            console.log('2--全量升级', data.fullUrl)
            handleConfirmUpgrade(data.fullUrl)
          })
        }
      }
    })
    source.value.addEventListener('client-version-upgrade', (data) => {
      console.log('client-version-upgrade', data)
      addLog(JSON.stringify(data))
    })
    source.value.addEventListener('network', (data) => {
      console.log('network', data)
      const rawJson = JSON.stringify(data, null, 2)
      const highlightedJSON = syntaxHighlight(rawJson)
      addLog(highlightedJSON)
    })
    source.value.addEventListener('cmd', (data) => {
      console.log('cmd', data)
      addLog(data)
    })
    source.value.connect()
  } catch (e) {
    console.error('connectSSE err', e)
    addLog(JSON.stringify(e))
  }
}

const handleSelectChange = (value: any) => {
  console.log('handleSelectChange---->', value)
  console.log('selectValue---->', selectValue)
  if (!value || value === '') {
    selectValue.value = {
      label: '',
      value: '',
      content: '',
    }
  }
}

const onClosed = () => {
  if (source.value) {
    console.log('close sse')
    source.value.close()
    source.value = null
  }
}

// 使用示例

const handleTest = () => {
  addLog('wahahaha')
  logs.value = []
}

const addLog = (context: string): void => {
  const newLog = `${new Date().toLocaleString()}: ${context}\r\n`
  logs.value.unshift(newLog)
  // 滚动到顶部
  if (logContainer.value) {
    logContainer.value.scrollTop = 0
  }
}

const openClientDetailDialog = (row: Client, p: FrpcConfiguration) => {
  console.log('打开对话框，row:', row, p)
  client.value = row
  profile.value = p
  showClientDialog.value = true
  connectSSE(row)
}

// 暴露方法供父组件调用
defineExpose({
  openClientDialog: openClientDetailDialog,
})

const handleReboot = () => {
  fetchApi('client-reboot', {})
}

const handleUninstall = () => {
  fetchApi('client-uninstall', {})
}

const handleNew = () => {
  fetchApi('client-new', {
    name: newClientForm.value.data.label,
    content: newClientForm.value.data.content,
  })
  newClientForm.value.showClientDialog = false
}

const handleConfirmUpgrade = (data: any) => {
  fetchApi('client-version-upgrade', { data: data })
}

const handleDelete = () => {
  fetchApi('client-delete', { name: selectValue.value.label })
}

const handleChange = () => {
  fetchApi('client-change', {
    name: selectValue.value.label,
    content: selectValue.value.content,
  })
}

const handleRefrsh = () => {
  fetchApi('client-refresh', {})
}
const handleCheckVersion = () => {
  fetchApi('client-version-check', {})
}
const handleNetwork = () => {
  fetchApi('network', {})
}
const handleCMD = () => {
  fetchApi('cmd', { data: cmdstring.value })
}

const fetchApi = (action: string, data: any) => {
  const body = {
    action: action,
    devIp: client.value?.devIp,
    devMac: client.value?.devMac,
    frpId: client.value?.frpId,
    secKey: client.value?.secKey,
    sseId: client.value?.sseId,
    data: data,
  }
  console.log('client', client.value)
  const loading = showLoading('请求中...')
  console.log('fetchApi', body)
  fetch('../api/client/cmd', {
    credentials: 'include',
    method: 'POST',
    body: JSON.stringify(body),
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
</script>
<style scoped>
.upgrade-popup-header h3 {
  line-height: 2.5;
  margin: 0;
}

.upgrade-popup-content {
  padding-left: 20px;
  padding-right: 20px;
}

.upgrade-popup-footer button {
  margin-left: 10px;
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

.autoWidth {
  width: auto;
  min-width: 250px; /* 初始最小宽度 */
  max-width: 400px; /* 初始最小宽度 */
  margin-left: 10px;
}
</style>
