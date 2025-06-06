<template>
  <div>
    <el-page-header :icon="null" style="width: 100%; margin-bottom: 20px">
      <template #title>
        <el-select
          v-model="selectValue"
          @change="handleSelectChange"
          placeholder="多客户端配置"
          clearable
          :fit-input-width="true"
          size="default"
          class="autoWidth1"
        >
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </template>
      <template #content>
        <div style="display: flex">
          <el-button type="primary" @click="upload" :loading="uploading" plain
            >更新
          </el-button>
          <el-button type="success" @click="refresh" :loading="loading" plain
            >刷新
          </el-button>

          <el-button-group class="ml-4">
            <el-button
              type="warning"
              @click="handleShowNewFrpc"
              :loading="loading"
              plain
              >新建客户端
            </el-button>
            <el-popconfirm
              title="确定删除客户端吗？"
              @confirm="deleteClient"
              v-if="selectValue !== '' && selectValue !== undefined"
            >
              <template #reference>
                <el-button type="danger" :loading="loading" plain
                  >删除客户端
                </el-button>
              </template>
            </el-popconfirm>

            <el-button type="warning" @click="handleShowNewProxyDrawer" plain
              >新建代理
            </el-button>

            <el-button type="warning" plain @click="handleImportConfigClick"
              >导入配置
            </el-button>
            <el-button type="warning" plain @click="handleExportConfig"
              >导出配置
            </el-button>
            <el-upload
              :http-request="handleImportConfigs"
              :limit="1"
              accept=".zip"
            >
              <template #trigger>
                <el-button ref="innerBtn" style="display: none"></el-button>
              </template>
            </el-upload>
          </el-button-group>
        </div>
      </template>
      <template #extra></template>
    </el-page-header>

    <div style="margin-left: 10px">
      <div>
        <span
          style="color: green; margin-right: 8px"
          v-if="meta && meta?.ports !== ''"
          >允许范围：{{ meta?.ports }}
        </span>
        <span
          style="color: green; margin-right: 8px"
          v-if="meta && meta?.domains !== ''"
          >允许域名：{{ meta?.domains }}
        </span>
        <span style="color: green" v-if="meta && meta?.domains !== ''"
          >允许子域名：{{ meta?.subdomains }}</span
        >
      </div>
      <el-input
        type="textarea"
        autosize
        v-model="textarea"
        placeholder="frpc configure file, can not be empty..."
      ></el-input>
    </div>
  </div>

  <!--新建客户端-->
  <el-dialog v-model="newClientFormVisible" width="700">
    <template #header><span>创建客户端</span></template>
    <template #default>
      <el-form ref="ruleFormRef" :model="newClientForm" :rules="rules">
        <el-form-item label="配置文件名：" prop="name">
          <el-input
            v-model="newClientForm.name"
            placeholder="请输入toml配置文件名"
          />
        </el-form-item>

        <el-form-item prop="toml">
          <el-input
            type="textarea"
            rows="13"
            v-model="newClientForm.toml"
            placeholder="请在此输入toml格式配置内容"
          />
        </el-form-item>
      </el-form>
      <el-upload :http-request="uploadToml" :limit="1">
        <template #trigger>
          <el-link type="primary">上传toml配置文件</el-link>
        </template>
      </el-upload>
    </template>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="newClientFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm(ruleFormRef)"
          >确定
        </el-button>
      </div>
    </template>
  </el-dialog>

  <!--  新建代理-->
  <el-drawer v-model="drawer" :with-header="true" direction="rtl" size="35%">
    <template #header>
      <h1>新建代理</h1>
    </template>
    <div class="demo-drawer__content">
      <el-tabs type="border-card" v-model="tabIndex">
        <el-tab-pane label="tcp" name="tcp">
          <el-form
            ref="ruleFormRef"
            :model="proxyForm"
            :rules="proxyRules"
            label-position="top"
          >
            <el-form-item label="代理名称：" prop="name">
              <el-input v-model="proxyForm.name" placeholder="代理名称">
                <template #append>
                  <el-button type="primary" @click="handleGenProxyName('tcp')"
                    >生成
                  </el-button>
                </template>
              </el-input>
            </el-form-item>

            <el-form-item required>
              <el-col :span="14">
                <el-form-item label="内网地址" prop="localIP">
                  <el-select
                    v-model="proxyForm.localIP"
                    placeholder="127.0.0.1"
                    @change="handleProxyLocalPort()"
                    loading-text="局域网主机扫描中..."
                    filterable
                    :loading="locaIPScanning"
                    clearable
                    allow-create
                  >
                    <el-option
                      v-for="item in ips"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col class="text-center" :span="2">
                <span class="text-gray-500"></span>
              </el-col>
              <el-col :span="8">
                <el-form-item label="内网端口" prop="localPort">
                  <el-select
                    v-model.number="proxyForm.localPort"
                    placeholder="请输入端口"
                    :loading="portLoading"
                    loading-text="端口扫描中..."
                    filterable
                    clearable
                    allow-create
                  >
                    <el-option
                      v-for="item in ports"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-form-item>

            <el-form-item label="外网端口：" prop="remotePort">
              <div style="display: flex">
                <el-select
                  style="width: 150px"
                  v-model.number="proxyForm.remotePort"
                  placeholder="请输入外网端口"
                  @change="handleProxyRemotePortCheck"
                  filterable
                  clearable
                  allow-create
                >
                  <el-option
                    v-for="item in remoteports"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  >
                  </el-option>
                </el-select>
                <div
                  v-if="
                    proxyForm.remotePort != null && proxyForm.remotePort > 0
                  "
                  style="
                    align-content: center;
                    align-items: center;
                    margin-left: 5px;
                  "
                >
                  <svg
                    v-if="checkPortErr.code === 0"
                    viewBox="0 0 1024 1024"
                    width="16"
                    height="16"
                  >
                    <path
                      fill="#67c23a"
                      d="M512 64a448 448 0 1 1 0 896 448 448 0 0 1 0-896zm-55.808 536.384-122.88-122.88a38.4 38.4 0 0 0-54.336 54.336l153.6 153.6a38.4 38.4 0 0 0 54.336 0l307.2-307.2a38.4 38.4 0 0 0-54.336-54.336L456.192 600.384z"
                    />
                  </svg>
                  <svg
                    viewBox="0 0 16 16"
                    width="16"
                    height="16"
                    v-if="checkPortErr.code !== 0"
                  >
                    <path
                      d="M2 2 L12 12 M2 12 L12 2"
                      stroke="red"
                      stroke-width="2"
                      stroke-linecap="round"
                      fill="none"
                    />
                  </svg>
                  <el-text v-if="checkPortErr.code !== 0"
                    >{{ checkPortErr.msg }}
                  </el-text>
                </div>
              </div>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="udp" name="udp">开发中...</el-tab-pane>
        <el-tab-pane label="http" name="http">开发中...</el-tab-pane>
        <el-tab-pane label="https" name="https">开发中...</el-tab-pane>
        <el-tab-pane label="stcp" name="stcp">开发中...</el-tab-pane>
        <el-tab-pane label="xtcp" name="xtcp">开发中...</el-tab-pane>
        <el-tab-pane label="sudp" name="sudp">开发中...</el-tab-pane>
        <el-tab-pane label="tcpmux" name="tcpmux">开发中...</el-tab-pane>
      </el-tabs>
      <div class="demo-drawer__footer">
        <el-button @click="null">取消</el-button>
        <el-button
          type="primary"
          @click="handleNewTCPProxy"
          :loading="proxyAddSaveLoading"
          >{{ proxyAddSaveLoading ? '保存中 ...' : '保存' }}
        </el-button>
      </div>
    </div>
  </el-drawer>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import {
  ElButton,
  ElMessage,
  ElMessageBox,
  FormInstance,
  FormRules,
} from 'element-plus'
import {
  downloadByPost,
  FrpcConfiguration,
  getProxyName,
  getTimestamp,
  put,
  showErrorTips,
  showInfoTips,
  showLoading,
  showSucessTips,
} from '../utils/utils.ts'

interface Option {
  value: string
  label: string
}

const innerBtn = ref<InstanceType<typeof ElButton>>()
const drawer = ref(false)
const checkPortErr = ref({
  code: -1,
  msg: '暂未检测',
})
const tabIndex = ref('tcp')
const locaIPScanning = ref(false)
const proxyAddSaveLoading = ref(false)
const portLoading = ref(false)
const newClientFormVisible = ref(false)
const loading = ref<boolean>(false)
const uploading = ref<boolean>(false)
const textarea = ref('')
const meta = ref<FrpcConfiguration>()
const selectValue = ref('')
const options = ref<Option[]>([])
const ports = ref<Option[]>([])
const remoteports = ref<Option[]>([])
const ips = ref<Option[]>([])
const newClientForm = ref({
  name: '',
  toml: '',
})

const proxyForm = ref({
  name: '',
  localIP: '',
  type: '',
  localPort: null as number | null,
  remotePort: null as number | null,
})

const proxyRules = reactive<FormRules>({
  name: [
    {
      required: true,
      message: '请输入代理名称',
      trigger: 'blur',
    },
  ],
  localIP: [
    {
      required: true,
      message: '请输入代理本地地址',
      trigger: 'blur',
    },
  ],
  localPort: [
    {
      required: true,
      message: '请输入代理本地端口',
      trigger: 'blur',
    },
  ],
  remotePort: [
    {
      required: true,
      message: '请输入代理远程端口',
      trigger: 'blur',
    },
  ],
})

const ruleFormRef = ref<FormInstance>()
const rules = reactive<FormRules>({
  name: [
    {
      required: true,
      message: '请输入配置文件名',
      trigger: 'blur',
    },
  ],
  toml: [
    {
      required: true,
      message: '请输入配置内容',
      trigger: 'blur',
    },
  ],
})

const scanLocalIp = () => {
  locaIPScanning.value = true
  // 使用 fetch 发送请求
  fetch(`../api/proxy/ips`, {
    method: 'GET',
  })
    .then((response) => {
      return response.json()
    })
    .then((json) => {
      console.log('ips', json)
      if (json.code === 0) {
        ips.value = json.data
      }
    })
    .finally(() => {
      locaIPScanning.value = false
    })
}

const handleProxyLocalPort = () => {
  if (proxyForm.value.localIP === '') {
    //showWarmTips('请先输入内网地址')
    return
  }
  portLoading.value = true
  ports.value = []
  // 使用 fetch 发送请求
  fetch(`../api/proxy/ports?localIP=${proxyForm.value.localIP}`, {
    method: 'GET',
  })
    .then((response) => {
      return response.json()
    })
    .then((json) => {
      console.log('ports', json)
      if (json.code === 0) {
        ports.value = json.data
      }
    })
    .finally(() => {
      portLoading.value = false
    })
}

const handleProxyRemotePortCheck = () => {
  ports.value = []
  const name = selectValue.value
  fetch(
    `../api/proxy/port/check?name=${name}&port=${proxyForm.value.remotePort}`,
    {
      method: 'GET',
    },
  )
    .then((response) => {
      return response.json()
    })
    .then((json) => {
      console.log('ports', json)
      checkPortErr.value = json
      if (json.code === 0) {
        showSucessTips(json.msg)
      } else {
        showErrorTips(json.msg)
      }
    })
    .catch((err) => {
      checkPortErr.value = {
        code: -1,
        msg: err.toString(),
      }
    })
}

const fetchRemotePorts = (name: string) => {
  fetch(`../api/proxy/remote/ports?name=${name}`, {
    method: 'GET',
  })
    .then((response) => {
      return response.json()
    })
    .then((json) => {
      console.log('fetchRemotePorts', json)
      if (json.code === 0) {
        remoteports.value = json.data
      }
    })
}

const handleGenProxyName = (prefix: string) => {
  proxyForm.value.name = getProxyName(prefix)
}

const handleShowNewFrpc = () => {
  newClientFormVisible.value = true
  newClientForm.value.name = `${getTimestamp()}.toml`
}

const handleImportConfigClick = () => {
  innerBtn.value?.$el.click()
}

const handleExportConfig = () => {
  console.log('配置导出中')
  downloadByPost('../api/client/config/export', {}).finally(() => {})
}

const handleImportConfigs = (options: any) => {
  const { file } = options
  const formData = new FormData()
  formData.append('file', file)
  const loading = showLoading('用户导入中...')
  // 使用 fetch 发送请求
  fetch('../api/client/config/import', {
    method: 'POST',
    body: formData,
  })
    .then((response) => {
      return response.json()
    })
    .finally(() => {
      loading.close()
    })
}

const handleShowNewProxyDrawer = () => {
  drawer.value = true
  proxyForm.value = {
    name: '',
    localIP: '',
    type: '',
    localPort: null as number | null,
    remotePort: null as number | null,
  }
  //获取远程可使用端口列表
  fetchRemotePorts(selectValue.value)
  //获取局域网活动主机IP
  scanLocalIp()
}

const handleNewTCPProxy = () => {
  const data = {
    type: tabIndex.value,
    name: proxyForm.value.name,
    localIP: proxyForm.value.localIP,
    localPort: proxyForm.value.localPort,
    remotePort: proxyForm.value.remotePort,
  }
  proxyAddSaveLoading.value = true
  fetch(`../api/proxy/tcp/add?name=${selectValue.value}`, {
    method: 'PUT',
    body: JSON.stringify(data),
  })
    .then((response) => {
      return response.json()
    })
    .then((json) => {
      console.log('fetchRemotePorts', json)
      if (json.code === 0) {
        remoteports.value = json.data
      }
      showInfoTips(json.msg)
    })
    .finally(() => {
      proxyAddSaveLoading.value = false
      drawer.value = false
    })
}

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      console.log('submit!')
      handleNewFrpcClient()
    } else {
      console.log('error submit!', fields)
    }
  })
}

const handleNewFrpcClient = () => {
  const body = JSON.stringify(newClientForm.value)
  put('客户端创建中...', '../api/client/create', body).finally(() => {
    newClientFormVisible.value = false
    fetchListData()
  })
}

// 自定义上传函数
const uploadToml = (options: any) => {
  const { file } = options
  const formData = new FormData()
  formData.append('file', file)
  const loading = showLoading('客户端创建中...')
  // 使用 fetch 发送请求
  fetch('../api/client/create', {
    method: 'POST',
    body: formData,
  })
    .then((response) => {
      return response.json()
    })
    .then((data) => {
      // 上传成功的回调
      options.onSuccess(data)
    })
    .catch((error) => {
      // 上传失败的回调
      options.onError(error)
    })
    .finally(() => {
      loading.close()
      newClientFormVisible.value = false
      setTimeout(function () {
        window.location.reload()
      }, 1000)
    })
}

const handleSelectChange = (value: string) => {
  console.log('handleSelectChange---->', value)
  meta.value = undefined
  if (value === '' || value === undefined) {
    fetchData()
    return
  }
  fetchConfig()
}
const fetchListData = () => {
  fetch('../api/client/list', { credentials: 'include' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      console.log('list', json)
      if (json.code === 0) {
        options.value = json.data
      }
    })
}
const fetchConfig = () => {
  const name = selectValue.value
  fetch(`../api/client/config/get?name=${name}`, { credentials: 'include' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      console.log('fetchConfig', json)
      if (json.code === 0 && json.data) {
        textarea.value = json.data.toml
        const cfg = {
          user: json.data.meta.user,
          token: json.data.meta.token,
          count: json.data.meta.count,
          comment: json.data.meta.comment,
          ports: json.data.meta.ports.join(','),
          domains: json.data.meta.domains.join(','),
          subdomains: json.data.meta.subdomains.join(','),
          enable: json.data.meta.enable,
          id: json.data.meta.id,
        }
        meta.value = cfg as FrpcConfiguration
        console.log('meta', meta.value)
      } else {
        showErrorTips(json.msg)
        meta.value = undefined
      }
    })
    .catch(() => {
      meta.value = undefined
    })
    .finally(() => {
      loading.value = false
    })
}
const fetchUpload = () => {
  const data = {
    name: selectValue.value,
    toml: textarea.value,
  }
  const body = JSON.stringify(data)
  fetch(`../api/client/config/set`, {
    credentials: 'include',
    method: 'POST',
    body: body,
  })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      showInfoTips(json.msg)
    })
    .finally(() => {
      uploading.value = false
    })
}
const refresh = () => {
  loading.value = true
  if (selectValue.value === undefined || selectValue.value === '') {
    fetchData()
  } else {
    fetchConfig()
  }
  fetchListData()
}

const deleteClient = () => {
  if (selectValue.value !== '' && selectValue.value !== undefined) {
    loading.value = true
    fetch(`../api/client/delete?name=${selectValue.value}`, {
      credentials: 'include',
      method: 'DELETE',
    })
      .then((res) => {
        return res.json()
      })
      .then((json) => {
        if (json.code === 0) showSucessTips(json.msg)
      })
      .catch(() => {
        ElMessage({
          showClose: true,
          message: 'delete frpc failed!',
          type: 'warning',
        })
      })
      .finally(() => {
        loading.value = false
        selectValue.value = ''
        fetchListData()
        fetchData()
      })
  }
}
const upload = () => {
  uploading.value = true
  if (selectValue.value === undefined || selectValue.value === '') {
    uploadConfig()
  } else {
    fetchUpload()
  }
}

const fetchData = () => {
  console.log('fetchData.meta', meta.value)
  fetch('../api/config', { credentials: 'include' })
    .then((res) => {
      return res.text()
    })
    .then((text) => {
      textarea.value = text
    })
    .catch(() => {
      ElMessage({
        showClose: true,
        message: 'Get configure content from frpc failed!',
        type: 'warning',
      })
    })
    .finally(() => {
      loading.value = false
    })
}
const uploadConfig = () => {
  ElMessageBox.confirm(
    'This operation will upload your frpc configure file content and hot reload it, do you want to continue?',
    'Notice',
    {
      confirmButtonText: 'Yes',
      cancelButtonText: 'No',
      type: 'warning',
    },
  )
    .then(() => {
      if (textarea.value == '') {
        ElMessage({
          message: 'Configure content can not be empty!',
          type: 'warning',
        })
        return
      }

      fetch('../api/config', {
        credentials: 'include',
        method: 'PUT',
        body: textarea.value,
      })
        .then(() => {
          fetch('../api/reload', { credentials: 'include' })
            .then(() => {
              ElMessage({
                type: 'success',
                message: 'Success',
              })
            })
            .catch((err) => {
              ElMessage({
                showClose: true,
                message: 'Reload frpc configure file error, ' + err,
                type: 'warning',
              })
            })
        })
        .catch(() => {
          ElMessage({
            showClose: true,
            message: 'Put config to frpc and hot reload failed!',
            type: 'warning',
          })
        })
        .finally(() => {
          uploading.value = false
        })
    })
    .catch(() => {
      ElMessage({
        message: 'Canceled',
        type: 'info',
      })
    })
    .finally(() => {
      uploading.value = false
    })
}

fetchData()
fetchListData()
</script>

<style>
#head {
  margin-bottom: 30px;
}

.autoWidth1 {
  width: auto;
  min-width: 250px; /* 初始最小宽度 */
  max-width: 400px; /* 初始最小宽度 */
  margin-left: 10px;
}

.success-icon {
  color: #67c23a; /* Element Plus 成功色 */
  font-size: 16px;
  margin-right: 8px; /* 调整图标与输入框右侧间距 */
}
</style>
