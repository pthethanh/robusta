<template>
    <el-row type="flex" justify="center" align="middle">
        <el-col :span=24>
            <el-menu
                active-text-color="red"
                background-color="white"
                text-color="black"
                mode="horizontal"
                :router=true
                >
                <el-menu-item index="1" route="/">Home</el-menu-item>
                <el-menu-item index="2" route="/news">News</el-menu-item>
                <el-menu-item index="3" route="/jobs">Jobs</el-menu-item>
                <el-menu-item index="4" route="/asks">Asks</el-menu-item>
                <el-menu-item index="5" route="/login" class="right-menu">Login</el-menu-item>
                <el-menu-item index="6" route="/login" class="right-menu">
                    <el-input v-model="search" class="hidden-xs-only"></el-input>
                </el-menu-item>
            </el-menu>
        </el-col>
    </el-row>
</template>

<script>
export default {
  data () {
    return {
      search: ''
    }
  }
}
</script>

<style scoped>
@import '~element-ui/lib/theme-chalk/display.css';
.el-menu--horizontal > .el-menu-item {
  height: 45px;
  line-height: 45px;
  font-weight: bold;
}
.el-menu {
  position: fixed;
}
.el-menu--horizontal > .right-menu{
  float: right;
}
</style>
