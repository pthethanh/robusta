<template>
  <div class="header">
    <el-menu
      mode="horizontal"
      :router="true"
    >
      <hamburger :toggle-click="toggleSideBar" :is-active="sidebar.opened" class="hamburger-container el-menu-item el-menu-item--horizontal"/>
      <el-menu-item index="1" route="/home">Home</el-menu-item>
      <el-menu-item index="2" route="/tutorials">Tutorials</el-menu-item>
      <el-menu-item index="3" route="/users/login">Login</el-menu-item>

    </el-menu>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Hamburger from '@/components/Hamburger'
export default {
  components: {
    Hamburger
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
