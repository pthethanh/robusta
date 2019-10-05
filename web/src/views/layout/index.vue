<template>
  <div class="layout">
    <el-container class="full-height">
      <el-header style="padding: 0px 0px; height: 48px; line-height: 48px;">
        <Header></Header>
      </el-header>
      <el-container>
        <sidebar></sidebar>
        <el-main style="padding: 0px; height:100%">
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
    <el-dialog :visible.sync="auth.open" modal top="0vh" :width="dialogWidth" :show-close="false">
      <login></login>
    </el-dialog>
  </div>
</template>

<script>
import Header from '@/components/Header'
import Sidebar from '@/components/Sidebar'
import Login from '@/components/Login'
import {
  mapGetters
} from 'vuex'
import ResizeMixin from './mixin/ResizeHandler'
export default {
  components: {
    Header,
    Sidebar,
    Login
  },
  computed: {
    ...mapGetters([
      'auth',
      'device'
    ]),
    dialogWidth () {
      if (this.device === 'mobile') {
        return '100%'
      }
      return '50%'
    }
  },
  mixins: [ResizeMixin]
}
</script>

<style lang="scss" scoped>
.layout {
  height: 100%;
  width: 100%;
  padding: 0px 0px;
  margin: 0px 0px;
}
</style>
