<template>
  <div id="app">
    <header class="grid-content header-color">
      <div class="header-content">
        <div class="brand">
          <a href="#">{{ title }}</a>
        </div>
        <div class="dark-switch">
          <div class="dark-reboot">
            <el-dropdown placement="bottom" split-button plain @click="restart">
              重启
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item @click="dialogFormVisible = true"
                    >升级服务
                  </el-dropdown-item>
                  <el-dropdown-item @click="showlog">查看日志</el-dropdown-item>
                  <el-dropdown-item @click="showVersion"
                    >查看版本</el-dropdown-item
                  >
                  <el-dropdown-item @click="newClientFormVisible = true"
                    >新建客户端</el-dropdown-item
                  >
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
          <el-switch
            v-model="darkmodeSwitch"
            active-text="深色"
            inactive-text="浅色"
            inline-prompt
            style="
              --el-switch-on-color: #444452;
              --el-switch-off-color: #589ef8;
            "
            @change="toggleDark"
          />
        </div>
      </div>
    </header>
    <section>
      <el-row>
        <el-col id="side-nav" :xs="24" :md="4">
          <el-menu
            default-active="1"
            mode="vertical"
            theme="light"
            router="false"
            @select="handleSelect"
          >
            <el-menu-item index="/">客户端信息</el-menu-item>
            <el-menu-item index="/configure">配置</el-menu-item>
            <el-menu-item index="">帮助</el-menu-item>
          </el-menu>
        </el-col>

        <el-col :xs="24" :md="20">
          <div id="content">
            <router-view></router-view>
          </div>
        </el-col>
      </el-row>
    </section>
    <footer></footer>
  </div>

  <el-dialog
    v-model="dialogFormVisible"
    align-center
    title="程序升级"
    width="500"
  >
    <el-input
      v-model="form.binUrl"
      autocomplete="off"
      placeholder="请输入程序Url地址～"
    />

    <template #footer>
      <div class="dialog-footer">
        <el-upload
          class="upload-demo"
          :http-request="customUpload"
          :limit="1"
        >
          <template #trigger>
            <el-button type="primary" :disabled="form.binUrl.length > 0"
              >上传文件升级</el-button
            >
          </template>
          <!-- 添加额外按钮 -->
          <el-button style="margin-left: 10px" type="danger" @click="upgrade">
            文件url升级
          </el-button>
        </el-upload>
      </div>
    </template>
  </el-dialog>

  <!-- 弹窗表格 -->
  <el-dialog v-model="versionDialogVisible" width="30%" title="版本信息">
    <!-- 空标题插槽 -->
    <template #title></template>
    <el-descriptions :column="1" :size="size" border>
      <el-descriptions-item width="100">
        <template #label>
          <div class="cell-item">软件名称</div>
        </template>
        {{ version.appName }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <div class="cell-item">软件版本</div>
        </template>
        {{ version.appVersion }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <div class="cell-item">编译时间</div>
        </template>
        {{ version.buildTime }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <div class="cell-item">frpc版本号</div>
        </template>
        {{ version.frpcVersion }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <div class="cell-item">git版本</div>
        </template>
        {{ version.gitRevision }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <div class="cell-item">go编译版本</div>
        </template>
        {{ version.goVersion }}
      </el-descriptions-item>
    </el-descriptions>
  </el-dialog>

  <!--新建客户端-->
  <el-dialog v-model="newClientFormVisible" title="创建客户端" width="700">
    <el-form ref="ruleFormRef" :model="newClientForm" :rules="rules">
      <el-form-item label="配置文件名：" prop="toml">
        <el-input
          v-model="newClientForm.name"
          placeholder="请输入toml配置文件名"
        />
      </el-form-item>

      <el-form-item prop="toml">
        <el-input
          type="textarea"
          rows="18"
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
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="newClientFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm(ruleFormRef)"
          >确定</el-button
        >
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useDark, useToggle } from '@vueuse/core'
import {
  showLoading,
  showWarmDialog,
  showErrorTips,
  showTips,
  showWarmTips,
  put,
} from './utils/utils.ts'
import { ComponentSize, FormInstance, FormRules } from 'element-plus'

const newClientForm = ref({
  name: '',
  toml: '',
})
const size = ref<ComponentSize>('default')

const newClientFormVisible = ref(false)
const versionDialogVisible = ref(false)
const version = ref({
  description: '',
  frpcVersion: '',
  buildTime: '',
  appVersion: '',
  appName: '',
  gitRevision: '',
  goVersion: '',
})
const title = ref<string>('Frpc')
const isDark = useDark()
const darkmodeSwitch = ref(isDark)
const toggleDark = useToggle(isDark)
const dialogFormVisible = ref(false)
const form = ref({
  binUrl: '',
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

const handleSelect = (key: string) => {
  if (key == '') {
    window.open('https://github.com/xxl6097/go-frp-panel')
  }
}
const showlog = () => {
  const host = window.origin
  window.open(`${host}/log/`)
}
// const showBox = (content: any) => {
//   ElMessageBox.alert(content, {
//     confirmButtonText: 'OK',
//   })
// }

const showVersion = () => {
  // versionDialogVisible.value = true
  fetch('../api/version', { credentials: 'include', method: 'GET' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      version.value = json
      versionDialogVisible.value = true
    })
    .catch(() => {
      showErrorTips('失败')
    })
}

const fetchData = () => {
  fetch('../api/version', { credentials: 'include', method: 'GET' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      if (json) {
        title.value = `Frpc客户端 v${json.frpcVersion}`
        document.title = title.value
      }
    })
    .catch(() => {
      showErrorTips('重启失败')
    })
}

const upgrade = () => {
  if (form.value.binUrl.length > 0) {
    const loading = showLoading('程序升级中...')
    dialogFormVisible.value = false
    fetch('../api/upgrade', {
      credentials: 'include',
      method: 'PUT',
      body: form.value.binUrl,
    })
      .then((res) => {
        return res.json()
      })
      .then((json) => {
        showTips(json.code, json.msg)
      })
      .catch(() => {
        showWarmTips('更新失败')
      })
      .finally(() => {
        loading.close()
        setTimeout(function () {
          window.location.reload()
        }, 1000)
      })
  } else {
    showWarmTips('请正确输入url地址')
  }
}

const handleNewFrpcClient = () => {
  const body = JSON.stringify(newClientForm.value)
  put('客户端创建中...', '../api/client/create', body).finally(() => {
    newClientFormVisible.value = false
  })
}

// 自定义上传函数
const uploadToml = (options: any) => {
  const { file } = options
  const formData = new FormData()
  formData.append('file', file)
  const loading = showLoading('客户端创建中...')
  dialogFormVisible.value = false
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
      dialogFormVisible.value = false
      setTimeout(function () {
        window.location.reload()
      }, 1000)
    })
}

// 自定义上传函数
const customUpload = (options: any) => {
  const { file } = options
  const formData = new FormData()
  formData.append('file', file)
  const loading = showLoading('程序更新中...')
  dialogFormVisible.value = false
  // 使用 fetch 发送请求
  fetch('../api/upgrade', {
    method: 'POST',
    body: formData,
  })
    .then((response) => {
      return response.json()
    })
    .then((data:any) => {
      // 上传成功的回调
      console.log(data)
    })
    .catch((error:any) => {
      // 上传失败的回调
      console.log(error)
    })
    .finally(() => {
      loading.close()
      dialogFormVisible.value = false
      setTimeout(function () {
        window.location.reload()
      }, 1000)
    })
}

// 自定义上传函数
// const customUpload = (options: any) => {
//   const { file } = options
//   const formData = new FormData()
//   formData.append('file', file)
//   const loading = showLoading('程序更新中...')
//
//   dialogFormVisible.value = false
//   xhrPromise({
//     url: '../api/upgrade',
//     method: 'POST',
//     data: formData,
//     onUploadProgress: (progress: string) => {
//       console.log(`上传进度：${progress}`)
//       loading.setText(`程序更新中...${progress}%`)
//     },
//   })
//     .then((data: any) => {
//       console.log('请求成功', data)
//       // 上传成功的回调
//       const json = JSON.parse(data.data)
//       if (json.code !== 0) {
//         if (json.msg !== '') {
//           showErrorTips(json.msg)
//         }
//       } else {
//         if (json.msg !== '') {
//           showSucessTips(json.msg)
//         }
//       }
//     })
//     .catch((error) => {
//       console.error('请求失败', error)
//       // 上传失败的回调
//       showErrorTips('上传失败的回调')
//     })
//     .finally(() => {
//       loading.close()
//       dialogFormVisible.value = false
//       setTimeout(function () {
//         window.location.reload()
//       }, 1000)
//     })
// }

const restart = () => {
  showWarmDialog(
    `确定重启吗？`,
    () => {
      const loading = showLoading('重启中...')
      fetch('../api/restart', { credentials: 'include' })
        .then((res) => {
          return res.json()
        })
        .then((json) => {
          showTips(json.code, json.msg)
          location.reload()
        })
        .catch(() => {
          showErrorTips('重启失败')
        })
        .finally(() => {
          setTimeout(function () {
            loading.close()
            window.location.reload()
          }, 2000)
        })
    },
    () => {},
  )
}
fetchData()
</script>

<style>
body {
  margin: 0px;
  font-family:
    -apple-system,
    BlinkMacSystemFont,
    Helvetica Neue,
    sans-serif;
}

header {
  width: 100%;
  height: 60px;
}

.header-color {
  background: #58b7ff;
}

html.dark .header-color {
  background: #395c74;
}

.header-content {
  display: flex;
  align-items: center;
}

#content {
  margin-top: 20px;
  padding-right: 40px;
}

.brand {
  display: flex;
  justify-content: flex-start;
}

.brand a {
  color: #fff;
  background-color: transparent;
  margin-left: 20px;
  line-height: 25px;
  font-size: 25px;
  padding: 15px 15px;
  height: 30px;
  text-decoration: none;
}

.dark-reboot {
  padding-right: 10px;
}

.dark-switch {
  display: flex;
  justify-content: flex-end;
  flex-grow: 1;
  padding-right: 40px;
}
</style>
