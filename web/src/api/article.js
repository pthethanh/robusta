import request from '@/utils/request'

export function fetchList (query) {
  let queryStr = ''
  if (query !== undefined) {
    queryStr += query + '&strategy=cache'
  }
  return request({
    url: '/api/v1/articles?' + queryStr,
    method: 'get'
  })
}

export function fetchArticle (id) {
  return request({
    url: '/api/v1/articles/' + id,
    method: 'get'
  })
}

export function createArticle (data) {
  return request({
    url: '/api/v1/articles',
    method: 'POST',
    data
  })
}

export function updateArticle (id, data) {
  return request({
    url: '/api/v1/articles/' + id,
    method: 'put',
    data
  })
}

export function updateArticleView (id) {
  return request({
    url: '/api/v1/articles/' + id + '?action=update_view',
    method: 'put'
  })
}

export function deleteArticle (id) {
  return request({
    url: '/api/v1/articles/' + id,
    method: 'delete'
  })
}

export function reactToArticle (id, typ) {
  let data = {
    target_type: 'article',
    target_id: id,
    type: typ
  }
  return request({
    url: '/api/v1/reactions',
    method: 'post',
    data
  })
}
