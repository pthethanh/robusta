export function getAPI(path) {
  return process.env.NODE_ENV === 'production' ? path : 'http://localhost:8080' + path
}
