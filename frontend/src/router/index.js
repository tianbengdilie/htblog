import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

import layout from '@/views/layout'
import authRouter from './modules/auth';

export const constantRoutes = [
  {
    name: 'home',
    path: '/',
    alias: '/home',
    component: layout
  },
  {
    name: 'blog',
    path: '/blog',
    component: layout,
    children: [{
      path: '/blog/:uid', //
      component: () => import('@/views/viewer')
    }]
  },
  {
    path: '*',
    component: () => import('@/views/error_page'),
  },

  ...authRouter,
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: constantRoutes,

})

export default router

export const asyncRoutes = [
  /** When your routing table is too long, you can split it into small modules */
];
