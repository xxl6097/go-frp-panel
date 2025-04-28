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
                    >升级服务
                  </el-dropdown-item>
                  <el-dropdown-item @click="checkVersion"
                    >版本检测</el-dropdown-item
                  >
                  <el-dropdown-item @click="showlog">查看日志</el-dropdown-item>
                  <el-dropdown-item @click="showVersion"
                    >查看版本</el-dropdown-item
                  >
                  <el-dropdown-item @click="uninstall"
                    >卸载自身
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
        <el-col id="side-nav" :xs="24" :md="4">
          <el-menu
            default-active="1"
            mode="vertical"
            theme="light"
            :router="'false'"
            @select="handleSelect"
          >
            <el-menu-item index="/">客户端信息</el-menu-item>
            <el-menu-item index="/configure">配置</el-menu-item>
            <el-menu-item index="/log">日志</el-menu-item>
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

  <UpgradeDialog ref="upgradeRef" />

  <!--  客户端程序升级-->
  <el-dialog v-model="dialogFormVisible" align-center width="500">
    <template #header><span>程序升级</span></template>
    <el-input
      v-model="form.binUrl"
      autocomplete="off"
      placeholder="请输入程序Url地址～"
    />

    <template #footer>
      <div class="dialog-footer">
        <el-upload class="upload-demo" :http-request="customUpload" :limit="1">
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

  <!-- 弹窗显示版本 -->
  <el-dialog v-model="versionDialogVisible" width="30%">
    <template #header><span>版本信息</span> </template>
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
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useDark, useToggle } from '@vueuse/core'
import {
  showLoading,
  showWarmDialog,
  showErrorTips,
  showTips,
  showWarmTips,
  showSucessTips,
  xhrPromise,
} from './utils/utils.ts'
import { ComponentSize } from 'element-plus'
import UpgradeDialog from './components/UpgradeDialog.vue'
const customColors = [
  { color: '#f56c6c', percentage: 20 },
  { color: '#e6a23c', percentage: 40 },
  { color: '#5cb87a', percentage: 60 },
  { color: '#1989fa', percentage: 80 },
  { color: '#6f7ad3', percentage: 100 },
]
const globalProgress = ref(0)
const size = ref<ComponentSize>('default')
const versionDialogVisible = ref(false)
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
const title = ref<string>('Frpc')
const isDark = useDark()
const darkmodeSwitch = ref(isDark)
const toggleDark = useToggle(isDark)
const dialogFormVisible = ref(false)
const form = ref({
  binUrl: '',
})

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

const uninstall = () => {
  // versionDialogVisible.value = true
  showWarmDialog(
    `确定要卸载程序吗，请认真思考！`,
    () => {
      const loading = showLoading('卸载中...')
      fetch('../api/uninstall', { credentials: 'include' })
        .then((res) => {
          return res.json()
        })
        .then((json) => {
          showTips(json.code, json.msg)
          location.reload()
        })
        .catch(() => {
          showErrorTips('卸载失败')
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

const fetchData = () => {
  fetch('../api/version', { credentials: 'include', method: 'GET' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      if (json) {
        version.value = json
        title.value = `Frpc ${json.data.appVersion}`
        document.title = title.value
      }
    })
}

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
    .catch(() => {
      showWarmTips('更新失败')
    })
    .finally(() => {
      setTimeout(function () {
        loading.close()
        window.location.reload()
      }, 4000)
    })
}

const upgrade = () => {
  if (form.value.binUrl.length > 0) {
    upgradeByUrl(form.value.binUrl)
  } else {
    showWarmTips('请正确输入url地址')
  }
}

// 自定义上传函数
// const customUpload = (options: any) => {
//   const { file } = options
//   const formData = new FormData()
//   formData.append('file', file)
//   const loading = showLoading('程序更新中...')
//   dialogFormVisible.value = false
//   // 使用 fetch 发送请求
//   fetch('../api/upgrade', {
//     method: 'POST',
//     body: formData,
//   })
//     .then((response) => {
//       return response.json()
//     })
//     .then((data:any) => {
//       // 上传成功的回调
//       console.log(data)
//     })
//     .catch((error:any) => {
//       // 上传失败的回调
//       console.log(error)
//     })
//     .finally(() => {
//       loading.close()
//       dialogFormVisible.value = false
//       setTimeout(function () {
//         window.location.reload()
//       }, 1000)
//     })
// }

// 自定义上传函数
const customUpload = (options: any) => {
  const { file } = options
  const formData = new FormData()
  formData.append('file', file)
  const loading = showLoading('程序更新中...')
  globalProgress.value = 0
  dialogFormVisible.value = false
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
    })
    .finally(() => {
      setTimeout(function () {
        loading.close()
        globalProgress.value = 0
        dialogFormVisible.value = false
        window.location.reload()
      }, 4000)
    })
}

const upgradeRef = ref<InstanceType<typeof UpgradeDialog> | null>(null)

const showUpgradeDialog = () => {
  if (upgradeRef.value) {
    upgradeRef.value.openUpgradeDialog()
  }
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

// const showUpdateDialog = (message: string, binurl: string) => {
//   showMessageDialog('升级提示', '升级', message).then(() => {
//     upgradeByUrl(binurl)
//   })
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
          }, 4000)
        })
    },
    () => {},
  )
}
fetchData()
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

.dark-reboot {
  padding-right: 10px;
}

.dark-switch {
  display: flex;
  justify-content: flex-end;
  flex-grow: 1;
  padding-right: 40px;
}

.global-progress-bar {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 9999;
  width: 100%;
}
</style>
