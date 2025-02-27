<template>
  <div>
    <el-row id="head">
      <el-button type="primary" @click="fetchData">刷新数据</el-button>
      <el-button type="primary" @click="uploadConfig">更新配置</el-button>
    </el-row>
    <el-input
      v-model="textarea"
      autosize
      placeholder="frpc configure file, can not be empty..."
      type="textarea"
    ></el-input>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { ElMessageBox } from 'element-plus'
import {
  showErrorTips,
  showLoading,
  showTips,
  showWarmTips,
} from '../utils/utils.ts'

let textarea = ref('')

const fetchData = () => {
  fetch('../api/server/config/get', { credentials: 'include' })
    .then((res) => {
      return res.text()
    })
    .then((text) => {
      textarea.value = text
    })
    .catch(() => {
      showErrorTips('获取配置失败')
    })
}

const uploadConfig = () => {
  ElMessageBox.confirm(
    '这个操作将更新frps服务的配置信息，确定要操作吗？',
    '注意',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    },
  )
    .then(() => {
      if (textarea.value == '') {
        showWarmTips('配置内容不能为空')
        return
      }
      const loading = showLoading('配置修改中...')
      fetch('../api/server/config/set', {
        credentials: 'include',
        method: 'PUT',
        body: textarea.value,
      })
        .then((res) => {
          return res.json()
        })
        .then((json) => {
          showTips(json.code, json.msg)
        })
        .catch(() => {
          //showErrorTips('配置失败')
        })
        .finally(() => {
          loading.close()
          setTimeout(function () {
            window.location.reload()
          }, 1000)
        })
    })
    .catch(() => {
      showErrorTips('配置失败')
    })
}

fetchData()
</script>

<style>
#head {
  margin-bottom: 30px;
}
</style>
