import { createApp } from 'vue'
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import App from './App.vue'
import router from './router'

// frp v0.70 设计体系：CSS 变量主题 + 深色模式
import './assets/css/var.css'
import './assets/css/dark.css'

const app = createApp(App)

app.use(router)

app.mount('#app')
