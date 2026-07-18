import { createRouter, createWebHashHistory } from 'vue-router'
import ServerOverview from '../views/ServerOverview.vue'
import Clients from '../views/Clients.vue'
import ClientDetail from '../views/ClientDetail.vue'
import Proxies from '../views/Proxies.vue'
import ProxyDetail from '../views/ProxyDetail.vue'
import ServerConfig from '../views/ServerConfig.vue'
import ClientConfig from '../views/ClientConfig.vue'
import Development from '../components/Development.vue'
import ClientList from '../components/client/ClientList.vue'

const router = createRouter({
  history: createWebHashHistory(),
  scrollBehavior() {
    return { top: 0 }
  },
  routes: [
    {
      path: '/',
      name: 'ServerOverview',
      component: ServerOverview,
    },
    {
      path: '/clients',
      name: 'Clients',
      component: Clients,
    },
    {
      path: '/clients/:key',
      name: 'ClientDetail',
      component: ClientDetail,
    },
    {
      path: '/proxies/:type?',
      name: 'Proxies',
      component: Proxies,
    },
    {
      path: '/proxy/:name',
      name: 'ProxyDetail',
      component: ProxyDetail,
    },
    {
      path: '/config',
      name: 'ServerConfig',
      component: ServerConfig,
    },
    {
      path: '/user',
      name: 'ClientConfig',
      component: ClientConfig,
    },
    {
      path: '/user/list',
      name: 'ClientList',
      component: ClientList,
    },
    {
      path: '/development',
      name: 'Development',
      component: Development,
    },
  ],
})

export default router
