<template>
  <div class="forgot-password">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="10" :md="10" :lg="8" :xl="8">
        <div class="title">{{ $t('user.forgot_password.title') }}</div>
        <p>{{ $t('user.forgot_password.info') }}</p>
        <el-form :model="user" :rules="rules" ref="forgot-password-form" @submit.native.prevent="sendResetPasswordRequest">
          <el-form-item prop="email" :error="error">
            <el-input v-model="user.email" :placeholder="$t('user.email')" type="email">
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" type="primary" @click="sendResetPasswordRequest">{{ $t('user.forgot_password.submit') }}</el-button>
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
          message: this.$i18n.t('validation.email_required'),
          trigger: 'blur'
        },
        {
          type: 'email',
          message: this.$i18n.t('validation.email_invalid'),
          trigger: 'blur'
        }
        ]
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
          message: this.$i18n.t('user.forgot_password.success_message')
        })
        this.error = ''
      }).catch((error) => {
        this.loading = false
        this.error = this.$i18n.t('user.forgot_password.failed_message') + ':' + error
      })
      this.loading = false
    }
  }
}
</script>
