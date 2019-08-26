<template>
  <div>
    <el-tag :key="tag" v-for="tag in tags" closable :disable-transitions="false" @close="handleClose(tag)" effect="plain">
      {{tag}}
    </el-tag>
    <el-input class="input-new-tag" v-if="inputVisible" v-model="inputValue" ref="saveTagInput" size="mini" @keyup.enter.native="handleInputConfirm" @blur="handleInputConfirm">
    </el-input>
    <el-button v-else class="button-new-tag" size="small" @click="showInput">+Tag</el-button>
  </div>
</template>

<script>
export default {
  props: {
    'initData': Array
  },
  data() {
    return {
      tags: [],
      inputVisible: false,
      inputValue: ''
    }
  },
  mounted() {
    if (this.initData !== null) {
      this.tags = this.initData
    }
  },
  methods: {
    handleClose(tag) {
      this.tags.splice(this.tags.indexOf(tag), 1)
    },
    showInput() {
      this.inputVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagInput.$refs.input.focus()
      })
    },
    handleInputConfirm() {
      let inputValue = this.inputValue
      if (inputValue) {
        this.tags.push(inputValue)
      }
      this.inputVisible = false
      this.inputValue = ''
    },
    getTags() {
      return this.tags
    }
  }
}
</script>

<style scoped>
.el-tag+.el-tag {
  margin-left: 0px;
}

.button-new-tag {
  margin-left: 0px;
  height: 32px;
  line-height: 30px;
  padding-top: 0;
  padding-bottom: 0;
}

.input-new-tag {
  width: 90px;
  margin-left: 0px;
  vertical-align: bottom;
}
</style>
