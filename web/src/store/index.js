import Vue from 'vue'
import Vuex from 'vuex'
// import db from 'localforage'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    slide: 0,
  },

  mutations: {
    set (state, payload) {
      Object.assign(state, payload)
    },

    setSlide (state, slide) {
      Object.assign(state, { slide })
    },
  },

  actions: {
  },
})
