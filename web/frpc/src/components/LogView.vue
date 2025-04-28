<template>
  <div class="homewrap">
    <div style="text-align: left; margin-left: 5px">
      <div style="display: flex; width: auto">
        <el-text size="small" type="primary">请输入SSE地址：</el-text>
        <el-input
          size="small"
          style="width: 250px"
          v-model="ssehost"
          placeholder="请输入SSE地址"
        ></el-input>
        <el-button
          style="margin-left: 5px"
          :type="datas.btncolor"
          size="small"
          @click="onStart()"
        >
          {{ datas.btntext }}
        </el-button>
        <el-button plain size="small" @click="onClearLog()">清空日志</el-button>
        <el-button plain size="small" @click="addLog">添加日志</el-button>
        <el-button plain size="small" @click="showDir">显示目录</el-button>
      </div>
    </div>

    <el-container>
      <div style="width: 100%">
        <div
          style="
            padding: 5px;
            width: 100%;
            height: 90%;
            overflow: auto;
            word-break: break-all;
          "
        >
          <el-card title="日志面板">
            <el-scrollbar ref="scrollbarRef" style="height: 700px">
              <div
                id="txtContent"
                ref="txtContent"
                v-for="(log, index) in logs"
                :key="index"
                :style="{
                  color: getLogColor(log),
                  textAlign: 'left',
                  overflow: 'auto',
                  wordWrap: 'break-word',
                }"
              >
                {{ log }}
              </div>
            </el-scrollbar>
          </el-card>
        </div>
      </div>
    </el-container>
  </div>

  <!--弹窗显示文件目录-->
  <el-dialog v-model="showFileDirDialog" width="700">
    <template #default>
      <div
        style="
          text-align: left;
          border: solid 1px #d9dede;
          box-shadow:
            0 2px 4px rgba(0, 0, 0, 0.12),
            0 0 6px rgba(0, 0, 0, 0.04);
          padding: 5px;
          margin-top: 5px;
        "
      >
        <el-scrollbar max-height="700px">
          <ul>
            <li>
              <el-text @click="onFileClick('..')">..</el-text>
            </li>
            <li v-for="item in files" :key="item.id">
              <el-text @click="onFileClick(item.label)">{{ item.label }}</el-text>
            </li>
          </ul>
        </el-scrollbar>

      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { onMounted, ref, nextTick } from 'vue'
import { Action, ElMessage, ElMessageBox, ElScrollbar } from 'element-plus'
const logs = ref<string[]>([])
const scrollbarRef = ref<InstanceType<typeof ElScrollbar>>()
const loading = ref<boolean>(false)
const status = ref<boolean>(false)
const apihost = ref<string>(window.location.origin)
const ssehost = ref<string>('')
const source = ref<EventSource>()
const datas = ref({
  btntext: '打开',
  btncolor: 'primary',
  websock: null,
})
interface Tree {
  id: string
  label: string
  children?: Tree[]
}


const files = ref<Tree[]>([])
const showFileDirDialog = ref<boolean>(false)
const logpath = ref<string>('/')

function getLogColor(logstring: string) {
  if (logstring.includes('[INFO]')) {
    return 'green'
  } else if (logstring.includes('[WARN]')) {
    return 'orange'
  } else if (logstring.includes('[ERROR]')) {
    return 'red'
  } else if (logstring.includes('[FATAL]')) {
    return 'blue'
  } else {
    return ''
  }
}

function addLogContent(content: string) {
  logs.value.push(content)
  nextTick(() => {
    if (scrollbarRef.value) {
      const scrollContainer = scrollbarRef.value.$el.querySelector(
        '.el-scrollbar__wrap',
      )
      if (scrollContainer) {
        scrollContainer.scrollTop = scrollContainer.scrollHeight
      }
    }
  })
}

function showDir() {
  showFileDirDialog.value = true
  fetchData(logpath.value)
}
function addLog() {
  const logTypes = ['[INFO]', '[WARN]', '[ERROR]', '[DEBUG]']
  const randomType = logTypes[Math.floor(Math.random() * logTypes.length)]
  addLogContent(`新的 ${randomType} 日志，时间：${new Date().toLocaleString()}`)
}

function initSSE() {
  if (ssehost.value === '') {
    loading.value = false
    ElMessageBox.alert('请填写ws地址哦～～', '警告', {
      // if you want to disable its autofocus
      // autofocus: false,
      confirmButtonText: 'OK',
      callback: (action: Action) => {
        ElMessage({
          type: 'info',
          message: `action: ${action}`,
        })
      },
    })
    return
  }
  try {
    showLog(`开始连接SSE:${ssehost.value}`)
    const s = new EventSource(ssehost.value)
    source.value = s
    s.onmessage = (event) => {
      console.log('收到消息:', event.data)
      showLog(event.data)
    }
    s.onopen = (e) => {
      console.log('SSE连接已建立', s.readyState) // readyState=1表示连接正常
      datas.value.btncolor = 'danger'
      datas.value.btntext = '关闭'
      loading.value = false
      status.value = true
      showLog('连接成功 ' + e.currentTarget?.toString())
      console.log('sse connect sucessully..', e)
    }
    s.onerror = (e) => {
      source.value?.close()
      source.value = undefined
      datas.value.btncolor = 'primary'
      datas.value.btntext = '打开'
      loading.value = false
      status.value = false
      showLog('连接错误:' + JSON.stringify(e))
      console.log('onerror received a message', e)
    }
  } catch (e) {
    console.log('sse init err', e)
    loading.value = false
    showLog(`连接SSE识别:${JSON.stringify(e)}`)
  }
}

function showLog(e: string) {
  console.log(e)
  addLogContent(e)
}
function onStart() {
  console.log('onStart', loading.value)
  showLog(`onStart:${loading.value}`)
  if (!status.value) {
    loading.value = true
    initSSE()
  } else {
    showLog(`onStart open :${loading.value}`)
    source.value?.close()
    source.value = undefined
    datas.value.btncolor = 'primary'
    datas.value.btntext = '打开'
    loading.value = false
    status.value = false
  }
}
function onClearLog() {
  logs.value = []
}

function onFileClick(e: string) {
  if (e === '..') {
    if (logpath.value === '/') return
    let list = logpath.value.split('/')
    let api = ''
    list.forEach((value, index, array) => {
      if (index < array.length - 2) {
        console.log('forEach', value, index, array.length) // Banana, Ma
        api = api.concat(value, '/')
      }
    })
    logpath.value = api
  } else {
    if (e.endsWith('/')) {
      logpath.value = logpath.value.concat(e)
    } else {
      window.open(apihost.value + '/fserver/' + logpath.value + e, '_blank')
      return
    }
  }
  fetchData(logpath.value)
}

function fetchData(path: string) {
  const data = {
    path: path,
  }
  const body = JSON.stringify(data)
  fetch('../api/files', {
    credentials: 'include',
    body: body,
    method: 'PUT',
  })
    .then((res) => {
      let isText = res.headers.get('File-Type')
      console.log('1-fetchData', res, res.statusText, isText)
      if (isText == 'text') {
        console.log('2-fetchData', res)
        return res.text()
      } else {
        return res.json()
      }
    })
    .then((json) => {
      console.log('4-fetchData', json)
      if (json && json.code === 0) {
        files.value = json.data
      }
    })
    .catch((err) => {
      ElMessage({
        showClose: true,
        message: 'Get status failed!' + err,
        type: 'warning',
      })
      showLog(err)
      console.log('3-fetchData', err)
    })
}

onMounted(() => {
  console.log('mounted', window.location)
  console.log('host', window.location.host)
  console.log('origin', window.location.origin)
  console.log('pathname', window.location.pathname)
  console.log('protocol', window.location.protocol)

  let url = window.location.pathname
  let list = url.split('/')
  console.log('list', list) // Banana, Ma
  let api = ''
  list.forEach((value, index, array) => {
    if (index > 0 && index < array.length - 2) {
      console.log('forEach', value, index, array.length) // Banana, Ma
      api = api.concat(value, '/')
    }
  })

  ssehost.value = `${window.location.origin}/api/sse-stream`
  //ssehost.value.concat(window.location.origin,'api/sse-stream')
  console.log('ssehost', ssehost.value) // Banana, Ma

  apihost.value = apihost.value.concat(api)
  console.log('apihost', apihost.value) // Banana, Ma
  initSSE()
})
</script>

<style scoped>
.homewrap {
  text-align: center;
}

.el-container {
  height: 880px;
}

.el-container .el-form-item {
  margin-bottom: 1px;
}

.el-aside {
  margin-left: 5px;
  margin-right: 5px;
}

.el-main {
  background-color: #a0dce6;
}

.el-tag {
  background-color: #409eff;
  width: 100%;
  color: #ffffff;
  font-size: 18px;
  margin-bottom: 4px;
  text-align: center;
}

.rightMenu {
  position: fixed;
  z-index: 99999999;
  cursor: pointer;
  border: 1px solid #eee;
  box-shadow: 0 0.5em 1em 2px rgba(0, 0, 0, 0.1);
  border-radius: 6px;
  color: #1a1a1a;
}

.rightMenu ul {
  list-style: none;
  margin: 0;
  padding: 0;
  border-radius: 6px;
}

.rightMenu ul li {
  padding: 6px 10px;
  background: #fff;
  border-bottom: 1px solid #000;
  box-sizing: border-box;
  display: flex;
  align-items: center;
  justify-content: space-around;
}

.rightMenu ul li:last-child {
  border: none;
}

.rightMenu ul li:hover {
  transition: all 1s;
  background: #92c9f6;
}

/* 为 li 元素添加下横线和蓝色字体样式 */
ul li {
  text-decoration: underline;
  cursor: pointer;
}
</style>
