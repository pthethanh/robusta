<template>
  <div class="sidebar">
    <el-menu default-active="1" v-bind:class="{
      'el-menu-vertical-demo': !isCollapse,
      'el-menu-vertical-demo hidden-sm-and-down': isCollapse
      }" @open="handleOpen" @close="handleClose" :collapse="isCollapse" :router=true>
      <el-menu-item index="1" route="/">
        <i class="el-icon-s-home"></i>
        <span slot="title">Blog</span>
      </el-menu-item>
      <el-menu-item index="2" route="/resources">
        <i class="el-icon-reading"></i>
        <span slot="title">Resources</span>
      </el-menu-item>
      <el-menu-item index="3" route="/articles/new">
        <i class="el-icon-circle-plus"></i>
        <span slot="title">Post</span>
      </el-menu-item>
      <el-menu-item index="4" route="/about">
        <i class="el-icon-info"></i>
        <span slot="title">About</span>
      </el-menu-item>
      <el-menu-item index="5" @click="onLogin()" v-if="!user.authenticated">
        <i class="el-icon-s-custom"></i>
        <span slot="title">Sign in</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<style lang="scss" scoped>
.sidebar {
  z-index: 1;
  position: fixed;
  height: 100%;

  .el-menu-vertical-demo {
    height: 100%;

    &:not(.el-menu--collapse) {
      width: 200px;
      height: 100%;
    }
  }
}
</style>

<script>
import {
  mapGetters
} from 'vuex'
export default {
  computed: {
    ...mapGetters([
      'sidebar',
      'user'
    ]),
    isCollapse () {
      return !this.sidebar.opened
    }
  },
  methods: {
    handleOpen (key, keyPath) {
      // eslint-disable-next-line
      console.log(key, keyPath)
    },
    handleClose (key, keyPath) {
      // eslint-disable-next-line
      console.log(key, keyPath)
    },
    onLogin () {
      this.$store.dispatch('ToggleLogin', true)
    }
  }
}
</script>
