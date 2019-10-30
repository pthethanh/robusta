import axios from 'axios'
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
    if (res.code !== 1000) {
      // eslint-disable-next-line
      return Promise.reject(res.message)
    } else {
      return response.data
    }
  },
  error => {
    console.log(error)
    var res = error.response
    if (res !== undefined && res !== null) {
      var code = res.data.code
      if (code === 1100 && (store.getters.token === '' || store.getters.token === undefined)) {
        store.dispatch('ToggleLogin', true)
        return Promise.reject(res.data)
      }
      return Promise.reject(res.data)
    }
    let data = { code: 1004, status: 500 }
    return Promise.reject(data)
  }
)

export default service
