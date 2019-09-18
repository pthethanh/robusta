import request from '@/utils/request'

export function listFolders (query) {
  let queryStr = ''
  if (query !== undefined) {
    queryStr += query
  }
  return request({
    url: '/api/v1/folders?' + queryStr,
    method: 'get'
  })
}

export function getFolder (id) {
  return request({
    url: '/api/v1/folders/' + id,
    method: 'get'
  })
}

export function createFolder (data) {
  return request({
    url: '/api/v1/folders',
    method: 'POST',
    data
  })
}
