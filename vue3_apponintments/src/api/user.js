import request from '@/utils/request.js'

// export const userRegisterService = (registerData)=>{
//     request.post('',registerData)
// }
//注册
export const userLoginService = (loginData)=>{
    return request.post('/login',loginData)
}
// 登录
export const userRegisterService = (loginData)=>{
    return request.post('/register',loginData)
}
//修改头像
export const userAvatarUpdateService = (HeadData)=>{
    return request.post('/user/UpdateHead',HeadData)
}
//修改个人信息
export const userInfoUpdateService = (userInfo)=>{
    return request.post('/user/updateuser',userInfo)
}
//获取用户列表
export const userListService = (params,user) => {
    return request.post('/user/userList', user, { params: params } )
}
//删除
export const userDeleteService = (id) => {
    return request.delete('/deluser/'+id)
}