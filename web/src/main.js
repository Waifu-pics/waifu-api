import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import VueNotification from "@kugatsu/vuenotification"

Vue.config.productionTip = false

Vue.use(VueNotification, {
  timer: 4,
  showLeftIcn: false,
})

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App),
}).$mount('#app')
