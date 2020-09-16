import Vue from 'vue'
import VueRouter from 'vue-router'
import More from '@/views/More.vue'
import Grid from '@/views/Grid.vue'
import Docs from '@/views/Docs.vue'
import Login from '@/views/Login.vue'
import Upload from '@/views/Upload.vue'
import Gen from '@/views/Gen.vue'
import Dashboard from '@/views/Dashboard.vue'

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
    path: '/upload',
    name: 'Upload',
    component: Upload,
  },
  // {
  //   path: '/gen',
  //   name: 'Generator',
  //   component: Gen,
  // },
  {
    path: '/admin/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/admin',
    name: 'Admin',
    component: Dashboard,
  },
  {
    path: '/:type',
    name: 'Grid',
    component: Grid,
  },
  {
    path: '/:type/:endpoint',
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
