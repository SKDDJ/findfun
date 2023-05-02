import Vue from 'vue';
import App from './App.vue';
import { createRouter, createWebHistory } from 'vue-router';
import { createI18n } from 'vue-i18n';
import Homepage from './components/Homepage.vue';
import PublishGathering from './components/PublishGathering.vue';
import SearchGathering from './components/SearchGathering.vue';
import MatchGathering from './components/MatchGathering.vue';
import EvaluateSystem from './components/EvaluateSystem.vue';

// Import Naive UI styles and components
import { create, NConfigProvider } from 'naive-ui';
import 'naive-ui/lib/naive-ui.css';
const naive = create();
Vue.use(naive);

const routes = [
  { path: '/', component: Homepage },
  { path: '/publish', component: PublishGathering },
  { path: '/search', component: SearchGathering },
  { path: '/match', component: MatchGathering },
  { path: '/evaluate', component: EvaluateSystem },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

const i18n = createI18n({
  locale: 'zh-CN',
  messages: {
    'zh-CN': {
      // Add translations here for simplified Chinese
    },
  },
});

Vue.config.productionTip = false;

new Vue({
  router,
  i18n,
  render: (h) => h(App),
}).$mount('#app');
