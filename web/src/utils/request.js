import axios from 'axios'
import { Message, MessageBox } from 'element-ui'
import store from '../store'
import { getToken } from '@/utils/auth'

const service = axios.create({
  baseURL: process.env.BASE_API || 'http://localhost:8080',
  timeout: 5000
})

service.interceptors.request.use(
  config => {
    if (store.getters.token) {
      config.headers['X-Token'] = getToken() // Let each request carry a custom token, please modify it according to the actual situation.
    }
    return config
  },
  error => {
    // Do something with request error
    console.log(error) // for debug
    Promise.reject('Error: ' + error)
  }
)

service.interceptors.response.use(
  response => {
    /**
     * Code is non-20000 is a mistake can be combined with their own business to modify
     */
    const res = response.data
    if (res.code !== 20000) {
      Message({
        message: res.message,
        type: 'error',
        duration: 5 * 1000
      })

      // 50008: Illegal token; 50012: Other clients logged in; 50014: Token expired;
      if (res.code === 50008 || res.code === 50012 || res.code === 50014) {
        MessageBox.confirm(
          'You have been logged out, you can cancel to stay on this page, or log in again',
          'Determined to log out',
          {
            confirmButtonText: 'Re-login',
            cancelButtonText: 'Cancel',
            type: 'warning'
          }
        ).then(() => {
          store.dispatch('FedLogOut').then(() => {
            location.reload() // To re-instantiate the vue-router object to avoid bugs
          })
        })
      }
      return Promise.reject('error')
    } else {
      return response.data
    }
  },
  error => {
    console.log('err' + error) // for debug
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
