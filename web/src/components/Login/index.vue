<template>
  <div class="login">
    <el-card class="card">
      <logo size="large" :contrast="true"></logo>
      <div class="info">
        {{ $t('login.info') }}
      </div>
      <el-form class="login-form" :model="user" :rules="rules" ref="form" @submit.native.prevent="login">
        <el-form-item prop="username" :error="error">
          <el-input v-model="user.username" :placeholder="$t('gen.email')" prefix-icon="fas fa-user">
          </el-input>
        </el-form-item>
        <el-form-item prop="password" :error="error">
          <el-input v-model="user.password" :placeholder="$t('gen.password')" type="password" prefix-icon="fas fa-lock">
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button :loading="loading" class="login-button" type="primary" native-type="submit" block>{{ $t('nav.sign_in') }}</el-button>
        </el-form-item>
        <el-button :loading="loading" class="login-button social-login-btn" type="plain" @click="loginGoogle">
          <svg-icon icon-class="google" class-name="svg-icon" />
          <span>{{ $t('login.google') }}</span>
        </el-button>
        <div class="footer">
          <div @click="goTo('/users/register')">{{ $t('login.register') }}</div>
          <div @click="goTo('/users/forgot-password')">{{ $t('login.forgot_password') }}</div>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import {
  login
} from '@/api/login'
import Logo from '@/components/Logo'

export default {
  name: 'login',
  components: {
    Logo
  },
  data () {
    return {
      user: {
        username: '',
        password: ''
      },
      loading: false,
      error: '',
      rules: {
        username: [{
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
        }]
      }
    }
  },
  methods: {
    loginLocal () {
      return new Promise((resolve, reject) => {
        var data = {
          provider: 'local',
          username: this.user.username,
          password: this.user.password
        }
        login(JSON.stringify(data)).then(response => {
          this.$store.dispatch('ToggleLogin', false)
          const data = response.data
          this.$store.dispatch('SetToken', data.token)
          this.$store.dispatch('SetInfo', data.user_info)
          this.$store.dispatch('GetInfo')
          this.$router.push(this.getRedirect())
          resolve()
        }).catch(error => {
          if (error.code) {
            this.error = this.$i18n.t('server.' + error.code)
          } else {
            this.error = this.$i18n.t('login.login_failed')
          }
          this.loading = false
          reject(error)
        })
      })
    },
    async login () {
      let valid = await this.$refs.form.validate()
      if (!valid) {
        return
      }
      this.loading = true
      await this.loginLocal()
      this.loading = false
    },
    loginGoogle () {
      this.loading = true
      this.oauthLogin('google')
      this.loading = false
    },
    oauthLogin (provider) {
      let target = process.env.NODE_ENV === 'production' ? '/auth/' + provider : 'http://mylocalhost.com/auth/' + provider
      target = target + '?redirect=' + this.getRedirect()
      window.location.href = target
      this.$store.dispatch('ToggleLogin', false)
    },
    getRedirect () {
      var redirect = this.$router.currentRoute.query.redirect
      if (redirect === undefined) {
        redirect = this.$router.currentRoute.path
      }
      if (redirect === '/users/register') {
        redirect = '/'
      }
      return redirect
    },
    goTo (p) {
      this.$store.dispatch('ToggleLogin', false)
      this.$router.push(p)
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/variables.scss';

.login {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: center;
  text-align: center;

  .card {
    border: 1px solid lightgrey;

    .info {
      max-width: 290px;
      word-break: keep-all;
      font-size: 0.95em;
      line-height: 1.4em;
      margin-bottom: 10px;
    }

    .login-form {
      min-width: 290px;

      .login-button {
        width: 100%;
        font-weight: $btnFontWeight;
      }

      .social-login-btn {
        border: 1px solid lightgrey;
        align-items: center;
        display: flex;
      }

      .svg-icon {
        font-size: 1.3em;
        vertical-align: text-bottom;
      }

      .footer {
        font-size: 0.95em;
        text-decoration: underline;
      }
    }

    @media only screen and (max-width: 600px) {
      .login-form {
        min-width: 240px;
      }
    }
  }
}
</style>
