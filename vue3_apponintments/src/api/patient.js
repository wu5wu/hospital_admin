// import request from '@/utils/request.js'

// export const getDate = () => {
//     return request.post('appointment/getdate' )
// }

// export const finfdoctorschedule = (data) => {
//     return request.post('appointment/getdate',data )
// }


import request from '@/utils/request.js'

export const getDate = () => {
    return request.post('appointment/getdate' )
}

export const finfdoctorschedule = (data) => {
    return request.post('appointment/getdate',data )
}

//患者查询已有的预约
export const findAppointment = (PatientId) => {
  return request.post("/appointment/getappointment", PatientId);
};

//查看已预约列表，可以进行取消预约
export const SetAppointment = (appointment) => {
  return request.post("/appointment/setappointment", appointment);
};

// //充值
// export const TopUp = (appointment) => {
//   return request.post("/patient/topup", appointment);
// };
//充值
export const TopUp = (appointment) => {
  return request.post("/alipay/sandbox/tradePagePay", appointment);
};


//修改
export const  Update= (data) => {
  return request.post("/patient/update", data);
};


//获取用户列表
export const  GetpatientList= (data) => {
  return request.post("/admin/patient/patientList", data);
};

//获取用户列表
export const  Deletepatient= (data) => {
  return request.post("/admin/patient/deletePatient", data);
};

//身份验证
export const  ForgotPassword= (data) => {
  return request.post("/patient/forgotPassword", data);
};

//获取余额
export const  Getbalance= (data) => {
  return request.post("/patient/getbalance", data);
};

//获取验证码
export const  GenCaptcha= () => {
  return request.post("/patient/genCaptcha");
};
