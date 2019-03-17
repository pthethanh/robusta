import request from '@/utils/request'

export function fetchList (query) {
  return request({
    url: '/api/v1/articles',
    method: 'get',
    params: query
  })
}

export function fetchArticle (id) {
  return request({
    url: '/api/v1/articles',
    method: 'get',
    params: { id }
  })
}

export function createArticle (data) {
  return request({
    url: '/api/v1/articles',
    method: 'post',
    data
  })
}

export function updateArticle (data) {
  return request({
    url: '/api/v1/articles',
    method: 'post',
    data
  })
}
