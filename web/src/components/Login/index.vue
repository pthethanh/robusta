<template>
  <div class="login">
    <el-card class="card">
      <Logo size="large"></Logo>
      <div class="info">
        Sign in to get personalized story recommendations, follow authors and topics you love.
      </div>
      <el-form class="login-form" :model="user" :rules="rules" ref="form" @submit.native.prevent="login">
        <el-form-item prop="username" :error="error">
          <el-input v-model="user.username" placeholder="Email" prefix-icon="fas fa-user">
          </el-input>
        </el-form-item>
        <el-form-item prop="password" :error="error">
          <el-input v-model="user.password" placeholder="Password" type="password" prefix-icon="fas fa-lock">
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button :loading="loading" class="login-button" type="primary" native-type="submit" block>Sign in</el-button>
        </el-form-item>
        <el-button :loading="loading" class="login-button social-login-btn" type="plain" @click="loginGoogle">
          <svg-icon icon-class="google" class-name="svg-icon" />
          <span>Sign in with Google</span>
        </el-button>
        <div class="footer">
          <div @click="goTo('/users/register')">Don't have an account? Register here!</div>
          <div @click="goTo('/users/forgot-password')">Forgot password ?</div>
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
        username: [
          {
            required: true,
            message: 'Email is required',
            trigger: 'blur'
          },
          {
            min: 4,
            message: 'Email length should be at least 4 characters',
            trigger: 'blur'
          }
        ],
        password: [
          {
            required: true,
            message: 'Password is required',
            trigger: 'blur'
          },
          {
            min: 3,
            message: 'Password length should be at least 3 characters',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  methods: {
    loginLocal () {
      var self = this
      return new Promise((resolve, reject) => {
        var data = {
          provider: 'local',
          username: this.user.username,
          password: this.user.password
        }
        login(JSON.stringify(data)).then(response => {
          self.$store.dispatch('ToggleLogin', false)
          const data = response.data
          self.$store.dispatch('SetToken', data.token)
          self.$store.dispatch('GetInfo')
          self.$router.push(this.getRedirect())
          resolve()
        }).catch(error => {
          this.loading = false
          this.error = 'Invalid email or password'
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
    oauthLogin(provider) {
      let target = process.env.NODE_ENV === 'production' ? '/auth/' + provider : 'http://mylocalhost.com/auth/' + provider
      target = target + '?redirect=' + this.getRedirect()
      window.location.href = target
      this.$store.dispatch('ToggleLogin', false)
    },
    getRedirect() {
      var redirect = this.$router.currentRoute.query.redirect
      if (redirect === undefined) {
        redirect = this.$router.currentRoute.path
      }
      if (redirect === '/users/register') {
        redirect = '/'
      }
      return redirect
    },
    goTo(p) {
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
