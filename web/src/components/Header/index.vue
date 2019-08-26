<template>
  <div class="navbar">
    <el-menu mode="horizontal" :router="true" text-color="#383838" active-text-color="#509CFA" background-color="white" :default-active="activeMenuIndex" @select="onSelect">
      <hamburger :toggle-click="toggleSideBar" :is-active="sidebar.opened" class="hamburger-container el-menu-item el-menu-item--horizontal hidden-sm-and-up" />
      <Logo class="logo" @click="onSelect('1', ['1'])"></Logo>
      <el-menu-item class="hidden-xs-only" index="1" route="/home">Blog</el-menu-item>
      <el-menu-item class="hidden-xs-only" index="2" route="/resources">Resources</el-menu-item>
      <el-menu-item class="hidden-xs-only" index="3" route="/about">About</el-menu-item>
      <el-dropdown v-if="user.authenticated" class="right-menu-item el-menu-item el-menu-item--horizontal" trigger="click">
        <avatar size="small" :src="user.info.avatar_url" :names="[user.info.first_name, user.info.email]"></avatar>
        <el-dropdown-menu slot="dropdown">
          <router-link to="/profile">
            <el-dropdown-item>Profile</el-dropdown-item>
          </router-link>
          <el-dropdown-item divided>
            <span style="display:block;" @click="onLogout">Sign Out</span>
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
      <div v-if="!user.authenticated" index="6" class="el-menu-item right-menu-item" @click="onLogin">Sign in</div>
      <el-menu-item index="5" route="/articles/new" class="right-menu-item">New Post</el-menu-item>
    </el-menu>
  </div>
</template>

<script>
import {
  mapGetters
} from 'vuex'
import Hamburger from '@/components/Hamburger'
import Avatar from '@/components/Avatar'
import Logo from '@/components/Logo'
export default {
  components: {
    Hamburger,
    Avatar,
    Logo
  },
  data () {
    return {
      keyword: '',
      activeIndex: '1'
    }
  },
  computed: {
    ...mapGetters([
      'sidebar',
      'user'
    ]),
    activeMenuIndex () {
      return this.activeIndex
    },
    firstLetterName() {
      var s = ''
      if (this.user.info.first_name !== undefined) {
        s = this.user.info.first_name
      } else {
        s = this.user.info.email
      }
      if (s === undefined) {
        return '^_^'
      }
      return s.substr(0, 1)
    }
  },
  methods: {
    toggleSideBar () {
      this.$store.dispatch('ToggleSideBar')
    },
    onLogin () {
      this.$store.dispatch('ToggleLogin', true)
    },
    onLogout () {
      this.$store.dispatch('LogOut')
    },
    onSelect (index, path) {
      this.activeIndex = index
    }
  }
}
</script>

<style lang="scss" scoped>
.navbar {
  height: 48px;
  position: fixed;
  z-index: 1;
  width: 100%;

  .right-menu-item {
    float: right;
    color: #3F3B3B;
  }
  .logo {
    float: left;
    height: 48px;
    line-height: 48px;
    margin: 0;
    border-bottom: 2px solid transparent;
    padding: 0 20px;
    &:hover {
      cursor: pointer;
    }
  }
}
</style>
