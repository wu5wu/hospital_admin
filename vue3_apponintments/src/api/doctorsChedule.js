import request from '@/utils/request.js'

export const getDoctorsCheduleList = (data) => {
    return request.post('admin/schedule/getdoctorschedule',data )
}

export const finfdoctorschedule = (data) => {
    return request.post('admin/schedule/finfdoctorschedule',data )
}
export const intelligentScheduling = () => {
    return request.get('admin/schedule/intelligentScheduling')
}

export const deletedoctorschedule = (scheduleId) => {
    return request.post('admin/schedule/deletedoctorschedule',scheduleId )
}
export const deletedoctorschedules = (scheduleId) => {
    return request.post('admin/schedule/deletedoctorschedules',scheduleId )
}

export const adddoctorschedule = (data) => {
    return request.post('admin/schedule/adddoctorschedule',data )
}
export const Setdoctorschedule = (data) => {
    return request.post('/admin/schedule/setdoctorschedule' ,data)
}
