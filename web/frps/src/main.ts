import { createApp } from 'vue'
import 'element-plus/theme-chalk/dark/css-vars.css'
import App from './App.vue'
import router from './router'

import './assets/css/var.css'
import './assets/css/dark.css'
// 自定义组件（服务器配置/客户端配置）所需的样式
import './assets/custom.css'

const app = createApp(App)

app.use(router)

app.mount('#app')
