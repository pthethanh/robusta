import request from '@/utils/request'

export function listSolutionInfo (query) {
  let queryStr = ''
  if (query !== undefined) {
    queryStr += query + '&q=info'
  }
  return request({
    url: '/api/v1/solutions?' + queryStr,
    method: 'get'
  })
}

export function completionReport (query) {
  let queryStr = ''
  if (query !== undefined) {
    queryStr += query + '&q=completion'
  }
  return request({
    url: '/api/v1/solutions/reports?' + queryStr,
    method: 'get'
  })
}
