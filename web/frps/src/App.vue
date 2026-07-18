<template>
  <div id="app">
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
          <span class="badge server-badge">Server</span>
        </div>

        <div class="header-controls">
          <!-- 设置菜单：保留旧面板的全部管理操作 -->
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
                <el-dropdown-item @click="dialogClientsVisible = true"
                  >上传客户端</el-dropdown-item
                >
                <el-dropdown-item @click="frpsForm.isShow = true"
                  >创建服务端</el-dropdown-item
                >
                <el-dropdown-item @click="handleGithubKeySetting"
                  >设置github</el-dropdown-item
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
          />
        </div>
      </div>
    </header>

    <div class="layout">
      <!-- Mobile overlay -->
      <div
        v-if="isMobile && sidebarOpen"
        class="sidebar-overlay"
        @click="closeSidebar"
      />

      <aside
        class="sidebar"
        :class="{ 'mobile-open': isMobile && sidebarOpen }"
      >
        <nav class="sidebar-nav">
          <router-link
            to="/"
            class="sidebar-link"
            :class="{ active: route.path === '/' }"
            @click="closeSidebar"
          >
            概览
          </router-link>
          <router-link
            to="/clients"
            class="sidebar-link"
            :class="{ active: route.path.startsWith('/clients') }"
            @click="closeSidebar"
          >
            客户端
          </router-link>
          <router-link
            to="/proxies"
            class="sidebar-link"
            :class="{
              active:
                route.path.startsWith('/proxies') ||
                route.path.startsWith('/proxy'),
            }"
            @click="closeSidebar"
          >
            代理
          </router-link>

          <div class="sidebar-divider"></div>

          <router-link
            to="/config"
            class="sidebar-link"
            :class="{ active: route.path === '/config' }"
            @click="closeSidebar"
          >
            服务器配置
          </router-link>
          <router-link
            to="/user"
            class="sidebar-link"
            :class="{ active: route.path === '/user' }"
            @click="closeSidebar"
          >
            客户端配置
          </router-link>

          <div class="sidebar-divider"></div>

          <a class="sidebar-link" href="javascript:void(0)" @click="showlog">
            日志
          </a>
          <router-link
            v-if="isDevelopment"
            to="/development"
            class="sidebar-link"
            :class="{ active: route.path === '/development' }"
            @click="closeSidebar"
          >
            开发中模式
          </router-link>
        </nav>
      </aside>

      <main id="content">
        <router-view></router-view>
      </main>
    </div>
  </div>

  <el-progress
    v-if="globalProgress > 0 && globalProgress < 100"
    :percentage="globalProgress"
    :stroke-width="2"
    :show-text="false"
    :color="customColors"
    class="global-progress-bar"
  />

  <UpgradeDialog ref="upgradeRef" />

  <!-- 程序升级 -->
  <el-dialog v-model="dialogFormVisible" align-center title="程序升级" width="500">
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
          <el-button
            style="margin-left: 10px"
            type="danger"
            @click="handleUrlToUpgrade"
            >文件url升级</el-button
          >
        </el-upload>
      </div>
    </template>
  </el-dialog>

  <!-- 上传客户端 -->
  <el-dialog
    v-model="dialogClientsVisible"
    align-center
    title="客户端上传"
    width="500"
  >
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
        <el-input v-model="frpsForm.adminPort" placeholder="请输入Admin管理端口" />
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
      <el-button type="primary" @click="fetchFrpsGen()" :loading="frpsForm.isLoading"
        >生成</el-button
      >
    </template>
  </el-dialog>

  <!-- github key 设置 -->
  <el-dialog v-model="githubApiForm.isShow" title="github key设置" width="500px">
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
      <el-button type="warning" @click="handleGithubKeySetting">清空</el-button>
      <el-button type="primary" @click="handleGithubKeySetting">确定</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, watch, provide, onMounted, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import { useDark } from '@vueuse/core'
import { Moon, Sunny, Setting } from '@element-plus/icons-vue'
import GitHubIcon from './assets/icons/github.svg?component'
import LogoIcon from './assets/icons/logo.svg?component'
import { useResponsive } from './composables/useResponsive'
import UpgradeDialog from './components/UpgradeDialog.vue'
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

const route = useRoute()
const isDark = useDark()
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

/* ---------- 面板标题 / 版本 ---------- */
const title = ref<string>('frps')
const version = ref<Version>()
provide('version', version)

/* ---------- 开发者模式（连续点击 logo 5 次） ---------- */
const isDevelopment = ref(false)
const clickCount = ref(0)
let devTimer: number | null = null
const handleDevelopment = () => {
  if (clickCount.value === 0) {
    devTimer = window.setTimeout(() => {
      clickCount.value = 0
      devTimer = null
    }, 1000)
  }
  clickCount.value++
  if (clickCount.value === 5) {
    showWarmTips('进入开发者模式')
    isDevelopment.value = true
    clickCount.value = 0
    if (devTimer) {
      clearTimeout(devTimer)
      devTimer = null
    }
  }
}

/* ---------- 全局上传进度条 ---------- */
const globalProgress = ref(0)
const customColors = [
  { color: '#f56c6c', percentage: 20 },
  { color: '#e6a23c', percentage: 40 },
  { color: '#5cb87a', percentage: 60 },
  { color: '#1989fa', percentage: 80 },
  { color: '#6f7ad3', percentage: 100 },
]

/* ---------- 升级对话框 ---------- */
const upgradeRef = ref<InstanceType<typeof UpgradeDialog> | null>(null)
const dialogFormVisible = ref(false)
const form = ref({ binUrl: '' })
const checkVersion = () => {
  upgradeRef.value?.openUpgradeDialog()
}

/* ---------- 上传客户端 ---------- */
const dialogClientsVisible = ref(false)

/* ---------- 生成服务端 ---------- */
const frpsForm = ref({
  bindPort: 6000,
  adminAddr: '0.0.0.0',
  adminPort: 6500,
  user: 'admin',
  pass: '',
  options: [] as any[],
  ops: {},
  isShow: false,
  isLoading: false,
})

/* ---------- github key ---------- */
const githubApiForm = ref({ clientId: '', clientSecret: '', isShow: false })

const handleGithubKeySetting = () => {
  if (githubApiForm.value.isShow) {
    fetch('../api/github/key', {
      credentials: 'include',
      method: 'post',
      body: JSON.stringify(githubApiForm.value),
    })
      .then((res) => res.json())
      .then((json) => {
        showTips(json.code, json.msg)
        if (json.code === 0) githubApiForm.value.isShow = false
      })
      .finally(() => {
        localStorage.setItem('githubKey', JSON.stringify(githubApiForm.value))
      })
  } else {
    fetch('../api/github/key', { credentials: 'include', method: 'get' })
      .then((res) => res.json())
      .then((json) => {
        if (json.code === 100) {
          githubApiForm.value.isShow = true
          if (json.data) {
            githubApiForm.value.clientId = json.data.clientId
            githubApiForm.value.clientSecret = json.data.clientSecret
          }
        } else {
          showTips(json.code, json.msg)
        }
      })
  }
}

const fetchOptions = () => {
  fetch('../api/frps/get', { credentials: 'include', method: 'GET' })
    .then((res) => res.json())
    .then((data) => {
      frpsForm.value.options = data && data.data ? data.data : []
    })
}

const fetchFrpsGen = () => {
  frpsForm.value.isLoading = true
  downloadByPost(
    'frps生成中',
    '../api/frps/gen',
    JSON.stringify(frpsForm.value),
  ).finally(() => {
    frpsForm.value.isLoading = false
  })
}

const doClientsUpload = async (options: any) => {
  const { file } = options
  globalProgress.value = 0
  const loading = showLoading('客户端上传中...')
  await piecesUpload(
    '../api/client/upload',
    'POST',
    file,
    (progress: any) => {
      loading.setText(`客户端上传中：${progress}%`)
      globalProgress.value = parseInt(progress)
    },
    () => {
      loading.close()
      dialogClientsVisible.value = false
      showSucessTips('上传成功')
      globalProgress.value = 0
    },
  )
}

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
    .catch((error) => {
      console.error('请求失败', error)
      ok.value = false
    })
    .finally(() => {
      setTimeout(function () {
        loading.close()
        globalProgress.value = 0
        dialogFormVisible.value = false
        if (ok.value) window.location.reload()
      }, 4000)
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
    .catch((error) => console.log('更新失败', error))
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

const showlog = () => {
  window.open(`${window.origin}/log/`)
}

const fetchVersionData = () => {
  fetch('../api/version', { credentials: 'include', method: 'GET' })
    .then((res) => res.json())
    .then((json) => {
      if (json) {
        const vv = json.data as Version
        version.value = vv
        title.value = `frps ${vv.appVersion}`
        document.title = `frps ${vv.hostName}`
      }
    })
}

onUnmounted(() => {
  if (devTimer) clearTimeout(devTimer)
})
onMounted(() => {
  const jsonString = localStorage.getItem('githubKey')
  if (jsonString) {
    const obj = JSON.parse(jsonString)
    if (obj) githubApiForm.value = obj
    githubApiForm.value.isShow = false
  }
})
fetchVersionData()
fetchOptions()
</script>

<style>
:root {
  --header-height: 50px;
  --sidebar-width: 200px;
  --header-bg: #ffffff;
  --header-border: #e4e7ed;
  --sidebar-bg: #ffffff;
  --text-primary: #303133;
  --text-secondary: #606266;
  --text-muted: #909399;
  --hover-bg: #efefef;
  --content-bg: #f9f9f9;
}

html.dark {
  --header-bg: #1e1e2e;
  --header-border: #3a3d5c;
  --sidebar-bg: #1e1e2e;
  --text-primary: #e5e7eb;
  --text-secondary: #b0b0b0;
  --text-muted: #888888;
  --hover-bg: #2a2a3e;
  --content-bg: #181825;
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

/* Header */
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
  color: var(--text-primary);
  letter-spacing: -0.5px;
}

.badge {
  font-size: 11px;
  font-weight: 500;
  color: var(--text-muted);
  background: var(--hover-bg);
  padding: 2px 8px;
  border-radius: 4px;
}

.badge.server-badge {
  background: linear-gradient(135deg, #3b82f6 0%, #06b6d4 100%);
  color: white;
  border: none;
  font-weight: 500;
}

html.dark .badge.server-badge {
  background: linear-gradient(135deg, #60a5fa 0%, #22d3ee 100%);
}

.header-controls {
  display: flex;
  align-items: center;
  gap: 16px;
}

.github-link {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border-radius: 6px;
  color: var(--text-secondary);
  transition: all 0.15s ease;
  text-decoration: none;
}

.github-link:hover {
  background: var(--hover-bg);
  color: var(--text-primary);
}

.github-icon {
  width: 18px;
  height: 18px;
}

/* 设置菜单按钮，复用 github-link 的尺寸风格 */
.icon-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border: none;
  background: transparent;
  border-radius: 6px;
  color: var(--text-secondary);
  cursor: pointer;
  padding: 0;
  font-size: 16px;
  transition: all 0.15s ease;
}

.icon-btn:hover {
  background: var(--hover-bg);
  color: var(--text-primary);
}

.sidebar-divider {
  height: 1px;
  background: var(--header-border);
  margin: 8px 4px;
}

.global-progress-bar {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 9999;
  width: 100%;
}

.theme-switch {
  --el-switch-on-color: #2c2c3a;
  --el-switch-off-color: #f2f2f2;
  --el-switch-border-color: var(--header-border);
}

html.dark .theme-switch {
  --el-switch-off-color: #333;
}

.theme-switch .el-switch__core .el-switch__inner .el-icon {
  color: #909399 !important;
}

/* Layout */
.layout {
  flex: 1;
  display: flex;
  overflow: hidden;
}

/* Sidebar */
.sidebar {
  width: var(--sidebar-width);
  flex-shrink: 0;
  background: var(--sidebar-bg);
  border-right: 1px solid var(--header-border);
  padding: 16px 12px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
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
  color: var(--text-secondary);
  padding: 10px 12px;
  border-radius: 6px;
  transition: all 0.15s ease;
}

.sidebar-link:hover {
  color: var(--text-primary);
  background: var(--hover-bg);
}

.sidebar-link.active {
  color: var(--text-primary);
  background: var(--hover-bg);
  font-weight: 500;
}

/* Hamburger button (mobile only) */
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
  transition: background 0.15s ease;
}

.hamburger-btn:hover {
  background: var(--hover-bg);
}

.hamburger-icon {
  font-size: 20px;
  line-height: 1;
  color: var(--text-primary);
}

/* Mobile overlay */
.sidebar-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  z-index: 99;
}

/* Content */
#content {
  flex: 1;
  min-width: 0;
  overflow-y: auto;
  padding: 40px;
}

#content > * {
  max-width: 1024px;
  margin: 0 auto;
}

/* Common page styles */
.page-title {
  font-size: 20px;
  font-weight: 600;
  color: var(--color-text-primary, var(--text-primary));
  margin: 0;
}

.page-subtitle {
  font-size: 14px;
  color: var(--color-text-muted, var(--text-muted));
  margin: 8px 0 0;
}

/* Element Plus global overrides */
.el-button {
  font-weight: 500;
}

.el-tag {
  font-weight: 500;
}

.el-switch {
  --el-switch-on-color: #606266;
  --el-switch-off-color: #dcdfe6;
}

html.dark .el-switch {
  --el-switch-on-color: #b0b0b0;
  --el-switch-off-color: #3a3d5c;
}

.el-form-item {
  margin-bottom: 16px;
}

.el-loading-mask {
  border-radius: 8px;
}

/* Select overrides */
.el-select__wrapper {
  border-radius: 8px !important;
  box-shadow: 0 0 0 1px var(--color-border-light, #e4e7ed) inset !important;
  transition: all 0.15s ease;
}

.el-select__wrapper:hover {
  box-shadow: 0 0 0 1px var(--color-border, #dcdfe6) inset !important;
}

.el-select__wrapper.is-focused {
  box-shadow: 0 0 0 1px var(--color-border, #dcdfe6) inset !important;
}

.el-select-dropdown {
  border-radius: 12px !important;
  border: 1px solid var(--color-border-light, #e4e7ed) !important;
  box-shadow:
    0 10px 25px -5px rgba(0, 0, 0, 0.1),
    0 8px 10px -6px rgba(0, 0, 0, 0.1) !important;
  padding: 4px !important;
}

.el-select-dropdown__item {
  border-radius: 6px;
  margin: 2px 0;
  transition: background 0.15s ease;
}

.el-select-dropdown__item.is-selected {
  color: var(--color-text-primary, var(--text-primary));
  font-weight: 500;
}

/* Input overrides */
.el-input__wrapper {
  border-radius: 8px !important;
  box-shadow: 0 0 0 1px var(--color-border-light, #e4e7ed) inset !important;
  transition: all 0.15s ease;
}

.el-input__wrapper:hover {
  box-shadow: 0 0 0 1px var(--color-border, #dcdfe6) inset !important;
}

.el-input__wrapper.is-focus {
  box-shadow: 0 0 0 1px var(--color-border, #dcdfe6) inset !important;
}

/* Card overrides */
.el-card {
  border-radius: 12px;
  border: 1px solid var(--color-border-light, #e4e7ed);
  transition: all 0.2s ease;
}

/* Scrollbar */
::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: #d1d1d1;
  border-radius: 3px;
}

/* Mobile */
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
    background: var(--sidebar-bg);
    transform: translateX(-100%);
    transition: transform 0.25s cubic-bezier(0.4, 0, 0.2, 1);
    border-right: 1px solid var(--header-border);
  }

  .sidebar.mobile-open {
    transform: translateX(0);
  }

  #content {
    width: 100%;
    padding: 20px;
  }
}
</style>
