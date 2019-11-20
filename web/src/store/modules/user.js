import {
  logout
} from '@/api/login'
import {
  getToken,
  removeToken,
  getUser,
  removeUser,
  setToken,
  setUser
} from '@/utils/auth'
import router from '../../router'

const user = {
  state: {
    token: getToken(),
    roles: [],
    info: {},
    authenticated: false
  },

  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    SET_INFO: (state, info) => {
      state.info = info
    },
    SET_AUTHENTICATED: (state, auth) => {
      state.authenticated = auth
    }
  },

  actions: {
    GetInfo ({ commit, state }) {
      var token = getToken()
      if (token === undefined) {
        return
      }
      commit('SET_TOKEN', token)
      var userJSON = atob(getUser())
      var user = JSON.parse(userJSON)
      if (user === undefined) {
        return
      }
      commit('SET_INFO', user)
      if (user.roles && user.roles.length > 0) {
        commit('SET_ROLES', user.roles)
      }
      commit('SET_AUTHENTICATED', true)
    },
    SetToken({ commit, state }, token) {
      commit('SET_TOKEN', token)
      setToken(token)
    },
    SetInfo({ commit, state }, info) {
      commit('SET_INFO', info)
      setUser(info)
    },
    LogOut ({ commit, state }) {
      return new Promise((resolve, reject) => {
        logout(state.info.provider).then(() => {
          commit('SET_TOKEN', '')
          commit('SET_ROLES', [])
          commit('SET_AUTHENTICATED', false)
          commit('SET_INFO', {})
          removeToken()
          removeUser()
          router.push('/')
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default user
