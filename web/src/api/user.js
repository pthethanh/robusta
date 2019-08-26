import request from '@/utils/request'

export function register(data) {
  return request({
    url: '/api/v1/users/registration',
    method: 'POST',
    data
  })
}
