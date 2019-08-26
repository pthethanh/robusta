const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  roles: state => state.user.roles,
  user: state => state.user,
  auth: state => state.app.auth,
  authenticated: state => state.user.authenticated
}
export default getters
