<template>
  <div class="register">
    <div class="form">
      <div class="title">Create a new account</div>
      <div>It's quick and easy.</div>
      <el-form class="register-form" :model="user" :rules="rules" ref="register-form" @submit.native.prevent="register">
        <el-form-item>
          <el-col :span="12">
            <el-input v-model="user.first_name" placeholder="First Name"></el-input>
          </el-col>
          <el-col :span="12">
            <el-input v-model="user.last_name" placeholder="Last Name"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item prop="email" :error="error">
          <el-input v-model="user.email" placeholder="Email" type="email">
          </el-input>
        </el-form-item>
        <el-form-item prop="password" :error="error">
          <el-input v-model="user.password" placeholder="Password" type="password">
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button :loading="loading" type="primary" @click="register">Register</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {
  register
} from '@/api/user'
export default {
  data () {
    return {
      user: {
        email: '',
        password: '',
        first_name: '',
        last_name: ''
      },
      loading: false,
      error: '',
      rules: {
        email: [{
          required: true,
          message: 'Email is required',
          trigger: 'blur'
        },
        {
          type: 'email',
          message: 'Please enter a valid email address',
          trigger: 'blur'
        }],
        password: [{
          required: true,
          message: 'Password is required',
          trigger: 'blur'
        },
        {
          min: 5,
          message: 'Password length should be at least 5 characters',
          trigger: 'blur'
        },
        {
          max: 25,
          message: 'Password length should be less than 25 characters',
          trigger: 'blur'
        }]
      }
    }
  },
  methods: {
    register () {
      var isValid = true
      this.$refs['register-form'].validate((valid) => {
        isValid = valid
      })
      if (!isValid) {
        return
      }
      this.loading = true
      var self = this
      register(JSON.stringify(this.user)).then(response => {
        self.$message({
          type: 'success',
          message: 'Account has been created. Go ahead to login!'
        })
        self.error = ''
        self.$router.push('/login?redirect=/')
      }).catch(error => {
        this.loading = false
        console.log(error)
        var res = error.response
        if (res !== undefined && res !== null) {
          this.error = res.data.message
        } else {
          this.error = 'There were some error during the registration. Please try again'
        }
      })
      this.loading = false
    }
  }
}
</script>

<style lang="scss" scoped>
.register {
  height: 100%;

  .form {
    float: right;
    background-color: white;
    padding: 20px 20px;

    .title {
      font-size: 2.5em;
      font-weight: 700;
    }

  }

}
</style>
