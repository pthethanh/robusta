<template>
  <div class="sidebar">
    <el-menu default-active="2"
      class="el-menu-vertical-demo"
      v-if="isCollapse"
      :collapse="false"
      :router="true"
      >
      <el-menu-item index="1" route="/">Home</el-menu-item>
      <el-menu-item index="3" route="/tutorials">Tutorials</el-menu-item>
    </el-menu>
  </div>
</template>

<style>
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
  }
</style>

<script>
import { mapGetters } from 'vuex'
export default {
  computed: {
    ...mapGetters([
      'sidebar'
    ]),
    isCollapse () {
      return !this.sidebar.opened
    }
  },
  methods: {
    handleOpen (key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose (key, keyPath) {
      console.log(key, keyPath)
    }
  }
}
</script>

<style>
  .sidebar {
  }
</style>
