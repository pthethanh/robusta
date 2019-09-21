import axios from 'axios'
import { Message } from 'element-ui'
import store from '../store'
import { getToken } from '@/utils/auth'

const service = axios.create({
  baseURL: process.env.NODE_ENV === 'production' ? '/' : 'http://localhost:8080',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json'
  }
})

service.interceptors.request.use(
  config => {
    if (store.getters.token) {
      config.headers['Authorization'] = getToken()
    }
    return config
  },
  error => {
    // Do something with request error
    // eslint-disable-next-line
    console.log(error)
    // eslint-disable-next-line
    Promise.reject('Error: ' + error)
  }
)

service.interceptors.response.use(
  response => {
    const res = response.data
    if (res.code !== 200 && res.code !== 201) {
      Message({
        message: res.message,
        type: 'error',
        duration: 5 * 1000
      })
      // eslint-disable-next-line
      return Promise.reject('error')
    } else {
      return response.data
    }
  },
  error => {
    var res = error.response
    if (res !== undefined && res !== null) {
      var code = res.data.code
      if (code === 401 && store.getters.token) {
        store.dispatch('ToggleLogin', true)
        return Promise.reject(error)
      }
      Message({
        message: error.message,
        type: 'error',
        duration: 5 * 1000
      })
      // eslint-disable-next-line
      console.log(res)
      return Promise.reject(error)
    }
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
