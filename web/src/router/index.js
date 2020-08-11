import Vue from 'vue'
import VueRouter from 'vue-router'
import More from '../views/More.vue'
import Grid from '../views/Grid.vue'
import Docs from '../views/Docs.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Grid,
  },
  {
    path: '/more',
    name: 'More',
    component: More,
  },
  {
    path: '/docs',
    name: 'Docs',
    component: Docs,
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
