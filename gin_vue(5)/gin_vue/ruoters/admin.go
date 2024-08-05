package ruoters

import (
	"gin_vue/config"
	"gin_vue/controllers"
	"github.com/gin-gonic/gin"
)

func AdminRuotersInit(r *gin.Engine) {
	patientGroup := r.Group("/admin")
	{
		patientController := controllers.AdminController{}

		patientGroup.POST("/register", patientController.Register)
		patientGroup.POST("/update", config.ValidateRegisteredClaims, patientController.UpdateDepartment)   //修改科室
		patientGroup.POST("/add", config.ValidateRegisteredClaims, patientController.AddDepartment)         //添加科室
		patientGroup.POST("/delete", config.ValidateRegisteredClaims, patientController.DeleteDepartment)   //删除科室
		patientGroup.POST("/deletes", config.ValidateRegisteredClaims, patientController.DeleteDepartments) //批量删除科室
		patientGroup.POST("/findAll", patientController.FindAllDepartment)                                  //遍历科室
		patientGroup.POST("/find", config.ValidateRegisteredClaims, patientController.FindDepartment)       //查找科室
		patientGroup.POST("/alldoctor", config.ValidateRegisteredClaims, patientController.FindAllDoctor)   //查找所有的医生
	}
	scheduleGroup := r.Group("/admin/schedule")
	{
		scheduleController := controllers.AdminController{}
		scheduleGroup.POST("/getdoctorschedule", scheduleController.GetDoctorsChedule)                                          //获取排班表
		scheduleGroup.GET("/intelligentScheduling", config.ValidateRegisteredClaims, scheduleController.IntelligentScheduling)  //对排班表进行初始创建,并返回排班结果
		scheduleGroup.POST("/setdoctorschedule", config.ValidateRegisteredClaims, scheduleController.SetDoctorsChedule)         //对排班表进行修改
		scheduleGroup.POST("/finfdoctorschedule", config.ValidateRegisteredClaims, scheduleController.FinfDoctorsChedule)       //对排班表进行查找
		scheduleGroup.POST("/deletedoctorschedule", config.ValidateRegisteredClaims, scheduleController.DeleteDoctorsChedule)   //对排班表进行删除
		scheduleGroup.POST("/deletedoctorschedules", config.ValidateRegisteredClaims, scheduleController.DeleteDoctorsChedules) //对排班表进行批量删除

		scheduleGroup.POST("/adddoctorschedule", config.ValidateRegisteredClaims, scheduleController.AddDoctorsChedule) //对排班表进行添加
	}
	patientsGroup := r.Group("/admin/patient")
	{
		patientController := controllers.AdminController{}
		patientsGroup.POST("/patientList", patientController.PatientList)     //获取用户列表
		patientsGroup.POST("/deletePatient", patientController.DeletePatient) //删除用户
	}

}
