<template>
  <div id="play-ground-editor">
    <el-row type="flex">
      <el-col :span="18">
        <el-row>
          <div class="btn-group">
            <el-button type="primary" size="mini" icon="el-icon-document-copy"></el-button>
            <el-button type="primary" size="mini">Run</el-button>
          </div>
          <codemirror ref="myCm" :value="code" :options="cmOptions" mode="text/x-go" class="editor" />
        </el-row>

      </el-col>
      <el-col :span="6">
        <div class="output" v-if="output !== ''">
          <pre>
<span>{{output}}</span>
<span v-if="exit" style="color:yellow; font-style:italic">Program exited.</span>
        </pre>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {
  runCode
} from '@/api/playground'
import 'codemirror/mode/go/go.js'
import {
  codemirror
} from 'vue-codemirror'

export default {
  components: {
    codemirror
  },
  props: {
    'code': {
      type: String,
      required: true
    }
  },
  data () {
    return {
      output: '',
      playing: null,
      showOutput: false,
      loading: false,
      exit: false,
      cmOptions: {
        tabSize: 4,
        mode: 'text/x-go',
        theme: 'eclipse',
        lineNumbers: true,
        line: true,
        extraKeys: {
          'F11'(cm) {
            cm.setOption('fullScreen', !cm.getOption('fullScreen'))
          },
          'Esc'(cm) {
            if (cm.getOption('fullScreen')) cm.setOption('fullScreen', false)
          }
        }
      }

    }
  },
  methods: {
    fullScreen () {
      this.codemirror.setOption('fullScreen', !this.codemirror.getOption('fullScreen'))
    },
    updateOutput (data) {
      this.showOutput = true
      this.exit = false
      var self = this
      var timeout

      function next () {
        if (!data.events || data.events.length === 0) {
          if (data.is_test) {
            if (data.tests_failed > 0) {
              self.output += '\n' + data.tests_failed + ' test(s) failed.'
            } else {
              self.output += '\nAll tests passed.'
            }
          } else {
            if (data.status > 0) {
              self.output += '\nstatus: ' + data.status + '.'
            } else {
              if (data.errors !== '') {
                self.output += '\nerror: ' + data.errors + '.'
              } else {
                // self.output += '\nProgram exited.'
                self.exit = true
              }
            }
          }
          return
        }
        var e = data.events.shift()
        timeout = setTimeout(function () {
          self.output += e.message
          next()
        }, e.delay / 1000000)
      }
      next()
      return {
        Stop: function () {
          clearTimeout(timeout)
        }
      }
    },
    run () {
      this.showOutput = true
      this.loading = true
      if (this.playing != null) {
        this.playing.Stop()
      }
      this.output = ''
      var self = this
      runCode(JSON.stringify({
        code: this.code
      })).then(function (response) {
        self.playing = self.updateOutput(response)
      }).catch(function (error) {
        // eslint-disable-next-line
        console.log(error)
        self.showOutput = false
      }).finally(function () {
        self.loading = false
      })
    },
    onCodeChange (newCode) {
      this.code = newCode
      this.$emit('code-change', this.code)
    },
    copyFullCode () {
      let self = this
      this.$copyText(this.code).then(function (e) {
        self.$message({
          type: 'success',
          message: 'Full runable code is copied into clipboard'
        })
      }, function (e) {
        self.$message({
          type: 'error',
          message: 'Failed to copy full runable code to clipboard'
        })
      })
    }
  },
  computed: {
    codemirror () {
      return this.$refs.myCm.codemirror
    }
  }
}
</script>

<style lang="scss" scoped>
@import '~codemirror/lib/codemirror.css';
@import '~codemirror/theme/eclipse.css';

.playground {
  .btn-group {
    position: absolute;
    right: 0px;
    z-index: 99999;
    top: 5px;
  }

  .output {
    border-top: 1px solid lightgrey;
    text-align: start;
    background-color: #1E1E1E;
    color: white;

    .pre {
      display: block;
      font-family: monospace;
      white-space: pre;
      font-size: 11px;
      line-height: 1.3em;
    }
  }
}
</style>
