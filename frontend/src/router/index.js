import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

// import home from '@/views/home'
// import viewer from '@/views/viewer'
// import error_page from '@/views/404'
// import login from '@/views/login'
import layout from '@/views/layout'
import permissionRouter from './modules/permission';
import authRouter from './modules/auth';

// const routes = [
//   {
//     path: '/',
//     name: 'Home',
//     component: home
//   },
//   {
//     path: '/login',
//     name: 'Login',
//     component: login
//   },
//   {
//     path: '/layout',
//     name: 'Layout',
//     component: layout
//   },
//   {
//     path: '/home',
//     redirect: '/'
//   },
//   {
//     path: '/viewer',
//     name: 'Viewer',
//     component: viewer
//   },
//   {
//     path: '*',
//     name: '404',
//     component: error_page
//   }
// ]

export const constantRoutes = [
  {
    path: '/',
    component: layout
  },

  ...authRouter,
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: constantRoutes
})

export default router

export const asyncRoutes = [
  /** When your routing table is too long, you can split it into small modules */
  permissionRouter,
  { path: '*', redirect: '/error/404', hidden: true },
];
