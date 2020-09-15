import Vue from 'vue'
import Vuex from 'vuex'
import { api } from '@/functions/api.js'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    endpoints: [],
  },
  mutations: {
    endpoints (state) {
      api.getEndpoints((res) => {
        state.endpoints = res
      })
    },
  },
  getters: {
    endpoints (state) {
      return state.endpoints
    },
  },
  actions: {
  },
  modules: {
  },
})
