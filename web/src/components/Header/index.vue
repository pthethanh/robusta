<template>
  <div class="header">
    <el-menu
      mode="horizontal"
      :router="true"
    >
      <hamburger :toggle-click="toggleSideBar" :is-active="sidebar.opened" class="hamburger-container el-menu-item el-menu-item--horizontal"/>
      <div class="el-menu-item el-menu-item--horizontal nav-title">
        <router-link to="/">
          <Logo></Logo>
        </router-link>
      </div>
      <el-menu-item index="1" route="/home">News</el-menu-item>
      <el-menu-item index="2" route="/tutorials">Tutorials</el-menu-item>
      <el-menu-item index="3" route="/interview/programs">Interview Preparation</el-menu-item>
      <el-menu-item index="4" route="/login" style="float:right">Login</el-menu-item>
    </el-menu>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Hamburger from '@/components/Hamburger'
import Logo from '@/components/Logo'

export default {
  components: {
    Hamburger,
    Logo
  },
  computed: {
    ...mapGetters([
      'sidebar'
    ])
  },
  methods: {
    toggleSideBar () {
      this.$store.dispatch('ToggleSideBar')
    }
  }
}
</script>

<style scoped>
.el-menu-item.el-menu-item--horizontal.nav-title {
  font-weight: 700;
  font-size: 1.5em;
  color: black;
}
.el-menu-item.el-menu-item--horizontal.nav-title:hover {
  color: black;
}
.el-menu-item {
  font-weight: 550;
}
.el-menu {
  overflow: hidden;
  position: fixed;
  top: 0;
  width: 100%;
  z-index: 1;
}
</style>
