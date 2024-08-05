package controllers

import (
	"gin_vue/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type ApplicationController struct {
}

// 请假（创建申请表）
func (con ApplicationController) AddApplication(c *gin.Context) {
	var Data models.Application
	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	if err := models.DB.Exec("INSERT INTO application (doctor_id, data,state,reason,c_time) VALUES (?, ?,?,?,?)", Data.DoctorId, Data.Data, "申请中", Data.Reason, models.UnixToTime()).Error; err != nil {
		c.JSON(200, gin.H{"message": err, "code": 201})
		return
	}
	c.JSON(200, gin.H{"message": "发送成功", "code": 200})
}

func (con ApplicationController) GetApplication(c *gin.Context) {
	var rdata []models.Application
	var Data models.Application
	if err := c.ShouldBindBodyWith(&Data, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	var Paging models.Paging
	if err := c.ShouldBindBodyWith(&Paging, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	skip := (Paging.Page - 1) * Paging.Limit

	if Data.DoctorId != 0 { //医生查看自己的申请
		models.DB.Raw("SELECT * FROM `application` WHERE doctor_id = ?", Data.DoctorId).Count(&Paging.Total)
		if err := models.DB.Raw("SELECT * FROM `application` WHERE doctor_id = ? ORDER BY c_time DESC LIMIT ? OFFSET ?", Data.DoctorId, Paging.Limit, skip).Scan(&rdata).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": rdata, "code": 200, "Total": Paging.Total})

	} else if Data.State != "" { //管理员查看对应状态的申请
		models.DB.Raw("ELECT * FROM `application` WHERE state = ?", Data.State).Count(&Paging.Total)
		if err := models.DB.Raw("SELECT * FROM `application` WHERE state = ? ORDER BY c_time DESC LIMIT ? OFFSET ?", Data.State, Paging.Limit, skip).Scan(&rdata).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": rdata, "code": 200, "Total": Paging.Total})
	} else {

		models.DB.Raw("SELECT * FROM `application`").Count(&Paging.Total)
		if err := models.DB.Raw("SELECT * FROM `application` ORDER BY c_time DESC LIMIT ? OFFSET ?", Paging.Limit, skip).Scan(&rdata).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": rdata, "code": 200, "Total": Paging.Total})
	}
}

// 对申请表的操作（修改）
func (con ApplicationController) SetApplication(c *gin.Context) {
	var Data models.Application
	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	var state string
	if err := models.DB.Raw("SELECT state FROM application WHERE application_id = ?", Data.ApplicationId).Scan(&state).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
		return
	}
	if state != "申请中" {
		c.JSON(http.StatusOK, gin.H{"message": "该状态不能进行操作", "code": 201})
		return
	}
	if err := models.DB.Model(&models.Application{}).Where("application_id = ?", Data.ApplicationId).Updates(&Data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient data"})
		return
	}

	if Data.State == "通过" {
		var r int
		models.DB.Raw("UPDATE doctors SET state = ? WHERE doctor_id = ? ", "忙碌", Data.DoctorId).Scan(r)

		if err := models.DB.Raw("SELECT department_id FROM doctors WHERE doctor_id = ?", Data.DoctorId).Scan(&r).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
			return
		}
		//查找出适合的医生进行替班
		var DoctorData models.Doctor
		models.DB.Raw("SELECT d.doctor_id FROM doctors d JOIN (SELECT doctor_id, SUM(TIME_TO_SEC(TIMEDIFF(end_time, start_time))) AS 总工作秒数 FROM doctorschedule WHERE data BETWEEN DATE_SUB(?, INTERVAL 7 DAY) AND ? AND state = 1 GROUP BY doctor_id) s ON d.doctor_id = s.doctor_id WHERE d.department_id = ? AND d.state != ? ORDER BY  s.总工作秒数 ASC", Data.Data, Data.Data, r, "忙碌").First(&DoctorData)
		models.DB.Raw("UPDATE doctorschedule SET doctor_id = ? WHERE doctor_id = ? AND data = ? AND timing = ?", DoctorData.DoctorId, Data.DoctorId, Data.Data, 0).Scan(r)
		models.DB.Raw("UPDATE doctorschedule SET doctor_id = ? WHERE doctor_id = ? AND data = ? AND timing = ?", DoctorData.DoctorId, Data.DoctorId, Data.Data, 1).Scan(r)

	}
	c.JSON(http.StatusOK, gin.H{"message": "操作成功", "code": 200})
}
