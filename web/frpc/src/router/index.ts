import { createRouter, createWebHashHistory } from 'vue-router'
import Overview from '../components/Overview.vue'
import ClientConfigure from '../components/ClientConfigure.vue'
import LogView from '../components/LogView.vue'
import Development from '../components/Development.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'Overview',
      component: Overview,
    },
    {
      path: '/configure',
      name: 'ClientConfigure',
      component: ClientConfigure,
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
  ],
})

export default router
