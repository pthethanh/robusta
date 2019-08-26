import Cookies from 'js-cookie'

const TokenKey = '_r_token'
const UserKey = '_r_user'

export function getToken () {
  return Cookies.get(TokenKey)
}

export function setToken (token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken () {
  return Cookies.remove(TokenKey, { path: '/', domain: '.' + window.location.hostname })
}

export function getUser() {
  return Cookies.get(UserKey)
}

export function setUser(user) {
  return Cookies.set(UserKey, user)
}

export function removeUser() {
  Cookies.remove(UserKey, { path: '/', domain: '.' + window.location.hostname })
}
