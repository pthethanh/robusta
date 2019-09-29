import request from '@/utils/request'

export function listChallenges (query) {
  let queryStr = ''
  if (query !== undefined) {
    queryStr += query
  }
  return request({
    url: '/api/v1/challenges?' + queryStr,
    method: 'get'
  })
}

export function getChallenge (id) {
  return request({
    url: '/api/v1/challenges/' + id,
    method: 'get'
  })
}

export function createChallenge (data) {
  return request({
    url: '/api/v1/challenges',
    method: 'POST',
    data
  })
}

export function updateChallenge (id, data) {
  return request({
    url: '/api/v1/challenges/' + id,
    method: 'PUT',
    data
  })
}

export function deleteChallenge (id) {
  return request({
    url: '/api/v1/challenges/' + id,
    method: 'delete'
  })
}
