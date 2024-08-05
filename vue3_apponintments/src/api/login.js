import request from '@/utils/request.js'


export const AdminRegisterService = (data) => {
    return request.post('/admin/register',data )
}

export const DoctorRegisterService = (data) => {
    return request.post('/doctor/register' ,data)
}

export const PatientRegisterService = (data) => {
    return request.post('/patient/register' ,data)
}


export const PatientLoginService = (data) => {
    return request.post('/patient/login' ,data)
}
