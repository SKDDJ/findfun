import { createApp } from 'vue'
import { createI18n } from 'vue-i18n'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import routes from './router/index.js'
import zhCN from './locales/zh-CN.json'

const i18n = createI18n({
  locale: 'zh-CN',
  messages: {
    'zh-CN': zhCN,
  },
})

const router = createRouter({
  history: createWebHistory(),
  routes,
})

createApp(App)
  .use(i18n)
  .use(router)
  .mount('#app')
