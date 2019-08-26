import request from '@/utils/request'

export function findComments (query) {
  let queryStr = ''
  if (query !== undefined) {
    queryStr += query
  }
  return request({
    url: '/api/v1/comments?' + queryStr,
    method: 'get'
  })
}

export function createComment (data) {
  return request({
    url: '/api/v1/comments',
    method: 'POST',
    data
  })
}

export function updateComment (id, data) {
  return request({
    url: '/api/v1/comments/' + id,
    method: 'put',
    data
  })
}

export function deleteComment (id) {
  return request({
    url: '/api/v1/comments/' + id,
    method: 'delete'
  })
}

export function reactToComment (id, typ) {
  let data = {
    type: typ,
    target_type: 'comment',
    target_id: id
  }
  return request({
    url: '/api/v1/reactions',
    method: 'post',
    data
  })
}
