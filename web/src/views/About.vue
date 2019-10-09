<template>
  <div class="about">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14" class="content">
        <h2>About</h2>
        <p>Hi there,</p>
        <p>
          <logo :contrast="true" :raw="true"></logo> is a website that I've used for practicing programming languages and hence it is supposed to be changed every day.
        </p>
        <p>If you see any issues with the site, please drop me an email at <a href="mailto:phamthethanh1088@gmail.com?Subject=Goway%3A%20Feedback" target="_top" style="color:#51AEE7">phamthethanh1088@gmail.com</a></p>
        <p>Thank you,</p>
        <p>Thanh</p>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import Logo from '@/components/Logo'

export default {
  components: {
    Logo
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/variables.scss';

.about {
  padding: 20px 20px;
  background-color: $backgroundColorSecondary;
  height: 100vh;

  .content {
    background-color: $backgroundColorInfo;
    padding: 20px 20px;
    font-family: 'Open Sans', sans-serif;
    font-size: 0.95em;
    border-radius: 5px;
  }
}
</style>
