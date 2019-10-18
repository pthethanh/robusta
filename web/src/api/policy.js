import request from '@/utils/request'

export function addPolicy(data) {
  return request({
    url: '/api/v1/policies?action=add-policy',
    method: 'POST',
    data
  })
}

export function addGroupPolicy(data) {
  return request({
    url: '/api/v1/policies?action=add-group-policy',
    method: 'POST',
    data
  })
}

export function listPolicyActions() {
  return request({
    url: '/api/v1/policies/actions',
    method: 'GET'
  })
}

export function listPolicy(q) {
  return request({
    url: '/api/v1/policies?' + q,
    method: 'GET'
  })
}

export function removePolicy(data) {
  return request({
    url: '/api/v1/policies',
    method: 'DELETE',
    data
  })
}
