<template>
  <el-progress
    v-if="globalProgress > 0 && globalProgress < 100"
    :percentage="globalProgress"
    :stroke-width="2"
    :show-text="false"
    :color="customColors"
    class="global-progress-bar"
  />
  <div id="app">
    <header class="grid-content header-color">
      <div class="header-content">
        <div class="brand" @click="handleDevelopment">
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
                  <el-dropdown-item @click="checkVersion"
                    >版本检测
                  </el-dropdown-item>
                  <el-dropdown-item @click="showlog">查看日志</el-dropdown-item>
                  <el-dropdown-item @click="handleClearData"
                    >清空数据
                  </el-dropdown-item>
                  <el-dropdown-item @click="dialogClientsVisible = true"
                    >上传客户端
                  </el-dropdown-item>
                  <el-dropdown-item @click="frpsForm.isShow = true"
                    >创建服务端
                  </el-dropdown-item>
                  <el-dropdown-item @click="handleGithubKeySetting"
                    >设置github
                  </el-dropdown-item>
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
        <el-col id="side-nav" :md="4" :xs="24">
          <el-menu
            :default-active="menuIndex"
            mode="vertical"
            router="false"
            theme="light"
            @select="handleSelect"
          >
            <el-menu-item index="/">服务器信息</el-menu-item>

            <el-menu-item index="/config">服务器配置</el-menu-item>
            <el-menu-item index="/user">客户端配置</el-menu-item>

            <el-sub-menu index="/proxies">
              <template #title>
                <span>代理列表</span>
              </template>
              <el-menu-item index="/proxies/tcp">TCP</el-menu-item>
              <el-menu-item index="/proxies/udp">UDP</el-menu-item>
              <el-menu-item index="/proxies/http">HTTP</el-menu-item>
              <el-menu-item index="/proxies/https">HTTPS</el-menu-item>
              <el-menu-item index="/proxies/tcpmux">TCPMUX</el-menu-item>
              <el-menu-item index="/proxies/stcp">STCP</el-menu-item>
              <el-menu-item index="/proxies/sudp">SUDP</el-menu-item>
            </el-sub-menu>
            <el-menu-item index="/log">日志</el-menu-item>
            <el-menu-item index="/development" v-if="isDevelopment"
              >开发中模式
            </el-menu-item>
            <el-menu-item index="">帮助</el-menu-item>
          </el-menu>
        </el-col>

        <el-col :md="20" :xs="24">
          <div id="content">
            <router-view></router-view>
          </div>
        </el-col>
      </el-row>
    </section>
    <footer></footer>
  </div>

  <UpgradeDialog ref="upgradeRef" />

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
          :http-request="handleUploadToUpgrade"
          :limit="1"
        >
          <template #trigger>
            <el-button type="primary" :disabled="form.binUrl.length > 0"
              >上传文件升级
            </el-button>
          </template>
          <!-- 添加额外按钮 -->
          <el-button
            style="margin-left: 10px"
            type="danger"
            @click="handleUrlToUpgrade"
          >
            文件url升级
          </el-button>
        </el-upload>
      </div>
    </template>
  </el-dialog>

  <!--  上传客户端-->
  <el-dialog
    v-model="dialogClientsVisible"
    align-center
    title="客户端上传"
    width="500"
  >
    <template #title>
      <!-- 空标题或自定义隐藏内容 -->
    </template>
    <el-upload
      class="upload-demo"
      :http-request="doClientsUpload"
      drag
      :accept="'.zip'"
    >
      <i class="el-icon el-icon--upload">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024">
          <path
            fill="currentColor"
            d="M544 864V672h128L512 480 352 672h128v192H320v-1.6c-5.376.32-10.496 1.6-16 1.6A240 240 0 0 1 64 624c0-123.136 93.12-223.488 212.608-237.248A239.808 239.808 0 0 1 512 192a239.872 239.872 0 0 1 235.456 194.752c119.488 13.76 212.48 114.112 212.48 237.248a240 240 0 0 1-240 240c-5.376 0-10.56-1.28-16-1.6v1.6z"
          ></path>
        </svg>
      </i>
      <div class="el-upload__text">拖拽到这里 <em>点击上传</em></div>
      <template #tip>
        <div class="el-upload__tip">
          请上传全平台架构的客户端程序放到dist文件夹并压缩，仅支持zip！
        </div>
      </template>
    </el-upload>
  </el-dialog>

  <!-- 生成服务端 -->
  <el-dialog v-model="frpsForm.isShow" title="生成Frps服务端" width="500px">
    <el-form label-width="130px">
      <el-form-item label="Frps绑定端口：">
        <el-input v-model="frpsForm.bindPort" placeholder="请输入bindport" />
      </el-form-item>
      <el-form-item label="Admin管理端口：">
        <el-input
          v-model="frpsForm.adminPort"
          placeholder="请输入Admin管理端口"
        />
      </el-form-item>
      <el-form-item label="Admin账户：">
        <el-input v-model="frpsForm.user" placeholder="请输入Admin账户" />
      </el-form-item>
      <el-form-item label="Admin密码：">
        <el-input v-model="frpsForm.pass" placeholder="请输入Admin密码" />
      </el-form-item>
      <el-form-item
        label="操作系统/架构"
        v-if="frpsForm.options && frpsForm.options.length > 0"
      >
        <el-cascader
          :options="frpsForm.options"
          clearable
          v-model="frpsForm.ops"
          placeholder="请选择"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        type="primary"
        @click="fetchFrpsGen()"
        :loading="frpsForm.isLoading"
        >生成
      </el-button>
    </template>
  </el-dialog>

  <!-- 填写云Api信息设置 -->
  <el-dialog
    v-model="githubApiForm.isShow"
    title="github key设置"
    width="500px"
  >
    <el-form label-width="130px">
      <el-form-item label="client id：">
        <el-input
          v-model="githubApiForm.clientId"
          placeholder="请输入github client id"
        />
      </el-form-item>
      <el-form-item label="client secret：">
        <el-input
          v-model="githubApiForm.clientSecret"
          placeholder="请输入请输入github client secret"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button type="primary" @click="handleGithubKeySetting">确定</el-button>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import UpgradeDialog from './components/UpgradeDialog.vue'

import { onMounted, ref, provide, onUnmounted } from 'vue'
import { useDark, useToggle } from '@vueuse/core'
import {
  downloadByPost,
  piecesUpload,
  showErrorTips,
  showLoading,
  showSucessTips,
  showTips,
  showWarmDialog,
  showWarmTips,
  xhrPromise,
} from './utils/utils.ts'
import { Version } from './utils/type.ts'
import { ElButton } from 'element-plus'
//https://element-plus-docs.bklab.cn/zh-CN/component/upload.html#upload-%E4%B8%8A%E4%BC%A0
const isDark = useDark()
const darkmodeSwitch = ref(isDark)
const toggleDark = useToggle(isDark)
const dialogFormVisible = ref(false)
const dialogClientsVisible = ref(false)
const isDevelopment = ref(false)
// const version = ref({
//   description: '', //应用描述
//   frpcVersion: '', //frpc版本号
//   buildTime: '', //编译时间
//   appVersion: '', //app版本号
//   appName: '', //app名称
//   osType: '', //操作系统
//   arch: '', //cpuType
//   gitRevision: '',
//   goVersion: '',
// })

const clickCount = ref(0)
const frpsForm = ref({
  bindPort: 6000,
  adminAddr: '0.0.0.0',
  adminPort: 6500,
  user: 'admin',
  pass: '',
  options: [],
  ops: {},
  isShow: false,
  isLoading: false,
})

const githubApiForm = ref({
  clientId: '',
  clientSecret: '',
  isShow: false,
})
const version = ref<Version>()
provide('version', version)

// const upgradeRef = ref(null)// 明确指定 ref 的类型
const upgradeRef = ref<InstanceType<typeof UpgradeDialog> | null>(null)

// 配置上传云端
const handleGithubKeySetting = () => {
  console.log('handleGithubKeySetting:', githubApiForm.value)
  if (githubApiForm.value.isShow) {
    fetch('../api/github/key', {
      credentials: 'include',
      method: 'post',
      body: JSON.stringify(githubApiForm.value),
    })
      .then((res) => res.json())
      .then((json) => {
        console.log('配置备份', json)
        showTips(json.code, json.msg)
        if (json.code === 0) {
          githubApiForm.value.isShow = false
        }
      })
      .finally(() => {
        localStorage.setItem('githubKey', JSON.stringify(githubApiForm.value))
      })
  } else {
    fetch('../api/github/key', {
      credentials: 'include',
      method: 'get',
    })
      .then((res) => res.json())
      .then((json) => {
        console.log('github key', json)
        if (json.code === 100) {
          githubApiForm.value.isShow = true
          if (json.data) {
            githubApiForm.value.clientId = json.data.clientId
            githubApiForm.value.clientSecret = json.data.clientSecret
          }
        }
        showTips(json.code, json.msg)
      })
      .finally(() => {})
  }
}

const showUpgradeDialog = () => {
  if (upgradeRef.value) {
    upgradeRef.value.openUpgradeDialog()
  }
}

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
// 获取平台数据
const fetchOptions = () => {
  fetch('../api/frps/get', {
    credentials: 'include',
    method: 'GET',
  })
    .then((res) => {
      return res.json()
    })
    .then((data) => {
      console.log('1-frpsForm.options', data)
      if (data) {
        frpsForm.value.options = data.data
      } else {
        frpsForm.value.options = []
      }
    })
}
// 创建frps安装程序
const fetchFrpsGen = () => {
  frpsForm.value.isLoading = true
  console.log('创建frps安装程序', frpsForm.value)
  downloadByPost(
    'frps生成中',
    '../api/frps/gen',
    JSON.stringify(frpsForm.value),
  ).finally(() => {
    frpsForm.value.isLoading = false
  })
}

const customColors = [
  { color: '#f56c6c', percentage: 20 },
  { color: '#e6a23c', percentage: 40 },
  { color: '#5cb87a', percentage: 60 },
  { color: '#1989fa', percentage: 80 },
  { color: '#6f7ad3', percentage: 100 },
]
const globalProgress = ref(0)
const form = ref({
  binUrl: '',
})
const menuIndex = ref('/')

const title = ref<string>('rps')

const doClientsUpload = async (options: any) => {
  const { file } = options
  globalProgress.value = 0
  dialogFormVisible.value = false
  const loading = showLoading('客户端上传中...')

  await piecesUpload(
    '../api/client/upload',
    'POST',
    file,
    (progress: any) => {
      console.log(`上传进度：${progress}`)
      loading.setText(`客户端上传中：${progress}%`)
      globalProgress.value = parseInt(progress)
    },
    () => {
      loading.close()
      dialogClientsVisible.value = false
      console.log(`上传成功！`)
      showSucessTips('上传成功')
      globalProgress.value = 0
    },
  )
}

// 文件上传更新
const handleUploadToUpgrade = (options: any) => {
  const { file } = options
  const formData = new FormData()
  formData.append('file', file)
  const loading = showLoading('程序更新中...')

  globalProgress.value = 0
  dialogFormVisible.value = false
  const ok = ref<boolean>(false)
  xhrPromise({
    url: '../api/upgrade',
    method: 'POST',
    data: formData,
    onUploadProgress: (progress: string) => {
      console.log(`上传进度：${progress}`)
      loading.setText(`程序更新中...${progress}%`)
      globalProgress.value = parseInt(progress)
    },
  })
    .then((data: any) => {
      console.log('请求成功', data)
      // 上传成功的回调
      const json = JSON.parse(data.data)
      if (json.code !== 0) {
        if (json.msg !== '') {
          showErrorTips(json.msg)
        }
      } else {
        if (json.msg !== '') {
          showSucessTips(json.msg)
        }
      }
    })
    .catch((error) => {
      console.error('请求失败', error)
      // 上传失败的回调
      //showErrorTips('上传失败的回调')
      ok.value = false
    })
    .finally(() => {
      setTimeout(function () {
        loading.close()
        globalProgress.value = 0
        dialogFormVisible.value = false
        if (ok.value) {
          window.location.reload()
        }
      }, 4000)
    })
}

const handleSelect = (key: string) => {
  if (key == '') {
    window.open('https://github.com/xxl6097/go-frp-panel')
  }
  console.log('menu.key', key)
  menuIndex.value = key
}

const checkVersion = () => {
  showUpgradeDialog()
}

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
          }, 4000)
        })
    },
    () => {},
  )
}

const handleClearData = () => {
  showWarmDialog(
    `确定清空应用数据吗？`,
    () => {
      fetch('../api/clear', { credentials: 'include', method: 'DELETE' })
        .then((res) => {
          return res.json()
        })
        .then((json) => {
          showTips(json.code, json.msg)
        })
        .catch(() => {
          showErrorTips('清空失败')
        })
    },
    () => {},
  )
}

const showlog = () => {
  const host = window.origin
  window.open(`${host}/log/`)
}
// const upgrade = () => {
//   if (form.value.binUrl.length > 0) {
//     const loading = showLoading('程序升级中...')
//     dialogFormVisible.value = false
//     xhrPromise({
//       url: '../api/upgrade',
//       method: 'PUT',
//       data: form.value.binUrl,
//       onUploadProgress: (progress: string) => {
//         console.log(`上传进度：${progress}`)
//         loading.setText(`程序更新中...${progress}%`)
//       },
//     })
//       .then((data: any) => {
//         console.log('请求成功', data)
//         // 上传成功的回调
//         const json = JSON.parse(data.data)
//         if (json.code !== 0) {
//           if (json.msg !== '') {
//             showErrorTips(json.msg)
//           }
//         } else {
//           if (json.msg !== '') {
//             showSucessTips(json.msg)
//           }
//         }
//       })
//       .catch((error) => {
//         console.error('请求失败', error)
//         // 上传失败的回调
//         showErrorTips('上传失败的回调')
//       })
//       .finally(() => {
//         loading.close()
//         dialogFormVisible.value = false
//         setTimeout(function () {
//           window.location.reload()
//         }, 1000)
//       })
//   } else {
//     showWarmTips('请正确输入url地址')
//   }
// }

const upgradeByUrl = (binUrl: string) => {
  console.log('upgradeByUrl', binUrl)
  const loading = showLoading('程序升级中...')
  dialogFormVisible.value = false
  fetch('../api/upgrade', {
    credentials: 'include',
    method: 'PUT',
    body: binUrl,
  })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      showTips(json.code, json.msg)
    })
    .catch((error) => {
      console.log('更新失败', error)
      //showWarmTips('更新失败' + JSON.stringify(error))
    })
    .finally(() => {
      setTimeout(function () {
        loading.close()
        window.location.reload()
      }, 4000)
    })
}

const handleUrlToUpgrade = () => {
  if (form.value.binUrl.length > 0) {
    upgradeByUrl(form.value.binUrl)
  } else {
    showWarmTips('请正确输入url地址')
  }
}

// const uploadFile = (file: any) => {
//   const loading = showLoading('客户端上传中...')
//   return new Promise((resolve, reject) => {
//     // 创建一个新的 XMLHttpRequest 对象
//     const xhr = new XMLHttpRequest()
//     // 打开一个 POST 请求，这里的 URL 可以根据实际情况修改
//     xhr.open('POST', '../api/client/upload', true)
//
//     // 监听上传进度事件
//     xhr.upload.addEventListener('progress', (event) => {
//       if (event.lengthComputable) {
//         const percentComplete = (event.loaded / event.total) * 100
//         console.log('--->', percentComplete + '%')
//         uploadPercent.value = percentComplete.toFixed(2)
//         loading.setText(`客户端上传中 ${uploadPercent.value}%`)
//       }
//     })
//
//     // 监听请求完成事件
//     xhr.addEventListener('load', () => {
//       if (xhr.status >= 200 && xhr.status < 300) {
//         resolve(xhr.response)
//       } else {
//         reject(new Error(`Upload failed with status ${xhr.status}`))
//       }
//     })
//
//     // 监听请求出错事件
//     xhr.addEventListener('error', () => {
//       reject(new Error('Network error occurred during upload'))
//     })
//
//     // 创建一个 FormData 对象
//     const formData = new FormData()
//     formData.append('file', file)
//
//     // 发送请求
//     xhr.send(formData)
//   })
// }

// const doClientsUpload1 = (options: any) => {
//   const { file } = options
//   const formData = new FormData()
//   formData.append('file', file)
//   const loading = showLoading('客户端上传中...')
//   dialogFormVisible.value = false
//   // 使用 fetch 发送请求
//   fetch('../api/client/upload', {
//     method: 'POST',
//     body: formData,
//   })
//     .then((response) => {
//       return response.json()
//     })
//     .then((data) => {
//       // 上传成功的回调
//       showSucessTips(JSON.stringify(data))
//     })
//     .catch((error) => {
//       // 上传失败的回调
//       console.log('errr', error)
//       showSucessTips(JSON.stringify(error))
//     })
//     .finally(() => {
//       loading.close()
//       dialogClientsVisible.value = false
//     })
// }

const fetchVersionData = () => {
  fetch('../api/version', { credentials: 'include', method: 'GET' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      if (json) {
        //version.value = json
        const vv = json.data as Version // 类型断言
        version.value = vv
        console.log('vv', vv)
        console.log('version', version.value)
        title.value = `frps ${vv.appVersion}`
        document.title = `frps ${vv.hostName}`
      }
    })
}
// 组件卸载时清除定时器（防止内存泄漏）
onUnmounted(() => {
  if (timer) clearTimeout(timer)
})
onMounted(() => {
  const mIndex = window.location.hash
  const result = mIndex.replace(/^#+/, '')
  console.log('index.menu.index', result)
  menuIndex.value = result

  const jsonString = localStorage.getItem('githubKey')
  if (jsonString) {
    const obj = JSON.parse(jsonString)
    if (obj) {
      githubApiForm.value = obj
    }
    githubApiForm.value.isShow = false
  }
})
fetchVersionData()
fetchOptions()
// checkVersion()
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

.dark-switch {
  display: flex;
  justify-content: flex-end;
  flex-grow: 1;
  padding-right: 40px;
}

.dark-reboot {
  padding-right: 10px;
}

.global-progress-bar {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 9999;
  width: 100%;
}
</style>
