import Vue from 'vue'
import VueRouter from 'vue-router'
import More from '../views/More.vue'
import Grid from '../views/Grid.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Grid',
    component: Grid,
  },
  {
    path: '/more',
    name: 'More',
    component: More,
  },
  {
    path: '/:endpoint',
    name: 'Grid',
    component: Grid,
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
})

export default router
