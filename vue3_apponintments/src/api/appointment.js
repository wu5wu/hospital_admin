import request from "@/utils/request.js";

//值班展示列表
export const getdoctorschedule = (data) => {
  return request.post("/appointment/getdoctorschedule", data);
};

//获取科室信息
export const getDepartmentList = () => {
  return request.post("/admin/findAll");
};

//获取日期
export const getAppointmentDate = () => {
    return request.get("/appointment/getdate");
};

//点击预约进行预约的创建
export const getAppointmentAdd = (data) => {
    return request.post("/appointment/add", data);
};

//获取个人预约列表
export const getAppointment = (data) => {
  return request.post("/appointment/getappointment", data);
};

//获取个人预约列表
export const deleteAppointment = (data) => {
  return request.post("/appointment/delete", data);
};

