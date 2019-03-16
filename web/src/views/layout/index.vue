<template>
  <div class="home">
    <el-container>
      <el-header>
        <Header></Header>
      </el-header>
      <el-main>
        <router-view/>
      </el-main>
      <el-footer></el-footer>
    </el-container>
  </div>
</template>

<script>
import Header from '@/components/Header'

export default {
  components: {
    Header
  }
}
</script>

<style>
  .home {
    height: 100%;
    width: 100%;
  }
</style>
