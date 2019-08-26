import request from '@/utils/request'

export function uploadImageByFile(data) {
  return request({
    url: '/api/v1/editor/image_by_file',
    method: 'POST',
    headers: { 'Content-Type': 'multipart/form-data' },
    data
  })
}

export function uploadImageByURL(data) {
  return request({
    url: '/api/v1/editor/image_by_url',
    method: 'POST',
    data
  })
}
