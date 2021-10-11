import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

// import viewer from '@/views/viewer'
import layout from '@/views/layout'
import authRouter from './modules/auth';
import errorPage from '@/views/error_page';

export const constantRoutes = [
  {
    path: '/',
    component: layout
  },
  {
    name: 'home',
    path: '/home',
    alias: '/'
  },
  {
    path: '*',
    component: errorPage,

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
