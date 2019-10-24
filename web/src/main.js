import Vue from 'vue'
import { i18n } from './plugins/i18n.js'
import request from '@/utils/request'
import '@/icons'
import '@/permission'
import './registerServiceWorker'
import './plugins/element.js'
import './plugins/highlight.js'
import './plugins/page_title.js'
import './plugins/hotkeys.js'
import './plugins/clipboard.js'

import * as filters from './filters'
import '@/styles/index.scss' // global css

import App from './App'
import store from './store/index'
import router from './router/index'

Vue.prototype.$http = request

// register global filters
Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

new Vue({ // eslint-disable-line no-new
  el: '#app',
  router,
  i18n,
  store,
  render: h => h(App)
})
