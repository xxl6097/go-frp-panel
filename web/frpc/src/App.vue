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
    <!-- 顶部导航（frp v0.70 风格） -->
    <header class="header">
      <div class="header-content">
        <div class="brand-section">
          <button
            v-if="isMobile"
            class="hamburger-btn"
            @click="toggleSidebar"
            aria-label="Toggle menu"
          >
            <span class="hamburger-icon">&#9776;</span>
          </button>
          <div class="logo-wrapper" @click="handleDevelopment">
            <LogoIcon class="logo-icon" />
          </div>
          <span class="divider">/</span>
          <span class="brand-name">{{ title }}</span>
          <span class="badge client-badge">Client</span>
        </div>

        <div class="header-controls">
          <el-dropdown trigger="click">
            <button class="icon-btn" aria-label="Settings">
              <el-icon><Setting /></el-icon>
            </button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="restart">重启服务</el-dropdown-item>
                <el-dropdown-item @click="dialogFormVisible = true"
                  >升级服务</el-dropdown-item
                >
                <el-dropdown-item @click="checkVersion"
                  >版本检测</el-dropdown-item
                >
                <el-dropdown-item @click="showlog">查看日志</el-dropdown-item>
                <el-dropdown-item @click="handleClearData"
                  >清空数据</el-dropdown-item
                >
                <el-dropdown-item @click="showVersion"
                  >查看版本</el-dropdown-item
                >
                <el-dropdown-item @click="uninstall">卸载自身</el-dropdown-item>
                <el-dropdown-item @click="githubProxyForm.isShow = true"
                  >设置proxy</el-dropdown-item
                >
              </el-dropdown-menu>
            </template>
          </el-dropdown>

          <a
            class="github-link"
            href="https://github.com/xxl6097/go-frp-panel"
            target="_blank"
            aria-label="GitHub"
          >
            <GitHubIcon class="github-icon" />
          </a>
          <el-switch
            v-model="isDark"
            inline-prompt
            :active-icon="Moon"
            :inactive-icon="Sunny"
            class="theme-switch"
            @change="toggleDark"
          />
        </div>
      </div>
    </header>

    <div class="layout">
      <div
        v-if="isMobile && sidebarOpen"
        class="sidebar-overlay"
        @click="closeSidebar"
      />

      <aside class="sidebar" :class="{ 'mobile-open': isMobile && sidebarOpen }">
        <nav class="sidebar-nav">
          <router-link
            to="/"
            class="sidebar-link"
            :class="{ active: route.path === '/' }"
            @click="closeSidebar"
            >客户端信息</router-link
          >
          <router-link
            to="/configure"
            class="sidebar-link"
            :class="{ active: route.path === '/configure' }"
            @click="closeSidebar"
            >配置</router-link
          >
          <a class="sidebar-link" href="javascript:void(0)" @click="showlog"
            >日志</a
          >
          <router-link
            v-if="isDevelopment"
            to="/development"
            class="sidebar-link"
            :class="{ active: route.path === '/development' }"
            @click="closeSidebar"
            >开发中模式</router-link
          >
        </nav>
      </aside>

      <main id="content">
        <router-view></router-view>
      </main>
    </div>
  </div>

  <UpgradeDialog ref="upgradeRef" />

  <!-- 客户端程序升级 -->
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
          <el-button style="margin-left: 10px" type="danger" @click="upgrade">
            文件url升级
          </el-button>
        </el-upload>
      </div>
    </template>
  </el-dialog>

  <!-- 版本信息 -->
  <el-dialog v-model="versionDialogVisible" width="30%">
    <template #header><span>版本信息</span></template>
    <el-descriptions :column="1" :size="size" border>
      <el-descriptions-item width="100">
        <template #label><div class="cell-item">软件名称</div></template>
        {{ version?.appName }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label><div class="cell-item">软件版本</div></template>
        {{ version?.appVersion }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label><div class="cell-item">编译时间</div></template>
        {{ version?.buildTime }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label><div class="cell-item">frpc版本号</div></template>
        {{ version?.frpcVersion }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label><div class="cell-item">git版本</div></template>
        {{ version?.gitRevision }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label><div class="cell-item">go编译版本</div></template>
        {{ version?.goVersion }}
      </el-descriptions-item>
    </el-descriptions>
  </el-dialog>

  <!-- github proxy -->
  <el-dialog
    v-model="githubProxyForm.isShow"
    title="设置github api代理"
    width="500px"
  >
    <el-input
      v-model="githubProxyForm.proxyUrl"
      placeholder="请输入代理，为空则清除"
    />
    <template #footer>
      <el-button type="primary" @click="handleNewGithubProxy">确定</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, watch, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import { useDark, useToggle } from '@vueuse/core'
import { Moon, Sunny, Setting } from '@element-plus/icons-vue'
import GitHubIcon from './assets/icons/github.svg?component'
import LogoIcon from './assets/icons/logo.svg?component'
import { useResponsive } from './composables/useResponsive'
import {
  showLoading,
  showWarmDialog,
  showErrorTips,
  showTips,
  showWarmTips,
  showSucessTips,
  xhrPromise,
  Version,
  showInfoTips,
} from './utils/utils.ts'
import { ComponentSize } from 'element-plus'
import UpgradeDialog from './components/UpgradeDialog.vue'

const route = useRoute()
const { isMobile } = useResponsive()
const sidebarOpen = ref(false)
const toggleSidebar = () => {
  sidebarOpen.value = !sidebarOpen.value
}
const closeSidebar = () => {
  sidebarOpen.value = false
}
watch(
  () => route.path,
  () => {
    if (isMobile.value) closeSidebar()
  },
)

const githubProxyForm = ref({
  proxyUrl: '',
  isShow: false,
})
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
const version = ref<Version>()
const title = ref<string>('frpc')
const isDark = useDark()
const toggleDark = useToggle(isDark)
const dialogFormVisible = ref(false)
const form = ref({ binUrl: '' })
const isDevelopment = ref(false)
const clickCount = ref(0)
let timer: number | null = null

const handleDevelopment = () => {
  if (clickCount.value === 0) {
    timer = window.setTimeout(() => {
      clickCount.value = 0
      timer = null
    }, 1000)
  }
  clickCount.value++
  if (clickCount.value === 5) {
    showWarmTips('进入开发者模式')
    isDevelopment.value = true
    clickCount.value = 0
    if (timer) {
      clearTimeout(timer)
      timer = null
    }
  }
}

const showlog = () => {
  window.open(`${window.origin}/log/`)
}

const handleNewGithubProxy = () => {
  const data = { proxyUrl: githubProxyForm.value.proxyUrl }
  fetch(`../api/proxy/github/api`, {
    method: 'PUT',
    body: JSON.stringify(data),
  })
    .then((response) => response.json())
    .then((json) => showInfoTips(json.msg))
    .finally(() => {
      githubProxyForm.value.isShow = false
    })
}

const showVersion = () => {
  fetch('../api/version', { credentials: 'include', method: 'GET' })
    .then((res) => res.json())
    .then((json) => {
      if (json.code === 0 && json.data) {
        version.value = json
        versionDialogVisible.value = true
      }
      showTips(json.code, json.msg)
    })
    .catch(() => showErrorTips('失败'))
}

const uninstall = () => {
  showWarmDialog(
    `确定要卸载程序吗，请认真思考！`,
    () => {
      const loading = showLoading('卸载中...')
      fetch('../api/uninstall', { credentials: 'include' })
        .then((res) => res.json())
        .then((json) => {
          showTips(json.code, json.msg)
          location.reload()
        })
        .catch(() => showErrorTips('卸载失败'))
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
        .then((res) => res.json())
        .then((json) => showTips(json.code, json.msg))
        .catch(() => showErrorTips('清空失败'))
    },
    () => {},
  )
}

const fetchData = () => {
  fetch('../api/version', { credentials: 'include', method: 'GET' })
    .then((res) => res.json())
    .then((json) => {
      if (json) {
        const vv = json.data as Version
        version.value = vv
        title.value = `frpc ${vv.appVersion}`
        document.title = `frpc ${vv.hostName}`
      }
    })
}

const upgradeByUrl = (binUrl: string) => {
  const loading = showLoading('程序升级中...')
  dialogFormVisible.value = false
  fetch('../api/upgrade', {
    credentials: 'include',
    method: 'PUT',
    body: binUrl,
  })
    .then((res) => res.json())
    .then((json) => showTips(json.code, json.msg))
    .catch(() => showWarmTips('更新失败'))
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
      loading.setText(`程序更新中...${progress}%`)
      globalProgress.value = parseInt(progress)
    },
  })
    .then((data: any) => {
      const json = JSON.parse(data.data)
      if (json.code !== 0) {
        if (json.msg !== '') showErrorTips(json.msg)
      } else {
        if (json.msg !== '') showSucessTips(json.msg)
      }
    })
    .catch((error) => console.error('请求失败', error))
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
const checkVersion = () => {
  upgradeRef.value?.openUpgradeDialog()
}

const restart = () => {
  showWarmDialog(
    `确定重启吗？`,
    () => {
      const loading = showLoading('重启中...')
      fetch('../api/restart', { credentials: 'include' })
        .then((res) => res.json())
        .then((json) => {
          showTips(json.code, json.msg)
          location.reload()
        })
        .catch(() => showErrorTips('重启失败'))
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

onUnmounted(() => {
  if (timer) clearTimeout(timer)
})
fetchData()
</script>
<style>
:root {
  --header-height: 50px;
  --sidebar-width: 200px;
  --header-bg: var(--color-bg-primary, #ffffff);
  --header-border: var(--color-border-light, #e4e7ed);
  --sidebar-bg: var(--color-bg-primary, #ffffff);
  --content-bg: var(--color-bg-secondary, #f9f9f9);
  --brand-text: var(--color-text-primary, #303133);
  --muted-text: var(--color-text-muted, #909399);
  --hover-bg: var(--color-bg-hover, #efefef);
}

body {
  margin: 0;
  font-family:
    ui-sans-serif, -apple-system, system-ui, Segoe UI, Helvetica, Arial,
    sans-serif;
}

*,
:after,
:before {
  box-sizing: border-box;
  -webkit-tap-highlight-color: transparent;
}

html,
body {
  height: 100%;
  overflow: hidden;
}

#app {
  height: 100vh;
  height: 100dvh;
  display: flex;
  flex-direction: column;
  background-color: var(--content-bg);
}

.header {
  flex-shrink: 0;
  background: var(--header-bg);
  border-bottom: 1px solid var(--header-border);
  height: var(--header-height);
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  padding: 0 20px;
}

.brand-section {
  display: flex;
  align-items: center;
  gap: 12px;
}

.logo-wrapper {
  display: flex;
  align-items: center;
  cursor: pointer;
}

.logo-icon {
  width: 28px;
  height: 28px;
}

.divider {
  color: var(--header-border);
  font-size: 22px;
  font-weight: 200;
}

.brand-name {
  font-weight: 600;
  font-size: 18px;
  color: var(--brand-text);
  letter-spacing: -0.5px;
}

.badge {
  font-size: 11px;
  font-weight: 500;
  color: var(--muted-text);
  background: var(--hover-bg);
  padding: 2px 8px;
  border-radius: 4px;
}

.badge.client-badge {
  background: linear-gradient(135deg, #8b5cf6 0%, #ec4899 100%);
  color: white;
  border: none;
}

html.dark .badge.client-badge {
  background: linear-gradient(135deg, #a78bfa 0%, #f472b6 100%);
}

.header-controls {
  display: flex;
  align-items: center;
  gap: 12px;
}

.icon-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border: none;
  background: transparent;
  border-radius: 6px;
  color: var(--color-text-secondary, #606266);
  cursor: pointer;
  padding: 0;
  font-size: 16px;
  transition: all 0.15s ease;
}

.icon-btn:hover {
  background: var(--hover-bg);
  color: var(--brand-text);
}

.github-link {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border-radius: 6px;
  color: var(--color-text-secondary, #606266);
  transition: all 0.15s ease;
  text-decoration: none;
}

.github-link:hover {
  background: var(--hover-bg);
  color: var(--brand-text);
}

.github-icon {
  width: 18px;
  height: 18px;
}

.theme-switch {
  --el-switch-on-color: #2c2c3a;
  --el-switch-off-color: #f2f2f2;
  --el-switch-border-color: var(--header-border);
}

html.dark .theme-switch {
  --el-switch-off-color: #333;
}

.layout {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.sidebar {
  width: var(--sidebar-width);
  flex-shrink: 0;
  background: var(--sidebar-bg);
  border-right: 1px solid var(--header-border);
  padding: 16px 12px;
  overflow-y: auto;
}

.sidebar-nav {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.sidebar-link {
  display: block;
  text-decoration: none;
  font-size: 15px;
  color: var(--color-text-secondary, #606266);
  padding: 10px 12px;
  border-radius: 6px;
  transition: all 0.15s ease;
  cursor: pointer;
}

.sidebar-link:hover {
  color: var(--brand-text);
  background: var(--hover-bg);
}

.sidebar-link.active {
  color: var(--brand-text);
  background: var(--hover-bg);
  font-weight: 500;
}

.hamburger-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  border: none;
  border-radius: 6px;
  background: transparent;
  cursor: pointer;
  padding: 0;
}

.hamburger-btn:hover {
  background: var(--hover-bg);
}

.hamburger-icon {
  font-size: 20px;
  line-height: 1;
  color: var(--brand-text);
}

.sidebar-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  z-index: 99;
}

#content {
  flex: 1;
  min-width: 0;
  overflow-y: auto;
  padding: 24px 32px;
}

.global-progress-bar {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 9999;
  width: 100%;
}

@media (max-width: 767px) {
  .header-content {
    padding: 0 16px;
  }

  .sidebar {
    position: fixed;
    top: var(--header-height);
    left: 0;
    bottom: 0;
    z-index: 100;
    transform: translateX(-100%);
    transition: transform 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .sidebar.mobile-open {
    transform: translateX(0);
  }

  #content {
    width: 100%;
    padding: 16px;
  }
}
</style>
