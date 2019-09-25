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
