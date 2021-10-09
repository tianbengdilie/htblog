import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'
import user from './modules/user'
import settings from './modules/settings'
import permission from './modules/permission'

Vue.use(Vuex)

const vuexLocal = new VuexPersistence({
  storage: window.localStorage,
  modules: ['user']
})

const store = new Vuex.Store({
  modules: {
    user,
    settings,
    permission
  },
  plugins: [vuexLocal.plugin]
})

export default store