import request from '@/utils/request'

export function register(data) {
  return request({
    url: '/api/v1/users/registration',
    method: 'POST',
    data
  })
}

export function listUsers() {
  return request({
    url: '/api/v1/users',
    method: 'GET'
  })
}

export function resetPassword(data) {
  return request({
    url: '/api/v1/users?action=reset-password',
    method: 'PUT',
    data
  })
}

export function genResetPasswordToken(data) {
  return request({
    url: '/api/v1/users?action=request-reset-password',
    method: 'PUT',
    data
  })
}
