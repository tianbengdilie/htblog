import router from '@/router'
import { getToken } from '@/utils/auth' // get token from cookie
import NProgress from 'nprogress' // progress bar

// permission control
const whiteList = ['/login', '/auth-redirect']

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
