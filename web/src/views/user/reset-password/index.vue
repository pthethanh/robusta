<template>
  <div class="reset-password">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="10" :md="10" :lg="8" :xl="8">
        <div class="title">{{ $t('user.reset_password.title') }}</div>
        <el-form :model="user" :rules="rules" ref="reset-password-form" @submit.native.prevent="reset">
          <el-form-item prop="new_password" :error="error">
            <el-input v-model="user.new_password" :placeholder="$t('user.password')" type="password">
            </el-input>
          </el-form-item>
          <el-form-item prop="re_type_new_password" :error="error">
            <el-input v-model="user.re_type_new_password" :placeholder="$t('user.password_confirm')" type="password">
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" type="primary" @click="reset">{{ $t('user.reset_password.reset') }}</el-button>
          </el-form-item>
        </el-form>
        <div v-html="$t('user.reset_password.link_to_forgot_password')"></div>
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
          message: this.$i18n.t('validation.password_required'),
          trigger: 'blur'
        },
        {
          min: 5,
          message: this.$i18n.t('validation.password_len_min'),
          trigger: 'blur'
        },
        {
          max: 25,
          message: this.$i18n.t('validation.password_len_max'),
          trigger: 'blur'
        }
        ],
        re_type_new_password: [{
          required: true,
          message: this.$i18n.t('validation.password_required'),
          trigger: 'blur'
        },
        {
          min: 5,
          message: this.$i18n.t('validation.password_len_min'),
          trigger: 'blur'
        },
        {
          max: 25,
          message: this.$i18n.t('validation.password_len_max'),
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
        this.error = this.$i18n.t('user.password_not_match')
        return
      }
      var token = this.$route.params.token
      if (token === undefined || token === '') {
        this.error = this.$i18n.t('user.reset_password.invalid_token')
        return
      }
      this.user.token = token
      this.loading = true
      resetPassword(JSON.stringify(this.user)).then((response) => {
        this.$message({
          type: 'success',
          message: this.$i18n.t('user.reset_password.success_message')
        })
        this.error = ''
      }).catch((error) => {
        this.loading = false
        this.error = this.$i18n.t('user.reset_password.failed_message') + ': ' + error
      })
      this.loading = false
    }
  }
}
</script>
