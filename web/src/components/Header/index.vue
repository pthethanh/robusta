<template>
  <div class="header">
    <el-menu
      mode="horizontal"
      :router="true"
      text-color="black"
      active-text-color="red"
      background-color="white"
    >
      <hamburger :toggle-click="toggleSideBar" :is-active="sidebar.opened" class="hamburger-container el-menu-item el-menu-item--horizontal"/>
      <div class="el-menu-item el-menu-item--horizontal nav-title">
        <router-link to="/">
          <Logo></Logo>
        </router-link>
      </div>
      <!-- temporary hidden the menu items-->
      <el-menu-item index="1" route="/home" hidden>Articles</el-menu-item>
      <el-menu-item index="2" route="/tutorials" hidden>Tutorials</el-menu-item>
      <el-menu-item index="3" route="/interview/programs" hidden>Interview Preparation</el-menu-item>
      <el-input class="el-menu el-menu-item" v-model="keyword" placeholder="tag: go-introduction"></el-input>
      <el-menu-item index="4" route="/login" style="float:right" hidden>Login</el-menu-item>
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
  data(){
    return {
      keyword: ""
    }
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
  color: black;
}
.el-menu {
  overflow: hidden;
  position: fixed;
  top: 0;
  width: 100%;
  z-index: 1;
}
</style>
