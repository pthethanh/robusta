import Vue from 'vue'

import VueHighlightJS from 'vue-highlight.js'
import go from 'highlight.js/lib/languages/go'
import 'highlight.js/styles/github.css'

// Highlight.js languages (Only required languages)
Vue.use(VueHighlightJS, {
  languages: {
    go
  }
})
