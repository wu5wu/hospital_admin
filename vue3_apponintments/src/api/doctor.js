import request from '@/utils/request.js'
//修改医生数据
export const UpdateDoctors = (data) => {
    return request.post('/doctor/update',data )
}
//获取医生列表
export const GetDoctors = (data) => {
    return request.post('/admin/alldoctor',data )
}
//删除医生
export const DeleteDoctors = (data) => {
    return request.post('/doctor/delete',data )
}
//添加医生
export const AddDoctors = (data) => {
    return request.post('/doctor/add',data )
}

//获取对应患者列表
export const Gtepatient = (data) => {
    return request.post('/doctor/gtepatient',data )
}

//预约已完成
export const Setappointment = (data) => {
    return request.post('/doctor/setappointment',data )
}

//获取请假表
export const Getapplication = (data) => {
    return request.post('/doctor/application/get',data )
}

//进行请假
export const Addapplication = (data) => {
    return request.post('/doctor/application/add',data )
}


//请假单的修改
export const Setapplication = (data) => {
    return request.post('/doctor/application/set',data )
}


//请假单的修改
export const Gtedoctorschedule = (data) => {
    return request.post('/doctor/gtedoctorschedule',data )
}
