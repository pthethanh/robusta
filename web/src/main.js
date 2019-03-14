import Vue from 'vue'

import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import locale from 'element-ui/lib/locale/lang/en' // lang i18n

import '@/styles/index.scss' // global css

import App from './App'
import store from './store/index'
import router from './router/index'

import '@/icons' // icon
import '@/permission' // permission control
import axios from 'axios'
import './registerServiceWorker'
import './plugins/element.js'

import '../mock' // simulation data

Vue.use(ElementUI, { locale })

Vue.config.productionTip = false
// Set base URL to backend API service
const backendAddr = process.env.BIRD_BACKEND_ADDRS || 'http://localhost:8080'
console.log(`BIRD_BACKEND_ADDRS: ${backendAddr}`)
axios.defaults.baseURL = backendAddr
Vue.prototype.$http = axios

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
