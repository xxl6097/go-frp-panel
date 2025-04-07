<template>
  <div>
    <div v-if="showUpgradeDialog" class="upgrade-popup-overlay">
      <div class="upgrade-popup">
        <div class="upgrade-popup-header">
          <h3>❤️ 发现新版本</h3>
          <button @click="handleClose" class="close-button">×</button>
        </div>
        <div class="upgrade-popup-content" v-html="updateContent"></div>
        <div class="upgrade-popup-footer">
          <el-button @click="handleClose">稍后提醒</el-button>
          <el-button type="primary" @click="handleConfirm">立即升级</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, defineExpose } from 'vue'
import {
  markdownToHtml,
  showInfoTips,
  showLoading,
  showTips,
} from '../utils/utils.ts'
const showUpgradeDialog = ref(false)

const showUpdateDialog = (message: string, binurl: string) => {
  showUpgradeDialog.value = true
  updateContent.value = markdownToHtml(message)
  binUrl.value = binurl
}

const upgradeByUrl = (binUrl: string) => {
  console.log('upgradeByUrl', binUrl)
  const loading = showLoading('程序升级中...')
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
      loading.close()
      setTimeout(function () {
        window.location.reload()
      }, 4000)
    })
}

const checkVersion = () => {
  fetch('../api/checkversion', { credentials: 'include' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      if (json.code === 0) {
        showInfoTips(json.msg)
      } else if (json.code === 1) {
        console.log('--------->', showUpgradeDialog.value)
        showUpdateDialog(json.msg, json.data)
      }
    })

  //   showUpdateDialog(`### ✨ 新特性
  //
  // * 程序以服务形式安装并运行，支持跨平台windows、linux、macos平台；
  // * 新增重启功能，用户可管理后台操作重启；
  // * 新增在线升级功能，可上传式升级和文件url式升级；
  // * 新增可在管理后台端查看日志功能；
  // * frps服务端可生成frpc客户端，密钥信息二进制内嵌在客户端程序中；
  // * 新增用户配置，可以配置授权用户供frpc端使用
  // * frpc客户端可运行多客户端
  // * 新增frpc用户配置导入导出
  //
  // ### ⚙️ 问题修复
  //
  // * Properly release resources in service.Close() to prevent resource leaks when used as a library.
  // ---
  // ### 🚀 github加速
  //
  // \`\`\`
  // [
  //   "https://ghfast.top/",
  //   "https://gh-proxy.com/",
  //   "https://ghproxy.1888866.xyz/"
  // ]
  // \`\`\`
  // `, '')
}

// 暴露方法供父组件调用
defineExpose({
  openUpgradeDialog: checkVersion,
})
const binUrl = ref<string>()
const updateContent = ref<string>()

const handleConfirm = () => {
  showUpgradeDialog.value = false
  upgradeByUrl(binUrl.value as string)
}

const handleClose = () => {
  showUpgradeDialog.value = false
  console.log('handleClose', showUpgradeDialog.value)
}

checkVersion()
</script>
<style scoped>
.upgrade-popup-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
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
</style>
