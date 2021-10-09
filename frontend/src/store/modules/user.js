import { login, getUserInfo } from '@/api/user'
import { setToken } from '@/utils/auth'
import router from '@/router'

const user = {
  namespaced: true,
  state: {
    userInfo: {
      nickName: '',
    },
    token: '',
  },
  mutations: {
    setUserInfo(state, userInfo) {
      // 这里的 `state` 对象是模块的局部状态
      state.userInfo = userInfo
    },
    setToken(state, token) {
      // 这里的 `state` 对象是模块的局部状态
      state.token = token
    },
    LoginOut(state) {
      state.userInfo = {}
      state.token = ''
      sessionStorage.clear()
      router.push({ name: 'Login', replace: true })
      window.location.reload()
    },
    ResetUserInfo(state, userInfo = {}) {
      state.userInfo = {
        ...state.userInfo,
        ...userInfo
      }
    }
  },
  actions: {
    async GetUserInfo({ commit }) {
      const res = await getUserInfo()
      if (res.code === 0) {
        commit('setUserInfo', res.data.userInfo)
      }
      return res
    },
    async LoginIn({ commit }, loginInfo) {
      const res = await login(loginInfo)
      console.log("response: ")
      console.log(res)
      if (res.data.code === 0) {
        let data = res.data.data
        commit('setUserInfo', data.user)
        commit('setToken', data.token)
        setToken(data.token)
        router.push({ path: '/' })
        return true
      }
    },
    // async LoginIn({ commit, dispatch, rootGetters, getters }, loginInfo) {
    //   const res = await login(loginInfo)
    //   console.log("response: ")
    //   console.log(res)
    //   if (res.data.code === 0) {
    //     commit('setUserInfo', res.data.user)
    //     commit('setToken', res.data.token)
    //     await dispatch('router/SetAsyncRouter', {}, { root: true })
    //     const asyncRouters = rootGetters['router/asyncRouters']
    //     asyncRouters.forEach(asyncRouter => {
    //       router.addRoute(asyncRouter)
    //     })
    //     router.push({ name: getters['userInfo'].authority.defaultRouter })
    //     return true
    //   }
    // },
    async LoginOut({ commit }) {
      commit('LoginOut')
    }
  },
  getters: {
    userInfo(state) {
      return state.userInfo
    },
    token(state) {
      return state.token
    }
  }
}

export default user
