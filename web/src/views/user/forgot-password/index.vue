<template>
  <div class="forgot-password">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="10" :md="10" :lg="8" :xl="8">
        <div class="title">Forgot Password</div>
        <p>Please enter the email that you registered below for resetting password.</p>
        <el-form :model="user" :rules="rules" ref="forgot-password-form" @submit.native.prevent="sendResetPasswordRequest">
          <el-form-item prop="email" :error="error">
            <el-input v-model="user.email" placeholder="Email" type="email">
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" type="primary" @click="sendResetPasswordRequest">Submit</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {
  genResetPasswordToken
} from '@/api/user'
export default {
  data () {
    return {
      user: {
        email: ''
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
        }]
      }
    }
  },
  methods: {
    sendResetPasswordRequest () {
      var isValid = true
      this.$refs['forgot-password-form'].validate((valid) => {
        isValid = valid
      })
      if (!isValid) {
        return
      }

      this.loading = true
      genResetPasswordToken(JSON.stringify(this.user)).then((response) => {
        this.$message({
          type: 'success',
          message: 'Reset Password request has been sent to your email. Please check.'
        })
        this.error = ''
      }).catch((error) => {
        this.loading = false
        var res = error.response
        if (res !== undefined && res !== null) {
          this.error = res.data.message
        } else {
          this.error = 'There were some error during send request for resetting password. Please try again.'
        }
      })
      this.loading = false
    }
  }
}
</script>
