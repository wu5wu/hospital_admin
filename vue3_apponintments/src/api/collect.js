import request from '@/utils/request.js'


export const collectListService = (params,collect) => {
    return request.post('/collect/getCollect', collect, { params: params } )
}

//取消收藏
export const collectDeleteService = (id) => {
    return request.delete('/collect/delCollect/'+id)
}