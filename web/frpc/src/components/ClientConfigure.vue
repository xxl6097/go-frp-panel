<template>
  <div>
    <el-page-header :icon="null" style="width: 100%; margin-bottom: 20px">
      <template #title>
        <el-select
          v-model="selectValue"
          @change="handleSelectChange"
          placeholder="请选择配置文件"
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
        <div class="flex items-center">
          <el-button type="primary" @click="upload" :loading="uploading" plain
            >更新</el-button
          >
          <el-button type="success" @click="refresh" :loading="loading" plain
            >刷新</el-button
          >
          <el-button
            type="warning"
            @click="handleShowNewFrpc"
            :loading="loading"
            plain
            >新建客户端</el-button
          >
          <div v-if="selectValue !== ''">
            <el-popconfirm title="确定删除客户端吗？" @confirm="deleteClient">
              <template #reference>
                <el-button type="danger" :loading="loading" plain
                  >删除客户端</el-button
                >
              </template>
            </el-popconfirm>
          </div>
          <el-button type="warning" @click="drawer = true" plain
            >新建代理</el-button
          >
        </div></template
      >
      <template #extra> </template>
    </el-page-header>

    <el-input
      type="textarea"
      autosize
      style="margin-left: 10px"
      v-model="textarea"
      placeholder="frpc configure file, can not be empty..."
    ></el-input>
  </div>

  <!--新建客户端-->
  <el-dialog v-model="newClientFormVisible" width="700">
    <template #header><span>创建客户端</span> </template>
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
          >确定</el-button
        >
      </div>
    </template>
  </el-dialog>

  <el-drawer
    v-model="drawer"
    title="I am the title"
    :with-header="true"
    direction="rtl"
    size="40%"
  >
    <template #header>
      <h1>新建代理</h1>
    </template>
    <template #default>
      <el-tabs type="border-card" :stretch="true">
        <el-tab-pane label="TCP">
          <el-form
            ref="ruleFormRef"
            :model="proxyForm"
            :rules="proxyRules"
            label-position="top"
          >
            <el-form-item label="代理名称：" prop="name">
              <el-input v-model="proxyForm.name" placeholder="代理名称">
                <template #append>
                  <el-button type="primary">生成</el-button>
                </template>
              </el-input>
            </el-form-item>

            <el-form-item required>
              <el-col :span="14">
                <el-form-item label="内网地址" prop="localIP">
                  <el-input
                    v-model="proxyForm.localIP"
                    placeholder="127.0.0.1"
                  />
                </el-form-item>
              </el-col>
              <el-col class="text-center" :span="2">
                <span class="text-gray-500"></span>
              </el-col>
              <el-col :span="8">
                <el-form-item label="内网端口" prop="localPort">
                  <el-col :span="15">
<!--                    <el-input-number-->
<!--                      v-model="proxyForm.localPort"-->
<!--                      controls-position="right"-->
<!--                      placeholder="请输入内网地端口"-->
<!--                    />-->
                    <el-select
                      v-model.number="proxyForm.localPort"
                      placeholder="请输入端口"
                      filterable
                      clearable
                      allow-create
                    >
                      <el-option
                        v-for="item in options1"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                      />
                    </el-select>
                  </el-col>

                  <el-col class="text-center" :span="1">
                    <span class="text-gray-50"></span>
                  </el-col>
                  <el-col :span="8">
                    <el-button type="primary" plain>内网端口</el-button>
                  </el-col>
                </el-form-item>
              </el-col>
            </el-form-item>

            <el-form-item label="外网端口：" prop="remotePort">
              <el-input-number
                v-model="proxyForm.remotePort"
                controls-position="right"
                placeholder="请输入外网端口"
              />
            </el-form-item>

            <el-form-item>
              <el-select
                v-model.number="value1"
                placeholder="Select"
                filterable
                clearable
                allow-create
                style="width: 240px"
              >
                <el-option
                  v-for="item in options1"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="UDP">udp</el-tab-pane>
        <el-tab-pane label="HTTP">http</el-tab-pane>
        <el-tab-pane label="HTTPS">https</el-tab-pane>
        <el-tab-pane label="STCP">stcp</el-tab-pane>
        <el-tab-pane label="SUDP">sudp</el-tab-pane>
        <el-tab-pane label="TCPMUX">tcpmux</el-tab-pane>
      </el-tabs>
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="null">cancel</el-button>
        <el-button type="primary" @click="null">confirm</el-button>
      </div>
    </template>
  </el-drawer>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import {
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
const drawer = ref(false)
const newClientFormVisible = ref(false)
const loading = ref<boolean>(false)
const uploading = ref<boolean>(false)
const textarea = ref('')
const selectValue = ref('')
const options = ref<Option[]>([])
const value1 = ref('')
const options1 = [
  {
    value: 'Option1',
    label: 'Option1',
  },
  {
    value: 'Option2',
    label: 'Option2',
  },
  {
    value: 'Option3',
    label: 'Option3',
  },
  {
    value: 'Option4',
    label: 'Option4',
  },
  {
    value: 'Option5',
    label: 'Option5',
  },
]
const newClientForm = ref({
  name: '',
  toml: '',
})

const proxyForm = ref({
  name: '',
  localIP: '',
  localPort: 0,
  remotePort: 0,
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

const handleShowNewFrpc = () => {
  newClientFormVisible.value = true
  newClientForm.value.name = `${getTimestamp()}.toml`
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
  const body = JSON.stringify(newClientForm)
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

////
const handleSelectChange = (value: string) => {
  console.log('---->', value)
  if (value === '') {
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
      if (json.code === 0) {
        textarea.value = json.data
      } else {
        showErrorTips(json.msg)
      }
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
  if (selectValue.value === '') {
    fetchData()
  } else {
    fetchConfig()
  }
  fetchListData()
}

const deleteClient = () => {
  if (selectValue.value !== '') {
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
  if (selectValue.value === '') {
    uploadConfig()
  } else {
    fetchUpload()
  }
}

const fetchData = () => {
  fetch('/api/config', { credentials: 'include' })
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

      fetch('/api/config', {
        credentials: 'include',
        method: 'PUT',
        body: textarea.value,
      })
        .then(() => {
          fetch('/api/reload', { credentials: 'include' })
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
</style>
