import router from '@/router'
import { getToken } from '@/utils/auth' // get token from cookie
import NProgress from 'nprogress' // progress bar
import authRouter from './modules/auth'

// permission control
const whiteList = (['/login', '/register']
    .concat(Array.from(authRouter, (route) => route.path))
    .concat(Array.from(authRouter, (route) => route.alias)))
    .filter((route) => route); // remove undefined element
console.log('[router.whiteList]', whiteList);

// router.beforeEach(async (to, from, next) => {
router.beforeEach((to, from, next) => {
    // start progress bar
    NProgress.start()

    // determine whether the user has logged in
    const hasToken = getToken()

    if (hasToken) {
        console.log("hasToken : ", hasToken)
        if (to.path == '/login') {
            NProgress.done()
            next('/')
        } else {
            NProgress.done()
            next()
        }
    } else {
        if (whiteList.indexOf(to.path) !== -1) {
            next()
        } else {
            next(`/login?redirect=${to.path}`)
            NProgress.done()
        }
    }
})

router.afterEach(() => {
    NProgress.done()
})

export default router
