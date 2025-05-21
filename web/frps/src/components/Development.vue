<template>
  <div>
    <el-row style="margin-top: 10px">
      <el-col :span="10">
        <div style="display: flex; margin-left: 5px; margin-right: 5px">
          <el-input
            v-model="input1"
            placeholder="请输入环境变量名称"
            style="width: 50%"
          ></el-input>
          <el-button @click="handleGetEnv" style="width: 50%"
          >获取环境变量
          </el-button>
        </div>
        <div style="display: flex; margin-left: 5px; margin-right: 5px">
          <el-input
            v-model="input2"
            style="width: 50%"
            placeholder="请输入命令"
          ></el-input>
          <el-button @click="handleCMD" style="width: 50%" plain
          >执行命令
          </el-button>
        </div>
        <div style="display: flex; margin-left: 5px; margin-right: 5px">
          <el-button @click="handleGetNetwork">获取网络信息</el-button>
        </div>
      </el-col>
      <el-col :span="14">
        <el-card title="日志面板" class="log-container">
          <div>
            <div ref="logContainer" class="log-container">
              <div v-for="(log, index) in logs" :key="index" class="log-item">
                <pre v-html="log"></pre>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { showLoading, syntaxHighlight } from '../utils/utils.ts'

const logs = ref<string[]>([])
const logContainer = ref<HTMLDivElement | null>(null)

const input1 = ref<string>()
const input2 = ref<string>()

const addLog = (context: string): void => {
  const newLog = `${new Date().toLocaleString()}: ${context}\r\n`
  logs.value.unshift(newLog)
  // 滚动到顶部
  if (logContainer.value) {
    logContainer.value.scrollTop = 0
  }
}

const handleGetEnv = () => {
  fetch(`../api/env?name=${input1.value}`, { credentials: 'include' })
    .then((res) => {
      return res.text()
    })
    .then((text) => {
      addLog(text)
    })
    .catch((err) => {
      addLog(err)
    })
}

const handleCMD = () => {
  fetchRunApi('cmd', { data: input2.value })
}

const handleGetNetwork = () => {
  fetchRunApi('network', {})
}

const fetchRunApi = (action: string, data: any) => {
  const body = {
    action: action,
    data: data,
  }
  console.log('body', body)
  const loading = showLoading('请求中...')
  console.log('fetchApi', body)
  fetch('../api/run', {
    credentials: 'include',
    method: 'POST',
    body: JSON.stringify(body),
  })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      console.log(json)
      if (json.code === 0) {
        const rawJson = JSON.stringify(json.data, null, 2)
        const highlightedJSON = syntaxHighlight(rawJson)
        addLog(highlightedJSON)
      } else {
        addLog(json.msg)
      }
    })
    .catch(() => {
      //showErrorTips('配置失败')
    })
    .finally(() => {
      loading.close()
    })
}
</script>

<style>
#head {
  margin-bottom: 30px;
}

.log-container {
  height: auto;
  max-height: 500px;
  overflow-y: auto;
  margin-left: 20px;
}

.log-item {
  margin-bottom: 5px;
}
</style>
