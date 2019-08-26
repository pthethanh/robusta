<template>
  <div class="codex-warning">
    <el-alert :title="title" :type="type" :description="data.message" effect="light" show-icon :closable=false>
    </el-alert>
  </div>
</template>

<script>
export default {
  props: {
    'data': Object
  },
  data() {
    return {
      title: '',
      type: ''
    }
  },
  mounted() {
    // we allowed type of warning can be passed in the title in form of: type:title
    this.type = 'warning'
    this.title = this.data.title

    var supportedTypes = ['warning', 'info', 'success', 'error']
    var values = this.data.title.trim().split(':')
    if (values.length === 1) { // if no type in the title, just use default value
      return
    }
    var type = values[0]
    if (supportedTypes.indexOf(type) < 0) { // if not supported, just use default value
      return
    }
    this.type = type
    this.title = this.data.title.substr(this.type.length + 1).trim()
  }
}
</script>

<style>
.codex-warning {
  line-height: 1.75em;
  border: 1px solid #D3D3D3;
}
</style>
