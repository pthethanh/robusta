import request from '@/utils/request'

export function fetchTutorials (query) {
  return request({
    url: '/api/v1/tutorials',
    method: 'get',
    params: query
  })
}

export function fetchTutorial (id) {
  return request({
    url: '/api/v1/tutorials/' + id,
    method: 'get'
  })
}

export function createTutorial (data) {
  return request({
    url: '/api/v1/tutorials',
    method: 'POST',
    data
  })
}

export function updateTutorial (id, data) {
  return request({
    url: '/api/v1/tutorials/' + id,
    method: 'put',
    data
  })
}

export function viewTutorial (id) {
  return request({
    url: '/api/v1/tutorials/' + id,
    method: 'post'
  })
}

export function deleteTutorial(id) {
  return request({
    url: '/api/v1/tutorials/' + id,
    method: 'delete'
  })
}
