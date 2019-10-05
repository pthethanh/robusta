<template>
  <div class="login">
    <Login></Login>
  </div>
</template>

<script>
import Login from '@/components/Login'
export default {
  components: {
    Login
  },
  title: 'Goway - Login',
  data () {
    return {
      user: {
        username: '',
        password: ''
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/variables.scss';

.login {
  height: 100%;
  vertical-align: middle;
  background-color: $backgroundColorPrimary;
}
</style>
