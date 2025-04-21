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
        <div class="brand">
          <el-tooltip placement="right" effect="light">
            <template #content>
              <el-descriptions :column="1" border label-width="110px">
                <el-descriptions-item
                  label="frpc版本号："
                  label-align="left"
                  align="center"
                  width="150px"
                >
                  <el-tag size="small">{{ version.frpcVersion }}</el-tag>
                </el-descriptions-item>

                <el-descriptions-item
                  label="面板版本号："
                  label-align="left"
                  align="center"
                >
                  <el-tag size="small">{{ version.appVersion }}</el-tag>
                </el-descriptions-item>

                <el-descriptions-item
                  label="编译时间："
                  label-align="left"
                  align="center"
                  width="200px"
                >
                  {{ version.buildTime }}
                </el-descriptions-item>

                <el-descriptions-item
                  label="操作系统："
                  label-align="left"
                  align="center"
                >
                  {{ version.osType }}
                </el-descriptions-item>

                <el-descriptions-item
                  label="CPU架构："
                  label-align="left"
                  align="center"
                >
                  {{ version.arch }}
                </el-descriptions-item>
              </el-descriptions>
            </template>
            <a href="#">{{ title }}</a>
          </el-tooltip>
        </div>
        <div class="dark-switch">
          <div class="dark-reboot">
            <el-dropdown placement="bottom" split-button plain @click="restart">
              重启
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item @click="dialogFormVisible = true"
                    >升级服务</el-dropdown-item
                  >
                  <el-dropdown-item @click="checkVersion"
                    >版本检测</el-dropdown-item
                  >
                  <el-dropdown-item @click="showlog">查看日志</el-dropdown-item>
                  <el-dropdown-item @click="dialogClientsVisible = true"
                    >上传客户端</el-dropdown-item
                  >
                  <el-dropdown-item @click="handleClearData"
                    >清空数据</el-dropdown-item
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
            <el-menu-item index="/user">用户配置</el-menu-item>

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
              >上传文件升级</el-button
            >
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
      <i class="el-icon el-icon--upload"
        ><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024">
          <path
            fill="currentColor"
            d="M544 864V672h128L512 480 352 672h128v192H320v-1.6c-5.376.32-10.496 1.6-16 1.6A240 240 0 0 1 64 624c0-123.136 93.12-223.488 212.608-237.248A239.808 239.808 0 0 1 512 192a239.872 239.872 0 0 1 235.456 194.752c119.488 13.76 212.48 114.112 212.48 237.248a240 240 0 0 1-240 240c-5.376 0-10.56-1.28-16-1.6v1.6z"
          ></path></svg
      ></i>
      <div class="el-upload__text">拖拽到这里 <em>点击上传</em></div>
      <template #tip>
        <div class="el-upload__tip">
          请上传全平台架构的客户端程序放到dist文件夹并压缩，仅支持zip！
        </div>
      </template>
    </el-upload>
  </el-dialog>
</template>

<script lang="ts" setup>
import UpgradeDialog from './components/UpgradeDialog.vue'

import { onMounted, ref } from 'vue'
import { useDark, useToggle } from '@vueuse/core'
import {
  piecesUpload,
  showErrorTips,
  showLoading,
  showSucessTips,
  showTips,
  showWarmDialog,
  showWarmTips,
  xhrPromise,
} from './utils/utils.ts'
//https://element-plus-docs.bklab.cn/zh-CN/component/upload.html#upload-%E4%B8%8A%E4%BC%A0
const isDark = useDark()
const darkmodeSwitch = ref(isDark)
const toggleDark = useToggle(isDark)
const dialogFormVisible = ref(false)
const dialogClientsVisible = ref(false)
const version = ref({
  description: '', //应用描述
  frpcVersion: '', //frpc版本号
  buildTime: '', //编译时间
  appVersion: '', //app版本号
  appName: '', //app名称
  osType: '', //操作系统
  arch: '', //cpuType
  gitRevision: '',
  goVersion: '',
})

// const upgradeRef = ref(null)// 明确指定 ref 的类型
const upgradeRef = ref<InstanceType<typeof UpgradeDialog> | null>(null)

const showUpgradeDialog = () => {
  if (upgradeRef.value) {
    upgradeRef.value.openUpgradeDialog()
  }
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

const title = ref<string>('Frps')

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
  // fetch('../api/checkversion', { credentials: 'include' })
  //   .then((res) => {
  //     return res.json()
  //   })
  //   .then((json) => {
  //     if (json.code === 0) {
  //       showInfoTips(json.msg)
  //     } else if (json.code === 1) {
  //       showUpdateDialog(json.msg, json.data)
  //     }
  //   })

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
        version.value = json
        title.value = `Frps ${json.appVersion}`
        document.title = title.value
      }
    })
}

onMounted(() => {
  const mIndex = window.location.hash
  const result = mIndex.replace(/^#+/, '')
  console.log('index.menu.index', result)
  menuIndex.value = result
})
fetchVersionData()
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
