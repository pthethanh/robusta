<template>
  <div class="layout">
    <el-container>
      <el-header style="padding: 0px 0px">
        <Header></Header>
      </el-header>
      <el-container>
        <el-aside width="200px" v-show="isCollapse">
          <Sidebar></Sidebar>
        </el-aside>
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import Header from '@/components/Header'
import Sidebar from '@/components/Sidebar'
import { mapGetters } from 'vuex'

export default {
  components: {
    Header,
    Sidebar
  },
  computed: {
    ...mapGetters([
      'sidebar'
    ]),
    isCollapse () {
      return !this.sidebar.opened
    }
  }
}
</script>

<style>
  .layout {
    height: 100%;
    width: 100%;
    padding: 0px 0px;
    margin: 0px 0px;
  }
</style>

<style>
  body {
    background-color: #D9DEE6;
  }
</style>
