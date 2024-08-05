package ruoters

import (
	"gin_vue/controllers"
	"github.com/gin-gonic/gin"
)

func AppointmentRuotersInit(r *gin.Engine) {
	patientGroup := r.Group("/appointment")
	{
		appointmentControllers := controllers.AppointmentControllers{}
		patientGroup.GET("/getdate", appointmentControllers.GetDate)                      //获取后七天日期
		patientGroup.POST("/add", appointmentControllers.AddAppointment)                  //进行预约（预约表的增加）
		patientGroup.POST("/delete", appointmentControllers.DeleteAppointment)            //取消预约（预约表的删除）
		patientGroup.POST("/getappointment", appointmentControllers.GetAppointment)       //获取预约表
		patientGroup.POST("/getdoctorschedule", appointmentControllers.GetDoctorsChedule) //获取医生值班表
		patientGroup.POST("/setappointment", appointmentControllers.SetAppointment)       //修改预约表状态（取消预约或完成或签到）
	}
}
