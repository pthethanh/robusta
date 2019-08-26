<template>
  <div id="vue-editor-js">
    <el-row>
      <div :id="holderId" class="editorjs" />
    </el-row>
    <el-row type="flex" justify="center">
      <el-dialog :visible.sync="preview" :center=true :modal=true :append-to-body=true :fullscreen=true>
        <el-row type="flex" justify="center">
          <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14">
            <view-me :data="data"></view-me>
          </el-col>
        </el-row>
      </el-dialog>
    </el-row>
  </div>
</template>

<script>
import EditorJS from '@editorjs/editorjs'
import * as configs from './config.js'
import ViewMe from '@/components/ViewMe'
import {
  getAPI
} from '@/utils/api'
import {
  uploadImageByFile,
  uploadImageByURL
} from '@/api/image'

const PLUGINS = {
  header: require('@editorjs/header'),
  list: require('@editorjs/list'),
  image: require('@editorjs/image'),
  inlineCode: require('@editorjs/inline-code'),
  embed: require('@editorjs/embed'),
  quote: require('@editorjs/quote'),
  marker: require('@editorjs/marker'),
  code: require('@editorjs/code'),
  link: require('@editorjs/link'),
  delimiter: require('@editorjs/delimiter'),
  table: require('@editorjs/table'),
  warning: require('@editorjs/warning'),
  paragraph: require('@editorjs/paragraph'),
  checklist: require('@editorjs/checklist'),
  codeExample: require('code-example')
  // raw: require('@editorjs/raw')
}

const PLUGIN_PROPS_TYPE = {
  type: [Boolean, Object],
  default: () => false,
  required: false
}

export default {
  name: 'vue-editor-js',
  components: {
    ViewMe
  },
  props: {
    holderId: {
      type: String,
      default: () => 'codex-editor',
      required: false
    },
    autofocus: {
      type: Boolean,
      default: () => false,
      required: false
    },
    initData: {
      type: Object,
      default: () => {},
      required: false
    },
    /**
     * Plugins
     */
    header: PLUGIN_PROPS_TYPE,
    list: PLUGIN_PROPS_TYPE,
    code: PLUGIN_PROPS_TYPE,
    inlineCode: PLUGIN_PROPS_TYPE,
    embed: PLUGIN_PROPS_TYPE,
    link: PLUGIN_PROPS_TYPE,
    marker: PLUGIN_PROPS_TYPE,
    table: PLUGIN_PROPS_TYPE,
    delimiter: PLUGIN_PROPS_TYPE,
    quote: PLUGIN_PROPS_TYPE,
    image: PLUGIN_PROPS_TYPE,
    warning: PLUGIN_PROPS_TYPE,
    paragraph: PLUGIN_PROPS_TYPE,
    checklist: PLUGIN_PROPS_TYPE,
    codeExample: PLUGIN_PROPS_TYPE
    // raw: PLUGIN_PROPS_TYPE
  },
  data () {
    return {
      editor: null,
      preview: false,
      data: null,
      customTools: {
        header: {
          class: PLUGINS['header'],
          inlineToolbar: ['link'],
          config: {
            placeholder: 'Header'
          },
          shortcut: 'CMD+SHIFT+H'
        },
        image: {
          class: PLUGINS['image'],
          config: {
            uploader: {
              uploadByFile (file) {
                let formData = new FormData()
                formData.append('image', file)
                return uploadImageByFile(formData)
              },
              uploadByUrl (url) {
                return uploadImageByURL({
                  url: url
                })
              }
            }
          }
        },
        list: {
          class: PLUGINS['list'],
          inlineToolbar: true,
          shortcut: 'CMD+SHIFT+L'
        },
        checklist: {
          class: PLUGINS['checklist'],
          inlineToolbar: true
        },
        quote: {
          class: PLUGINS['quote'],
          inlineToolbar: true,
          config: {
            quotePlaceholder: 'Enter a quote',
            captionPlaceholder: 'Quote\'s author'
          },
          shortcut: 'CMD+SHIFT+O'
        },
        marker: {
          class: PLUGINS['marker'],
          shortcut: 'CMD+SHIFT+M'
        },
        code: {
          class: PLUGINS['code'],
          shortcut: 'CMD+SHIFT+C'
        },
        inlineCode: {
          class: PLUGINS['inlineCode'],
          shortcut: 'CMD+SHIFT+C'
        },
        table: {
          class: PLUGINS['table'],
          inlineToolbar: true,
          shortcut: 'CMD+ALT+T'
        },
        linkTool: {
          class: PLUGINS['link'],
          inlineToolbar: true,
          config: {
            endpoint: getAPI('/api/v1/editor/fetch-url')
          }
        },
        embed: {
          class: PLUGINS['embed'],
          inlineToolbar: true,
          config: {
            services: configs.embedServices
          }
        },
        delimiter: {
          class: PLUGINS['delimiter'],
          inlineToolbar: true
        },
        warning: {
          class: PLUGINS['warning'],
          inlineToolbar: true
        },
        paragraph: {
          class: PLUGINS['paragraph'],
          inlineToolbar: true
        },
        codeExample: {
          class: PLUGINS['codeExample'],
          shortcut: 'CMD+SHIFT+X'
        }
        // raw: PLUGINS['raw']
      }
    }
  },
  mounted () {
    this.editor = new EditorJS({
      holderId: this.holderId,
      autofocus: this.autofocus,
      onReady: () => {
        this.$emit('ready')
      },
      onChange: () => {
        this.$emit('change')
      },
      data: this.initData,
      tools: this.getTools(),
      onRender: () => {
        // eslint-disable-next-line
        console.log('render')
      }
    })
  },
  methods: {
    async save () {
      const raw = await this.editor.save()
      this.$emit('save', raw)
    },
    getTools () {
      return this.customTools
    },
    async openReview () {
      this.data = await this.editor.save()
      this.preview = true
    }
  }
}
</script>

<style scoped>
.editorjs {
  background-color: white;
  width: 100%;
}

.ce-block__content {
  position: relative;
  margin: 0 auto;
}
</style>
