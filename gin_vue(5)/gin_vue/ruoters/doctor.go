package ruoters

import (
	"gin_vue/config"
	"gin_vue/controllers"
	"github.com/gin-gonic/gin"
)

func DoctorRuotersInit(r *gin.Engine) {
	patientGroup := r.Group("/doctor")
	{
		patientController := controllers.DoctorController{}
		patientGroup.POST("/add", patientController.AddDoctor)                                                        //医生的创建
		patientGroup.POST("/register", patientController.Register)                                                    //医生登录
		patientGroup.POST("/update", config.ValidateRegisteredClaims, patientController.Update)                       //修改
		patientGroup.POST("/delete", config.ValidateRegisteredClaims, patientController.Delete)                       //删除
		patientGroup.POST("/gtedoctorschedule", config.ValidateRegisteredClaims, patientController.GteDoctorsChedule) //查看个人排班
		patientGroup.POST("/gtepatient", patientController.Gtepatient)                                                //查看个人患者信息
		patientGroup.POST("/setappointment", patientController.SetAppointment)
	}

	applicationGroup := r.Group("/doctor/application")
	{
		applicationController := controllers.ApplicationController{}
		applicationGroup.POST("/add", applicationController.AddApplication) //申请签到（医生）
		applicationGroup.POST("/get", applicationController.GetApplication) //获取申请表（医生）（管理员）
		applicationGroup.POST("/set", applicationController.SetApplication) //对申请表的操作（修改）（管理员）
	}
}
