<template>
  <div class="code-ex">
    <el-row>
      <div class="code-ex-btn-group">
        <el-button @click="copyFullCode" type="primary" size="mini" v-if="code != ''" icon="el-icon-document-copy"></el-button>
        <el-button @click="run" type="primary" size="mini" v-if="code != ''">{{ $t('gen.run') }}</el-button>
      </div>
      <div class="code-ex-input">
        <highlight-code lang="go">{{sample}}</highlight-code>

      </div>
    </el-row>
    <el-row v-show="showOutput" v-loading="loading">
      <div class="code-ex-output">
        <pre>
<span>{{output}}</span>
<span v-if="exit" style="color:yellow; font-style:italic">{{ $t('playground.program_exit') }}</span>
        </pre>
      </div>
    </el-row>
  </div>
</template>

<script>
import {
  runCode
} from '@/api/playground'
export default {
  components: {

  },
  props: {
    'code': {
      type: String,
      required: true
    },
    'sample': {
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
      exit: false
    }
  },
  methods: {
    updateOutput (data) {
      this.showOutput = true
      this.exit = false
      var self = this
      var timeout

      function next () {
        if (!data.events || data.events.length === 0) {
          if (data.is_test) {
            if (data.tests_failed > 0) {
              self.output += '\n' + data.tests_failed + ' ' + this.$i18n.t('playground.tests_failed')
            } else {
              self.output += '\n' + this.$i18n.t('playground.all_tests_passed')
            }
          } else {
            if (data.status > 0) {
              self.output += '\n' + this.$i18n.t('playground.status') + ': ' + data.status + '.'
            } else {
              if (data.errors !== '') {
                self.output += '\n' + this.$i18n.t('playground.error') + ': ' + data.errors + '.'
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
      this.$copyText(this.code).then((e) => {
        self.$message({
          type: 'success',
          message: this.$i18n.t('gen.copied')
        })
      }, (e) => {
        self.$message({
          type: 'error',
          message: this.$i18n.t('gen.copy_failed')
        })
      })
    }
  },
  watch: {
    code: function (o, n) {
      this.output = ''
      this.exit = false
      this.showOutput = false
    }
  }
}
</script>

<style scoped>
.code-ex .code-ex-input code,
.code-ex .code-ex-input pre,
.code-ex .code-ex-input div {
  background-color: #FBFFDA;
}

.code-ex-input pre {
  margin: 0px;
}

.code-ex-btn-group {
  position: absolute;
  right: 0px;
}

.code-ex-output {
  border-top: 1px solid lightgrey;
  text-align: start;
  background-color: #1E1E1E;
  color: white;
}

.code-ex-output pre {
  display: block;
  font-family: monospace;
  white-space: pre;
  font-size: 11px;
  line-height: 1.3em;
}
</style>
