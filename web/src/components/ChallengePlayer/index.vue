<template>
  <div id="play-ground-editor">
    <el-row>
      <div class="btn-group">
        <el-button type="primary" size="mini" icon="el-icon-document-copy" @click="copyFullCode"></el-button>
        <el-button type="primary" :loading="loading" size="mini" @click="runTest">{{ $t('gen.run') }}</el-button>
      </div>
      <codemirror ref="myCm" :value="code" :options="cmOptions" mode="text/x-go" class="editor" @input="onCodeChange" />
      <div class="output" v-if="output !== ''">
        <pre>
<span>{{_output}}</span>
        </pre>
      </div>
    </el-row>
  </div>
</template>

<script>
import {
  evaluate
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
    code: {
      type: String,
      required: true
    },
    challengeId: {
      type: String,
      required: true
    },
    folderId: {
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
          F11 (cm) {
            cm.setOption('fullScreen', !cm.getOption('fullScreen'))
          },
          Esc (cm) {
            if (cm.getOption('fullScreen')) cm.setOption('fullScreen', false)
          }
        }
      }
    }
  },
  methods: {
    fullScreen () {
      this.codemirror.setOption(
        'fullScreen',
        !this.codemirror.getOption('fullScreen')
      )
    },
    onCodeChange (newCode) {
      this.code = newCode
      this.$emit('code-change', this.code)
    },
    copyFullCode () {
      this.$copyText(this.code).then((e) => {
        this.$message({
          type: 'success',
          message: this.$i18n.t('gen.copied')
        })
      }, (e) => {
        this.$message({
          type: 'error',
          message: this.$i18n.t('gen.copy_failed')
        })
      })
    },
    runTest () {
      this.loading = true
      this.output = ''
      var errorStr = this.$i18n.t('playground.error')
      evaluate(
        JSON.stringify({
          challenge_id: this.challengeId,
          folder_id: this.folderId,
          solution: this.code
        })
      ).then(response => {
        var data = response.data
        if (data.error !== '') {
          this.output += errorStr + ': ' + data.error + '\n'
          this.$emit('run-completed', false, this.code)
          return
        }
        var status = this.$i18n.t('playground.passed')
        if (data.tests_failed > 0 || data.status !== 0) {
          status = this.$i18n.t('playground.failed')
        }
        this.output = this.$i18n.t('playground.status') + ': ' + status + '\n'
        if (data.tests_failed > 0) {
          this.output += data.tests_failed + ' ' + this.$i18n.t('playground.test_failed') + '\n'
        }
        if (data.status > 1 && data.tests_failed === 0) {
          this.output += this.$i18n.t('playground.runtime_error') + '\n'
        }
        var problems = data.problems
        if (problems.length > 0) {
          this.output += this.$i18n.t('playground.warnings') + ': \n'
          for (var i = 0; i < problems.length; i++) {
            this.output += 'prog.go:' + problems[i].Position.Line + ':' + problems[i].Position.Column + ': ' + problems[i].Text + '\n'
          }
        }
        this.$emit('run-completed', status === this.$i18n.t('playground.passed'), this.code)
      }).catch((error) => {
        var res = error.response
        if (res !== undefined && res !== null) {
          this.output += errorStr + ': ' + res.data.message + '\n'
          return
        }
        this.output += errorStr + ': ' + error + '\n'
        this.$emit('run-completed', false, this.code)
      }).finally(() => {
        this.loading = false
      })
    }
  },
  computed: {
    codemirror () {
      return this.$refs.myCm.codemirror
    },
    _output () {
      var max = 100
      var v = this.output
      if (v.length > max) {
        v = this.output.substring(0, max) + '...'
      }
      return v
    }
  },
  watch: {
    challengeId: function (o, n) {
      this.output = ''
    }
  }
}
</script>

<style lang="scss" scoped>
@import "~codemirror/lib/codemirror.css";
@import "~codemirror/theme/eclipse.css";

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
    background-color: #1e1e1e;
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
