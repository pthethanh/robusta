<template>
  <div class="reset-password">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="10" :md="10" :lg="8" :xl="8">
        <div class="title">Reset Password</div>
        <el-form :model="user" :rules="rules" ref="reset-password-form" @submit.native.prevent="reset">
          <el-form-item prop="new_password" :error="error">
            <el-input v-model="user.new_password" placeholder="Password" type="password">
            </el-input>
          </el-form-item>
          <el-form-item prop="re_type_new_password" :error="error">
            <el-input v-model="user.re_type_new_password" placeholder="Re-type Password" type="password">
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" type="primary" @click="reset">Reset</el-button>
          </el-form-item>
        </el-form>
        <div>
            Click <a href="/users/forgot-password" style="color:red">here</a> to request a new reset password link if the current link doesn't work or already expired.
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {
  resetPassword
} from '@/api/user'
export default {
  data () {
    return {
      user: {
        token: '',
        new_password: '',
        re_type_new_password: ''
      },
      loading: false,
      error: '',
      rules: {
        new_password: [{
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
          }
        ],
        re_type_new_password: [{
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
          }
        ]
      }
    }
  },
  methods: {
    reset () {
      var isValid = true
      this.$refs['reset-password-form'].validate((valid) => {
        isValid = valid
      })
      if (!isValid) {
        return
      }
      if (this.user.new_password !== this.user.re_type_new_password) {
        this.error = 'Re-type Password does not match'
        return
      }
      var token = this.$route.params.token
      if (token === undefined || token === '') {
        this.error = 'Invalid token'
        return
      }
      this.user.token = token
      this.loading = true
      resetPassword(JSON.stringify(this.user)).then((response) => {
        this.$message({
          type: 'success',
          message: 'Password has been reset succesfully. Please try to login again.'
        })
        this.error = ''
      }).catch((error) => {
        this.loading = false
        var res = error.response
        if (res !== undefined && res !== null) {
          this.error = res.data.message
        } else {
          this.error = 'There were some error during resetting password. Please try again.'
        }
      })
      this.loading = false
    }
  }
}
</script>
