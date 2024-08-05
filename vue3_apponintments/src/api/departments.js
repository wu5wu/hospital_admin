import request from '@/utils/request.js'

export const AddDepartmentService = (data) => {
    return request.post('/admin/add' ,data)
}

export const DepartmentDeleteService = (data) => {
    return request.post('/admin/delete' ,data)
}

export const DepartmentDeleteServices = (data) => {
    return request.post('/admin/deletes' ,data)
}

export const DepartmentUpdateService = (data) => {
    return request.post('/admin/update' ,data)
}

export const DepartmentFindService = (data) => {
    return request.post('/admin/find' ,data)
}

