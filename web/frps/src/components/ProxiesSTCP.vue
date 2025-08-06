<template>
  <!--  <ProxyView :proxies="proxies" proxyType="stcp" @refresh="fetchData" />-->
  <ProxyView :proxies="proxyArray" proxyType="tcp" @refresh="fetchData" />
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ProxyConfig, STCPProxy, TCPProxy } from '../utils/proxy.js'
import ProxyView from './ProxyView.vue'

let proxies = ref<STCPProxy[]>([])
const proxyArray = ref<ProxyConfig[]>([])

const fetchData = () => {
  fetch('../api/proxy/stcp', { credentials: 'include' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      proxies.value = []
      proxyArray.value = []
      const proxiesMap = ref<Map<string, TCPProxy[]>>(new Map())
      for (let proxyStats of json.proxies) {
        const p = new STCPProxy(proxyStats)
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
        console.log(p)
        proxyArray.value.push(p)
      }
    })
}
fetchData()
</script>

<style></style>
