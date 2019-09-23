<template>
  <div id="play-ground-editor">
    <el-row>
      <div class="btn-group">
        <el-button type="primary" size="mini" icon="el-icon-document-copy" @click="copyFullCode"></el-button>
        <el-button type="primary" :loading="loading" size="mini" @click="runTest">Run</el-button>
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
    challenge_id: {
      type: String,
      required: true
    },
    folder_id: {
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
      let self = this
      this.$copyText(this.code).then(
        function (e) {
          self.$message({
            type: 'success',
            message: 'Full runable code is copied into clipboard'
          })
        },
        function (e) {
          self.$message({
            type: 'error',
            message: 'Failed to copy full runable code to clipboard'
          })
        }
      )
    },
    runTest () {
      this.loading = true
      this.output = ''
      evaluate(
        JSON.stringify({
          challenge_id: this.challenge_id,
          folder_id: this.folder_id,
          solution: this.code
        })
      ).then(response => {
        var data = response.data
        if (data.error !== '') {
          this.output += 'Error: ' + data.error + '\n'
          return
        }
        var status = 'PASSED'
        if (data.is_test_failed) {
          status = 'FAILED'
        }
        this.output = 'Status: ' + status + '\n'
        if (data.tests_failed > 0) {
          this.output += data.tests_failed + ' tests failed\n'
        }
        var problems = data.problems
        if (problems.length > 0) {
          this.output += 'Warnings: \n'
          for (var i = 0; i < problems.length; i++) {
            this.output += 'prog.go:' + problems[i].Position.Line + ':' + problems[i].Position.Column + ': ' + problems[i].Text + '\n'
          }
        }
        this.$emit('run-completed', status === 'PASSED')
      }).catch((error) => {
        var res = error.response
        if (res !== undefined && res !== null) {
          this.output += 'Error: ' + res.data.message + '\n'
          return
        }
        this.output += 'Error: ' + error + '\n'
        this.$emit('run-completed', false)
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
