import request from '@/utils/request'

export function runCode (data) {
  return request({
    url: '/api/v1/playground/run',
    method: 'POST',
    data
  })
}

export function evaluate (data) {
  return request({
    url: '/api/v1/playground/evaluate',
    method: 'POST',
    data
  })
}
