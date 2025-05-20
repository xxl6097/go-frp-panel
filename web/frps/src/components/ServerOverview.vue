<template>
  <div>
    <el-row>
      <el-col :md="12">
        <div class="source">
          <el-form
            label-position="left"
            label-width="220px"
            class="server_info"
          >
            <el-form-item label="内核版本号">
              <span>{{ version?.frpcVersion }}</span>
            </el-form-item>
            <el-form-item label="面板版本">
              <span>{{ version?.appVersion }}</span>
            </el-form-item>
            <el-form-item label="编译日期">
              <span class="single-line">{{ version?.buildTime }}</span>
            </el-form-item>
            <el-form-item label="编译环境">
              <span class="single-line">{{ version?.goVersion }}</span>
            </el-form-item>
            <el-form-item label="IP地址">
              <span class="single-line">{{ version?.ipv4 }}</span>
            </el-form-item>
            <el-form-item label="网口">
              <span class="single-line">{{ version?.netName }}</span>
            </el-form-item>
            <el-form-item label="Mac地址">
              <span class="single-line">{{ version?.macAddress }}</span>
            </el-form-item>
            <el-form-item label="操作系统">
              <span class="single-line"
                >{{ version?.osType }}/{{ version?.arch }}</span
              >
            </el-form-item>
            <el-form-item label="剩余空间">
              <span class="single-line">{{ version?.freeSize }}</span>
            </el-form-item>
            <el-form-item label="使用空间">
              <span class="single-line">{{ version?.usedSize }}</span>
            </el-form-item>
            <el-form-item label="总空间">
              <span class="single-line">{{ version?.totalSize }}</span>
            </el-form-item>
            <el-form-item label="TCP绑定端口" v-if="data.bindPort != 0">
              <span>{{ data.bindPort }}</span>
            </el-form-item>
            <el-form-item label="KCP绑定端口" v-if="data.kcpBindPort != 0">
              <span>{{ data.kcpBindPort }}</span>
            </el-form-item>
            <el-form-item label="QUIC绑定端口" v-if="data.quicBindPort != 0">
              <span>{{ data.quicBindPort }}</span>
            </el-form-item>
            <el-form-item label="HTTP端口" v-if="data.vhostHTTPPort != 0">
              <span>{{ data.vhostHTTPPort }}</span>
            </el-form-item>
            <el-form-item label="HTTPS端口" v-if="data.vhostHTTPSPort != 0">
              <span>{{ data.vhostHTTPSPort }}</span>
            </el-form-item>
            <el-form-item
              label="TCPMUX 端口"
              v-if="data.tcpmuxHTTPConnectPort != 0"
            >
              <span>{{ data.tcpmuxHTTPConnectPort }}</span>
            </el-form-item>
            <el-form-item label="二级域名后缀" v-if="data.subdomainHost != ''">
              <LongSpan :content="data.subdomainHost" :length="30"></LongSpan>
            </el-form-item>
            <el-form-item label="每个代理最大连接池大小">
              <span>{{ data.maxPoolCount }}</span>
            </el-form-item>
            <el-form-item label="单个客户端最大同时存在代理数">
              <span>{{ data.maxPortsPerClient }}</span>
            </el-form-item>
            <el-form-item label="端口限制" v-if="data.allowPortsStr">
              <LongSpan :content="data.allowPortsStr" :length="30"></LongSpan>
            </el-form-item>
            <el-form-item
              label="仅接受启用 TLS 的客户端"
              v-if="data.tlsForce === true"
            >
              <span>{{ data.tlsForce }}</span>
            </el-form-item>
            <el-form-item label="心跳连接超时时间">
              <span>{{ data.heartbeatTimeout }}</span>
            </el-form-item>
            <el-form-item label="当前连接数">
              <span>{{ data.curConns }}</span>
            </el-form-item>
            <el-form-item label="客户端总数">
              <span>{{ data.clientCounts }}</span>
            </el-form-item>
            <el-form-item label="代理总数">
              <span>{{ data.proxyCounts }}</span>
            </el-form-item>
          </el-form>
        </div>
      </el-col>
      <el-col :md="12">
        <div
          id="traffic"
          style="width: 400px; height: 250px; margin-bottom: 30px"
        ></div>
        <div id="proxies" style="width: 400px; height: 250px"></div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, inject, Ref } from 'vue'
import { DrawTrafficChart, DrawProxyChart } from '../utils/chart'
import LongSpan from './LongSpan.vue'
import { showErrorTips } from '../utils/utils.ts'
import { Version } from '../utils/type.ts'

const version = inject<Ref<Version>>('version')

let data = ref({
  version: '',
  bindPort: 0,
  kcpBindPort: 0,
  quicBindPort: 0,
  vhostHTTPPort: 0,
  vhostHTTPSPort: 0,
  tcpmuxHTTPConnectPort: 0,
  subdomainHost: '',
  maxPoolCount: 0,
  maxPortsPerClient: '',
  allowPortsStr: '',
  tlsForce: false,
  heartbeatTimeout: 0,
  clientCounts: 0,
  curConns: 0,
  proxyCounts: 0,
  appName: '',
  displayName: '',
  description: '',
  appVersion: '',
  buildVersion: '',
  buildTime: '',
  gitRevision: '',
  gitBranch: '',
  goVersion: '',
})

const fetchData = () => {
  fetch('../api/serverinfo', { credentials: 'include' })
    .then((res) => res.json())
    .then((json) => {
      data.value.version = json.version
      data.value.bindPort = json.bindPort
      data.value.kcpBindPort = json.kcpBindPort
      data.value.quicBindPort = json.quicBindPort
      data.value.vhostHTTPPort = json.vhostHTTPPort
      data.value.vhostHTTPSPort = json.vhostHTTPSPort
      data.value.tcpmuxHTTPConnectPort = json.tcpmuxHTTPConnectPort
      data.value.subdomainHost = json.subdomainHost
      data.value.maxPoolCount = json.maxPoolCount
      data.value.maxPortsPerClient = json.maxPortsPerClient
      if (data.value.maxPortsPerClient == '0') {
        data.value.maxPortsPerClient = 'no limit'
      }
      data.value.allowPortsStr = json.allowPortsStr
      data.value.tlsForce = json.tlsForce
      data.value.heartbeatTimeout = json.heartbeatTimeout
      data.value.clientCounts = json.clientCounts
      data.value.curConns = json.curConns
      data.value.proxyCounts = 0
      if (json.proxyTypeCount != null) {
        if (json.proxyTypeCount.tcp != null) {
          data.value.proxyCounts += json.proxyTypeCount.tcp
        }
        if (json.proxyTypeCount.udp != null) {
          data.value.proxyCounts += json.proxyTypeCount.udp
        }
        if (json.proxyTypeCount.http != null) {
          data.value.proxyCounts += json.proxyTypeCount.http
        }
        if (json.proxyTypeCount.https != null) {
          data.value.proxyCounts += json.proxyTypeCount.https
        }
        if (json.proxyTypeCount.stcp != null) {
          data.value.proxyCounts += json.proxyTypeCount.stcp
        }
        if (json.proxyTypeCount.sudp != null) {
          data.value.proxyCounts += json.proxyTypeCount.sudp
        }
        if (json.proxyTypeCount.xtcp != null) {
          data.value.proxyCounts += json.proxyTypeCount.xtcp
        }
      }

      // draw chart
      DrawTrafficChart('traffic', json.totalTrafficIn, json.totalTrafficOut)
      DrawProxyChart('proxies', json)
    })
    .catch(() => {
      showErrorTips('获取服务器信息失败')
    })
}

// const fetchPanelData = () => {
//   fetch('../api/panelinfo', { credentials: 'include' })
//     .then((res) => res.json())
//     .then((json) => {
//       data.value.appVersion = json.data.appVersion
//       data.value.buildVersion = json.data.buildVersion
//       data.value.buildTime = json.data.buildTime
//       data.value.gitRevision = json.data.gitRevision
//       data.value.gitBranch = json.data.gitBranch
//       data.value.goVersion = json.data.goVersion
//       console.log('--->', json)
//       // draw chart
//     })
//     .catch(() => {
//       showErrorTips('获取配置失败')
//     })
// }
// fetchPanelData()
fetchData()
</script>

<style>
.source {
  border-radius: 4px;
  transition: 0.2s;
  padding-left: 24px;
  padding-right: 24px;
}

.server_info {
  margin-left: 40px;
  font-size: 0px;
}

.server_info .el-form-item__label {
  color: #99a9bf;
  height: 40px;
  line-height: 40px;
}

.server_info .el-form-item__content {
  height: 40px;
  line-height: 40px;
}

.server_info .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 100%;
}
.single-line {
  white-space: nowrap; /* 强制文本不换行 */
  overflow: hidden; /* 隐藏超出部分 */
  text-overflow: ellipsis; /* 超出部分显示为省略号 */
  display: inline-block; /* 确保宽度生效 */
  /* max-width: 200px;设置最大宽度 */
}
</style>
