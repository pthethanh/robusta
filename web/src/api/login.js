import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/api/v1/auth',
    method: 'POST',
    data
  })
}

export function getInfo (token) {
  return request({
    url: '/user/info',
    method: 'GET',
    params: { token }
  })
}

export function logout (provider) {
  return request({
    url: '/logout/' + provider,
    method: 'POST'
  })
}
