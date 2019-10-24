<template>
  <div class="navbar">
    <el-menu mode="horizontal" :router="true" :text-color="variables.headerTextColor" :active-text-color="variables.headerActiveTextColor" :background-color="variables.headerBackgroundColor" :default-active="activeMenuIndex" @select="onSelect">
      <hamburger :toggle-click="toggleSideBar" :is-active="sidebar.opened" class="hamburger-container" />
      <logo class="logo" @click="onSelect('1', ['1'])"></logo>
      <el-menu-item class="hidden-xs-only menu-item-bold" index="1" route="/articles">{{ $t('nav.blog') }}</el-menu-item>
      <el-menu-item class="hidden-xs-only menu-item-bold" index="2" route="/challenges">{{ $t('nav.challenges') }}</el-menu-item>
      <el-menu-item class="hidden-xs-only menu-item-bold" index="3" route="/resources">{{ $t('nav.resources') }}</el-menu-item>
      <el-menu-item class="hidden-xs-only menu-item-bold" index="4" route="/about">{{ $t('nav.about') }}</el-menu-item>
      <el-dropdown v-if="user.authenticated" class="right-menu-item el-menu-item el-menu-item--horizontal" trigger="click">
        <avatar size="small" :src="user.info.avatar_url" :names="[user.info.first_name, user.info.email]"></avatar>
        <el-dropdown-menu slot="dropdown">
          <router-link to="/profile">
            <el-dropdown-item>{{ $t('nav.profile') }}</el-dropdown-item>
          </router-link>
          <el-dropdown-item divided>
            <span style="display:block;" @click="onLogout">{{ $t('nav.sign_out') }}</span>
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
      <div v-if="!user.authenticated" index="7" class="el-menu-item right-menu-item menu-item-bold" @click="onLogin">{{ $t('nav.sign_in') }}</div>
      <el-menu-item index="6" route="/articles/new" class="right-menu-item menu-item-bold" style="color: #F2CD02;">{{ $t('nav.new_post') }}</el-menu-item>
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
import variables from '@/styles/variables.scss'
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
    firstLetterName () {
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
    },
    variables () {
      return variables
    }
  },
  methods: {
    toggleSideBar () {
      this.$store.dispatch('ToggleSideBar')
    },
    onLogin() {
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
@import '@/styles/variables.scss';

.navbar {
  height: 48px;
  position: fixed;
  z-index: 999;
  width: 100%;

  .right-menu-item {
    float: right;
    color: $headerTextColor;
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
