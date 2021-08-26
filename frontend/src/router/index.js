import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

import home from '@/views/home'
import viewer from '@/views/viewer'
import error_page from '@/views/404'
import login from '@/views/login'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: home
  }, {
    path: '/login',
    name: 'Login',
    component: login
  },
  {
    path: '/home',
    redirect: '/'
  },
  {
    path: '/viewer',
    name: 'Viewer',
    component: viewer
  },
  {
    path: '*',
    name: '404',
    component: error_page
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
