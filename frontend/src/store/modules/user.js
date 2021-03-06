import { loginByEmail, getUserInfo } from '@/api/mock_user'

const user = {
  namespaced: true,
  state: {
    token: '',
    roles: [],
    user: '',
    name: '',
    avatar: '',
    status: '',
    code: '',
    introduction: '',
    setting: {
      articlePlatform: [],
    },
  },
  getters: {
    token: (state) => state.token,
    roles: (state) => state.roles,
    name: (state) => state.name,
    user: (state) => state.user,
    avatar: (state) => state.avatar,
    status: (state) => state.status,
    introduction: (state) => state.introduction,
    setting: (state) => state.setting,
  },
  mutations: {
    SET_USER_INFO: (state, payload) => {
      if (payload.logout) {
        state.token = '';
        state.roles = [];
        state.user = '';
        state.name = '';
        state.avatar = '';
        state.code = '';
      } else {
        state.token = payload.token || state.token;
        state.roles = payload.roles || state.roles;
        state.user = payload.user || state.user;
        state.name = payload.name || state.name;
        state.avatar = payload.avatar || state.avatar;
        state.code = payload.code || state.code;
      }
    },
    SET_TOKEN: (state, token) => {
      console.log("SET_TOKEN:", token);
      state.token = token;
    },
  }, actions: {
    // Login user
    LoginByEmail: async ({ commit, dispatch }, payload) => {
      try {
        const response = await loginByEmail(payload.email.trim(), payload.password);
        console.log('[vuex.user] LoginByEmail', payload, response);
        await commit('SET_TOKEN', response.user.token);
        await commit('SET_USER_INFO', response.user);
        await dispatch('permission/GenerateRoutes', response.user, { root: true });
      } catch (err) {
        console.warn('[vuex.user] LoginByEmail', err);
      }
    },

    GetUserInfo: async ({ commit, state }) => {
      console.log('[vuex.user] GetUserInfo');
      try {
        const response = await getUserInfo(state.token);

        // Since mockjs does not support custom status codes, it can only be hacked like this
        if (!response) {
          throw new Error('Verification failed, please login again.');
        }

        // Verify returned roles are a non-null array
        if (response.user.roles && response.user.roles.length === 0) {
          throw new Error('getInfo: roles must be a non-null array!');
        }

        commit('SET_USER_INFO', response.user);
      } catch (err) {
        console.warn('[vuex.user] GetUserInfo', err);
      }
    },

    LogOut: async ({ commit }) => {
      try {
        console.log('[vuex.user] LogOut');
        await commit('SET_USER_INFO', { logout: true });
      } catch (err) {
        console.warn('[vuex.user] LogOut', err);
      }
    },
  },
};

export default user
