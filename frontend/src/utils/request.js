import axios from 'axios' // 引入axios
// import store from '@/store'
import { emitter } from '@/utils/bus.js'

const service = axios.create({
    baseURL: "http://9.135.152.118:3000",
    timeout: 99999,
    withCredentials: true
})
let acitveAxios = 0
let timer
const showLoading = () => {
    acitveAxios++
    if (timer) {
        clearTimeout(timer)
    }
    timer = setTimeout(() => {
        if (acitveAxios > 0) {
            emitter.emit('showLoading')
        }
    }, 400)
}

const closeLoading = () => {
    acitveAxios--
    if (acitveAxios <= 0) {
        clearTimeout(timer)
        emitter.emit('closeLoading')
    }
}
// http request 拦截器
service.interceptors.request.use(
    config => {
        if (!config.donNotShowLoading) {
            showLoading()
        }
        // const token = store.getters['user/token']
        config.data = JSON.stringify(config.data)
        config.headers = {
            'Content-Type': 'application/json',
            // 'x-token': token,
        }
        return config
    },
    error => {
        closeLoading()
        return error
    }
)

// // http response 拦截器
// service.interceptors.response.use(
//     response => {
//         closeLoading()
//         if (response.headers['new-token']) {
//             store.commit('user/setToken', response.headers['new-token'])
//         }
//         if (response.data.code === 0 || response.headers.success === 'true') {
//             if (response.headers.msg) {
//                 response.data.msg = decodeURI(response.headers.msg)
//             }
//             return response.data
//         } else {
//             if (response.data.data && response.data.data.reload) {
//                 store.commit('user/LoginOut')
//             }
//             return response.data.msg ? response.data : response
//         }
//     },
//     error => {
//         closeLoading()
//         console.log(error.response)
//         return error
//     }
// )

export default service