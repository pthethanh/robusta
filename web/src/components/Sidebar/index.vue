<template>
  <div class="sidebar">
    <el-menu default-active="1" v-bind:class="{
      'el-menu-vertical-demo': !isCollapse,
      'el-menu-vertical-demo hidden-lg-and-down': isCollapse
      }" @select="onSelect" :collapse="isCollapse" :router=true background-color="#545c64" text-color="#fff" active-text-color="#ffd04b">
      <div v-for="(route,index) in routes()" :key="route.path">
        <el-menu-item v-if="visibleChildren(route).length === 0 || visibleChildren(route).length === 1" :index="index + 1 + ''" :route="route.path">
          <i v-bind:class="route.icon"></i>
          <span slot="title">{{route.name}}</span>
        </el-menu-item>
        <el-submenu :index="index+1 + ''" v-if="visibleChildren(route).length > 1">
          <template slot="title">
            <i v-bind:class="route.icon"></i>
            <span slot="title">{{route.name}}</span>
          </template>
          <el-menu-item v-for="(child,childIndex) in visibleChildren(route)" :key="childIndex" :index="index + '-' + childIndex" :route="route.path + '/' + child.path">
            {{child.name}}
          </el-menu-item>
        </el-submenu>
      </div>
      <el-menu-item :index="(routes().length + 1) + ''" @click="onLogin()" v-if="!user.authenticated">
        <i class="el-icon-s-custom"></i>
        <span slot="title">Sign in</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<style lang="scss" scoped>
.sidebar {
  z-index: 9999;
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
import {
  routes
} from '@/router'
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
      var requiredRoles = route.roles
      if (requiredRoles === undefined) {
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
    }
  }
}
</script>
