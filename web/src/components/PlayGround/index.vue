<template>
  <div>
    <codemirror ref="myCm" :value="code" :options="cmOptions" mode="text/x-go" class="code-editor" />
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

<style scoped>
@import '~codemirror/lib/codemirror.css';
@import '~codemirror/theme/eclipse.css';
</style>
