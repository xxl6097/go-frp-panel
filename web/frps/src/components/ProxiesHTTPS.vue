<template>
  <!--  <ProxyView :proxies="proxies" proxyType="https" @refresh="fetchData" />-->
  <ProxyView :proxies="proxyArray" proxyType="tcp" @refresh="fetchData" />
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { HTTPSProxy, ProxyConfig, TCPProxy } from '../utils/proxy.js'
import ProxyView from './ProxyView.vue'

let proxies = ref<HTTPSProxy[]>([])
const proxyArray = ref<ProxyConfig[]>([])

const fetchData = () => {
  let vhostHTTPSPort: number
  let subdomainHost: string
  fetch('../api/serverinfo', { credentials: 'include' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      vhostHTTPSPort = json.vhostHTTPSPort
      subdomainHost = json.subdomainHost
      if (vhostHTTPSPort == null || vhostHTTPSPort == 0) {
        return
      }
      fetch('../api/proxy/https', { credentials: 'include' })
        .then((res) => {
          return res.json()
        })
        .then((json) => {
          proxies.value = []
          // for (let proxyStats of json.proxies) {
          //   proxies.value.push(
          //     new HTTPSProxy(proxyStats, vhostHTTPSPort, subdomainHost),
          //   )
          // }

          proxyArray.value = []
          const proxiesMap = ref<Map<string, TCPProxy[]>>(new Map())
          for (let proxyStats of json.proxies) {
            const p = new HTTPSProxy(proxyStats, vhostHTTPSPort, subdomainHost)
            proxies.value.push(p)
            if (proxiesMap.value.has(p.baseName)) {
              // 键存在：向现有数组追加
              proxiesMap.value.get(p.baseName)!.push(p) // 使用 ! 断言非空
            } else {
              // 键不存在：创建新数组
              proxiesMap.value.set(p.baseName, [p])
            }
          }
          for (const [key, value] of proxiesMap.value) {
            const p = new ProxyConfig(key, value)
            proxyArray.value.push(p)
          }
        })
    })
}
fetchData()
</script>

<style></style>
