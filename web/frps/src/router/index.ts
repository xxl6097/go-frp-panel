import { createRouter, createWebHashHistory } from 'vue-router'
import ServerOverview from '../components/ServerOverview.vue'
import ProxiesTCP from '../components/ProxiesTCP.vue'
import ProxiesUDP from '../components/ProxiesUDP.vue'
import ProxiesHTTP from '../components/ProxiesHTTP.vue'
import ProxiesHTTPS from '../components/ProxiesHTTPS.vue'
import ProxiesTCPMux from '../components/ProxiesTCPMux.vue'
import ProxiesSTCP from '../components/ProxiesSTCP.vue'
import ProxiesSUDP from '../components/ProxiesSUDP.vue'
import ServerConfig from '../components/ServerConfig.vue'
import UserConfig from '../components/UserConfig.vue'
import LogView from '../components/LogView.vue'
import ClientList from '../components/client/ClientList.vue'
import Development from '../components/Development.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'ServerOverview',
      component: ServerOverview,
    },
    {
      path: '/config',
      name: 'ServerConfig',
      component: ServerConfig,
    },
    {
      path: '/user',
      name: 'UserConfig',
      component: UserConfig,
    },
    {
      path: '/user/list',
      name: 'ClientList',
      component: ClientList,
      props: (to: any) => ({
        profile: JSON.parse(to.query.profileData || '{}'),
      }),
    },
    {
      path: '/log',
      name: 'LogView',
      component: LogView,
    },
    {
      path: '/development',
      name: 'Development',
      component: Development,
    },
    {
      path: '/proxies/tcp',
      name: 'ProxiesTCP',
      component: ProxiesTCP,
    },
    {
      path: '/proxies/udp',
      name: 'ProxiesUDP',
      component: ProxiesUDP,
    },
    {
      path: '/proxies/http',
      name: 'ProxiesHTTP',
      component: ProxiesHTTP,
    },
    {
      path: '/proxies/https',
      name: 'ProxiesHTTPS',
      component: ProxiesHTTPS,
    },
    {
      path: '/proxies/tcpmux',
      name: 'ProxiesTCPMux',
      component: ProxiesTCPMux,
    },
    {
      path: '/proxies/stcp',
      name: 'ProxiesSTCP',
      component: ProxiesSTCP,
    },
    {
      path: '/proxies/sudp',
      name: 'ProxiesSUDP',
      component: ProxiesSUDP,
    },
  ],
})

export function registerRoute(name: string, path: string, component: any) {
  router.addRoute({
    path: path,
    name: name,
    component: component,
  })
}

export default router
