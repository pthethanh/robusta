<template>
  <el-avatar :size="_size" :src="src" v-bind:style="{ 'background-color': _color}">{{_symbol}}</el-avatar>
</template>

<script>
export default {
  props: {
    src: String,
    size: String,
    names: Array
  },
  computed: {
    _size () {
      if (this.size === '') {
        return 'small'
      }
      return this.size
    },
    _color () {
      var colors = ['#A7A7A7', '#20242A', '#25B7D4', '#20B9D3', '#23B6D6', '#F87442', '#F77441']
      return colors[this._symbol.charCodeAt(0) % colors.length]
    },
    _symbol () {
      var s = '^_^'
      for (var i = 0; i < this.names.length; i++) {
        if (this.names[i] !== '' && this.names[i] !== undefined) {
          s = this.names[i]
          break
        }
      }
      return s.substr(0, 1)
    }
  }
}
</script>
