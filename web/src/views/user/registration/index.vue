<template>
  <div class="register">
    <div class="form">
      <div class="title">{{ $t('user.registration.title') }}</div>
      <div>{{ $t('user.registration.info') }}</div>
      <el-form class="register-form" :model="user" :rules="rules" ref="register-form" @submit.native.prevent="register">
        <el-form-item>
          <el-col :span="12">
            <el-input v-model="user.first_name" :placeholder="$t('user.first_name')"></el-input>
          </el-col>
          <el-col :span="12">
            <el-input v-model="user.last_name" :placeholder="$t('user.last_name')"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item prop="email" :error="error">
          <el-input v-model="user.email" :placeholder="$t('user.email')" type="email">
          </el-input>
        </el-form-item>
        <el-form-item prop="password" :error="error">
          <el-input v-model="user.password" :placeholder="$t('user.password')" type="password">
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button :loading="loading" type="primary" @click="register">{{ $t('user.registration.register') }}</el-button>
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
          message: this.$i18n.t('validation.email_required'),
          trigger: 'blur'
        },
        {
          type: 'email',
          message: this.$i18n.t('validation.email_invalid'),
          trigger: 'blur'
        }
        ],
        password: [{
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
          message: this.$i18n.t('user.registration.success_message')
        })
        self.error = ''
        self.$router.push('/login?redirect=/')
      }).catch((error) => {
        this.loading = false
        this.error = this.$i18n.t('user.registration.failed_message') + ': ' + error
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
