import service from '@/utils/request'

export const uploadFile = (data) => {
    return service({
        url: '/api/v1/blog/save_blog',
        method: 'post',
        data: data
    })
}

export const downloadFile = (data) => {
    return service({
        url: '/api/v1/blog/get_blog',
        method: 'get',
        data: data
    })
}

export const uploadImg = (data) => {
    return service({
        url: '/api/v1/blog/save_img', //
        method: 'post', //
        data: data
    })
}

export const downloadImg = (data) => { //
    return service({ //
        url: '/api/v1/blog/get_img',
        method: 'get', //
        data: data
    })
}