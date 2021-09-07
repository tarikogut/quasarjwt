const user = JSON.parse(localStorage.getItem('user'));
const initialState = user
  ? {status: {loggedIn: true}, user}
  : {status: {loggedIn: false}, user: null};
export const login = ({commit, dispatch}, user) => {
  return AuthService.login(user).then(user => {
    commit('loginSucces', user)
    return Promise.resolve(user)
  })
}
