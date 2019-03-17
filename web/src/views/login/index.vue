<template>
  <div class="login">
    <el-row type="flex" justify="center">
      <el-col :span=8>
        <div>
          <h1>Join <Logo style="display:inline"></Logo></h1>
        </div>
        <el-form>
          <el-form-item label="Username">
            <el-input v-model="user.username"></el-input>
          </el-form-item>
          <el-form-item label="Password">
            <el-input v-model="user.password" type="password"></el-input>
          </el-form-item>
          <el-form-item>
              <el-button type="primary">Login</el-button>
              <div>
                  Not a user? Click
                  <a><router-link to="/users/register"><strong>here to register</strong></router-link></a>
              </div>
              <div>
                  <a><router-link to="/">Return <strong>Home</strong></router-link></a>
              </div>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import Logo from '@/components/Logo/index'
export default {
  components: {
    Logo
  },
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

<style scoped>
  .el-form {
    border: 1px solid lightgrey;
    padding: 10px;
  }
  .login {
    height: 100%;
    background-color: white;
  }
</style>
