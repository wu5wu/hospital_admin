package ruoters

import (
	"gin_vue/controllers"
	"github.com/gin-gonic/gin"
)

//func LoginRuotersInit(r *gin.Engine) {
//	r.POST("/patient/login", controllers.PatientController{}.Login)
//	r.POST("/patient/register", controllers.PatientController{}.Register)
//	r.POST("/patient/update", controllers.PatientController{}.Update)
//}

func LoginRuotersInit(r *gin.Engine) {
	patientGroup := r.Group("/patient")
	{
		patientController := controllers.PatientController{}

		patientGroup.POST("/login", patientController.Login)                           //注册
		patientGroup.POST("/register", patientController.Register)                     //登录
		patientGroup.POST("/update", patientController.Update)                         //更新数据
		patientGroup.POST("/fingdoctorschedule", patientController.FindDoctorsChedule) //获取排班信息
		patientGroup.POST("/topup", controllers.AlipayController{}.TopUp)              //进行余额充值
		patientGroup.POST("/doupload", controllers.ToolsController{}.Login)            //上传图片
		patientGroup.GET("/getimg", controllers.ToolsController{}.GetImage)            //获取图片
		patientGroup.POST("/forgotPassword", patientController.ForgotPassword)         //验证身份，找回密码
		patientGroup.POST("/getbalance", controllers.AlipayController{}.GetBalance)    //获取余额
		patientGroup.POST("/genCaptcha", controllers.YanzhengController{}.GenCaptcha)  //生成验证码
	}
}
