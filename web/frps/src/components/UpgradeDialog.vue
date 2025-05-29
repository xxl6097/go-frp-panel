<template>
  <el-dialog v-model="showUpgradeDialog" title="❤️ 发现新版本">
    <div class="upgrade-popup-content" v-html="updateContent"></div>
    <template #footer>
      <el-button @click="handleClose">稍后提醒</el-button>
      <el-button type="warning" @click="handleConfirm" v-if="patchUrl !== ''"
        >差量升级
      </el-button>
      <el-button type="primary" @click="handleConfirm"
        >{{ patchUrl === '' ? '升级' : '全量升级' }}
      </el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, defineExpose } from 'vue'
import {
  markdownToHtml,
  showInfoTips,
  showLoading,
  showTips,
} from '../utils/utils.ts'
import { ElButton } from 'element-plus'

const showUpgradeDialog = ref(false)
const patchUrl = ref<string>()
const binUrl = ref<string>()
const updateContent = ref<string>()

const showUpdateDialog = (
  patchurl: string,
  fullurl: string,
  message: string,
) => {
  showUpgradeDialog.value = true
  updateContent.value = markdownToHtml(message)
  binUrl.value = fullurl
  patchUrl.value = patchurl
  console.log('fullurl', binUrl)
  console.log('patchUrl', patchurl)
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
      if (json.code === 0) {
        setTimeout(function () {
          window.location.reload()
        }, 1000)
      }
    })
    .catch((error) => {
      console.log('更新失败', error)
      //showWarmTips('更新失败' + JSON.stringify(error))
    })
    .finally(() => {
      loading.close()
    })
}

const checkVersion = () => {
  fetch('../api/checkversion', { credentials: 'include' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      console.log('checkVersion', json)
      if (json.code === 0) {
        showUpdateDialog(
          json.data.patchUrl,
          json.data.fullUrl,
          json.data.releaseNotes,
        )
      } else {
        showInfoTips(json.msg)
      }
    })
}

// 暴露方法供父组件调用
defineExpose({
  openUpgradeDialog: checkVersion,
})

const handleConfirm = () => {
  showUpgradeDialog.value = false
  if (patchUrl.value !== '') {
    upgradeByUrl(patchUrl.value as string)
  } else {
    upgradeByUrl(binUrl.value as string)
  }
}

const handleClose = () => {
  showUpgradeDialog.value = false
  console.log('handleClose', showUpgradeDialog.value)
}

// checkVersion()
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
  z-index: 9999; /* 设置较高的 z-index 值，确保在最顶部 */
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
