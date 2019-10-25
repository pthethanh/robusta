<template>
  <el-drawer :visible.sync="sidebar.opened" @close="closeSideBar" direction="ltr" :size="width" :show-close="false" class="sidebar">
    <template v-slot:title>
      <div class="title">
        <hamburger :raw="true" :toggle-click="toggleSideBar" :is-active="true" class="title-item"></hamburger>
        <logo :raw="true" class="title-item"></logo>
      </div>
    </template>
    <el-menu default-active="1" @select="onSelect" :router=true :background-color="variables.sideBarBackgroundColor" :text-color="variables.sideBarTextColor" :active-text-color="variables.sideBarActiveTextColor">
      <div v-for="(route,index) in routes()" :key="route.path">
        <el-menu-item v-if="visibleChildren(route).length === 0 || visibleChildren(route).length === 1" :index="index + 1 + ''" :route="route.path">
          <i v-bind:class="route.meta.icon"></i>
          <span slot="title">{{route.meta.name}}</span>
        </el-menu-item>
        <el-submenu :index="index+1 + ''" v-if="visibleChildren(route).length > 1">
          <template slot="title">
            <i v-bind:class="route.meta.icon"></i>
            <span slot="title">{{route.meta.name}}</span>
          </template>
          <el-menu-item v-for="(child,childIndex) in visibleChildren(route)" :key="childIndex" :index="index + '-' + childIndex" :route="route.path + '/' + child.path">
            {{child.meta.name}}
          </el-menu-item>
        </el-submenu>
      </div>
      <el-menu-item :index="(routes().length + 1) + ''" @click="onLogin()" v-if="!user.authenticated">
        <i class="el-icon-s-custom"></i>
        <span slot="title">{{ $t('nav.sign_in')}}</span>
      </el-menu-item>
    </el-menu>
  </el-drawer>
</template>

<script>
import {
  mapGetters
} from 'vuex'
import {
  routes
} from '@/router'
import variables from '@/styles/variables.scss'
import Hamburger from '@/components/Hamburger'
import Logo from '@/components/Logo'

export default {
  components: {
    Hamburger,
    Logo
  },
  computed: {
    ...mapGetters([
      'sidebar',
      'user',
      'device'
    ]),
    variables () {
      return variables
    },
    width () {
      if (this.device === 'mobile') {
        return '100%'
      }
      return '20%'
    }
  },
  methods: {
    onSelect (index, indexPath) {
      this.$store.dispatch('ToggleSideBar')
    },
    onLogin () {
      this.$store.dispatch('ToggleLogin', true)
    },
    routes () {
      var visibleRoutes = []
      for (var i = 0; i < routes.length; i++) {
        if (this.isVisible(routes[i])) {
          visibleRoutes.push(routes[i])
        }
      }
      return visibleRoutes
    },
    visibleChildren (route) {
      if (route.children === undefined) {
        return []
      }
      var result = []
      for (var i = 0; i < route.children.length; i++) {
        if (this.isVisible(route.children[i])) {
          result.push(route.children[i])
        }
      }
      return result
    },
    isVisible (route) {
      if (route.hidden) {
        return false
      }
      var requiredRoles = []
      if (route.meta !== undefined) {
        requiredRoles = route.meta.roles
      }
      if (requiredRoles === undefined || requiredRoles.length === 0) {
        return true
      }
      var currentUser = this.user
      for (var j = 0; j < requiredRoles.length; j++) {
        for (var k = 0; k < currentUser.roles.length; k++) {
          if (requiredRoles[j] === currentUser.roles[k]) {
            return true
          }
        }
      }
    },
    toggleSideBar () {
      this.$store.dispatch('ToggleSideBar')
    },
    closeSideBar() {
      this.$store.dispatch('CloseSideBar')
    }
  }
}
</script>

<style lang="scss">
@import '@/styles/variables.scss';

.sidebar .el-drawer__open .el-drawer.ltr {
  background-color: $sideBarBackgroundColor;
}

.sidebar .el-drawer__header {
  height: 48px;
}
</style>

<style lang="scss" scoped>
@import '@/styles/variables.scss';

.sidebar {
  z-index: 9999;
  position: fixed;
  height: 100%;

  .el-drawer {
    .el-menu {
      height: 100vh;
    }
  }

  .title {
    margin-top: 20px;
    left: 20px;
    font-size: 1.2em;

    .title-item {
      margin-right: 20px;
      vertical-align: middle;
    }
  }
}
</style>
