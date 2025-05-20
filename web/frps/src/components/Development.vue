<template>
  <div>
    <el-row style="margin-top: 10px">
      <el-col :span="10">
        <div style="display: flex">
          <el-input v-model="input1"></el-input>
          <el-button @click="handleGetEnv">获取环境变量</el-button>
        </div>
      </el-col>
      <el-col :span="14">
        <el-card title="日志面板" class="log-container">
          <div>
            <div ref="logContainer" class="log-container">
              <div v-for="(log, index) in logs" :key="index" class="log-item">
                {{ log }}
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

const logs = ref<string[]>([])
const logContainer = ref<HTMLDivElement | null>(null)

const input1 = ref<string>()

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
</script>

<style>
#head {
  margin-bottom: 30px;
}
</style>
