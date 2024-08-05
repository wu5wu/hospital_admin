package controllers

import (
	"errors"
	"gin_vue/config"
	"gin_vue/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type DoctorController struct {
}

// Register 登录
func (con DoctorController) Register(c *gin.Context) {

	var Data models.Doctor
	if err := c.ShouldBindBodyWith(&Data, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	var param CaptchaData
	if err := c.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	if !CaptchaVerify(param) {
		c.JSON(http.StatusOK, gin.H{"error": "验证码错误"})
		return
	}

	// 根据用户名查询数据库，验证用户信息
	var doctor models.Doctor
	err := models.DB.Raw("SELECT * FROM doctors WHERE username = ?", Data.Username).Scan(&doctor).Error
	if err != nil {
		// 如果查询结果为空，返回用户未找到的错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		// 查询出错，返回错误信息
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user data"})
		return
	}

	// 验证密码是否匹配
	//if doctor.Password != Data.Password {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
	//	return
	//}

	if err := bcrypt.CompareHashAndPassword([]byte(doctor.Password), []byte(Data.Password)); err != nil {
		c.JSON(201, gin.H{"error": "密码错误"})
		return
	}

	// 登录成功，生成 JWT Token
	tokenString, err := config.GenRegisteredClaims(Data.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 登录成功，返回成功响应和 Token
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString, "code": 200, "data": doctor})
}

// Update 进行数据修改更新
func (con DoctorController) Update(c *gin.Context) {

	//获取数据进行更新
	var updateData models.Doctor
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	// 执行更新操作  利用gorm可以实现动态sql的效果，若字段为空，gorm会自动忽略不进行修改
	if err := models.DB.Model(&models.Doctor{}).Where("doctor_id = ?", updateData.DoctorId).Updates(&updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient data"})
		return
	}

	// 返回更新成功的响应
	c.JSON(http.StatusOK, gin.H{"message": "数据修改成功", "code": 200})
}

// 查看个人值班
func (con DoctorController) GteDoctorsChedule(c *gin.Context) {
	var updateData models.Doctor
	var Data []models.DoctorsChedule

	if err := c.ShouldBindBodyWith(&updateData, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	var Paging models.Paging
	if err := c.ShouldBindBodyWith(&Paging, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	skip := (Paging.Page - 1) * Paging.Limit

	if updateData.DoctorId != 0 {

		models.DB.Raw("SELECT a.*,b.name,c.name as departments FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id  And b.doctor_id = ? ORDER BY a.`data` DESC LIMIT ? OFFSET ?", updateData.DoctorId, Paging.Limit, skip).Scan(&Data)
		models.DB.Raw("SELECT a.*,b.name,c.name as departments FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id  And b.doctor_id = ? ", updateData.DoctorId).Count(&Paging.Total)

		c.JSON(200, gin.H{"message": Data, "Total": Paging.Total, "code": 200})
		return

	} else {
		c.JSON(201, gin.H{"message": "查询失败", "code": 201})
		return
	}

	return
}

type dataclass struct {
	DoctorId int
	Date     string
	Name     string
	Status   string
}
type redata struct {
	Name          string
	Identity      string
	Phone         string
	AppointmentId int
	Created       string
	Data          string
	SignTime      string
	Status        string
	Img           string
}

// 查看个人患者
func (con DoctorController) Gtepatient(c *gin.Context) {
	var patients []redata
	var Data dataclass

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

	if Data.Date != "" && Data.Status == "" {
		models.DB.Raw("SELECT a.`name`,a.identity,a.phone,b.appointment_id ,b.created,d.`data`,b.sign_time,b.`status` FROM patient a, appointment b,doctorschedule d WHERE b.patient_id = a.patient_id AND b.schedule_id = d.schedule_id AND d.doctor_id = ? AND d.`data` = ?", Data.DoctorId, Data.Date).Count(&Paging.Total)
		models.DB.Raw("SELECT a.`name`,a.identity,a.phone,b.appointment_id ,b.created,d.`data`,b.sign_time,b.`status` ,a.img FROM patient a, appointment b,doctorschedule d WHERE b.patient_id = a.patient_id AND b.schedule_id = d.schedule_id AND d.doctor_id = ? AND d.`data` = ?  LIMIT ? OFFSET ?", Data.DoctorId, Data.Date, Paging.Limit, skip).Scan(&patients)
		c.JSON(200, gin.H{"message": patients, "code": 200, "Total": Paging.Total})
	} else if Data.Name != "" {
		models.DB.Raw("SELECT a.`name`,a.identity,a.phone,b.appointment_id ,b.created,d.`data`,b.sign_time,b.`status` FROM patient a, appointment b,doctorschedule d WHERE b.patient_id = a.patient_id AND b.schedule_id = d.schedule_id AND d.doctor_id = ? AND a.`name` LIKE ?", Data.DoctorId, "%"+Data.Name+"%").Count(&Paging.Total)
		models.DB.Raw("SELECT a.`name`,a.identity,a.phone,b.appointment_id ,b.created,d.`data`,b.sign_time,b.`status` ,a.img FROM patient a, appointment b,doctorschedule d WHERE b.patient_id = a.patient_id AND b.schedule_id = d.schedule_id AND d.doctor_id = ? AND a.`name` LIKE ? LIMIT ? OFFSET ?", Data.DoctorId, "%"+Data.Name+"%", Paging.Limit, skip).Scan(&patients)
		c.JSON(200, gin.H{"message": patients, "code": 200})
	} else if Data.Status != "" && Data.Date == "" {
		models.DB.Raw("SELECT a.`name`,a.identity,a.phone,b.appointment_id ,b.created,d.`data`,b.sign_time ,b.`status` FROM patient a, appointment b,doctorschedule d WHERE b.patient_id = a.patient_id AND b.schedule_id = d.schedule_id AND d.doctor_id = ? AND b.`status`=? ", Data.DoctorId, Data.Status).Count(&Paging.Total)
		models.DB.Raw("SELECT a.`name`,a.identity,a.phone,b.appointment_id ,b.created,d.`data`,b.sign_time ,b.`status` ,a.img FROM patient a, appointment b,doctorschedule d WHERE b.patient_id = a.patient_id AND b.schedule_id = d.schedule_id AND d.doctor_id = ? AND b.`status`=? ORDER BY b.sign_time ASC LIMIT ? OFFSET ? ", Data.DoctorId, Data.Status, Paging.Limit, skip).Scan(&patients)
		c.JSON(200, gin.H{"message": patients, "code": 200, "Total": Paging.Total})
	} else if Data.Status != "" && Data.Date != "" {
		models.DB.Raw("SELECT a.`name`,a.identity,a.phone,b.appointment_id ,b.created,d.`data`,b.sign_time ,b.`status` FROM patient a, appointment b,doctorschedule d WHERE b.patient_id = a.patient_id AND b.schedule_id = d.schedule_id AND d.doctor_id = ? AND b.`status`=?  AND d.`data` = ?", Data.DoctorId, Data.Status, Data.Date).Count(&Paging.Total)
		models.DB.Raw("SELECT a.`name`,a.identity,a.phone,b.appointment_id ,b.created,d.`data`,b.sign_time ,b.`status` ,a.img FROM patient a, appointment b,doctorschedule d WHERE b.patient_id = a.patient_id AND b.schedule_id = d.schedule_id AND d.doctor_id = ? AND b.`status`=?  AND d.`data` = ? ORDER BY b.sign_time ASC LIMIT ? OFFSET ? ", Data.DoctorId, Data.Status, Data.Date, Paging.Limit, skip).Scan(&patients)
		c.JSON(200, gin.H{"message": patients, "code": 200, "Total": Paging.Total})
	} else {

		models.DB.Raw("SELECT a.`name`,a.identity,a.phone,b.appointment_id ,b.created,d.`data`,b.sign_time,b.`status` FROM patient a, appointment b,doctorschedule d WHERE b.patient_id = a.patient_id AND b.schedule_id = d.schedule_id AND d.doctor_id = ? ORDER BY b.created DESC", Data.DoctorId).Count(&Paging.Total)
		models.DB.Raw("SELECT a.`name`,a.identity,a.phone,b.appointment_id ,b.created,d.`data`,b.sign_time,b.`status` ,a.img FROM patient a, appointment b,doctorschedule d WHERE b.patient_id = a.patient_id AND b.schedule_id = d.schedule_id AND d.doctor_id = ? ORDER BY b.created DESC LIMIT ? OFFSET ? ", Data.DoctorId, Paging.Limit, skip).Scan(&patients)
		c.JSON(200, gin.H{"message": patients, "code": 200, "Total": Paging.Total})
	}
}

// 修改预约表状态为已完成
func (con DoctorController) SetAppointment(c *gin.Context) {
	var data models.Appointment

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(401, gin.H{"error": "Invalid data format"})
		return
	}
	var Status string
	models.DB.Raw("SELECT `status` FROM `appointment` WHERE appointment_id = ?", data.AppointmentId).Scan(&Status)
	if Status == "已签到" {
		result := models.DB.Exec("UPDATE appointment SET `status` = ? WHERE appointment_id = ?", data.Status, data.AppointmentId)
		if result.Error != nil {
			// 处理错误
			c.JSON(201, gin.H{"message": "修改失败", "code": 201})
			return
		}
		c.JSON(200, gin.H{"message": "操作成功", "code": 200})
	} else if Status == "已完成" {
		c.JSON(200, gin.H{"message": "请不要重复操作", "code": 201})
	} else if Status == "已取消" {
		c.JSON(200, gin.H{"message": "该预约已取消", "code": 201})
	} else {
		c.JSON(200, gin.H{"message": "预约还未签到", "code": 201})
	}

}

// Login 注册
func (con DoctorController) AddDoctor(c *gin.Context) {

	var Data models.Doctor
	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	var mode string
	if err := models.DB.Raw("SELECT username FROM doctors WHERE username = ?", Data.Username).Scan(&mode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
		return
	}

	if mode == Data.Username {
		c.JSON(http.StatusOK, gin.H{"error": "该用户名已存在"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return
	}
	img := "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png"
	// 执行插入操作
	if err := models.DB.Exec("INSERT INTO doctors (username, password, name, department_id, price,introduction,img,state) VALUES (?, ?, ?, ?, ?,?,?,?)", Data.Username, hashedPassword, Data.Name, Data.DepartmentId, Data.Price, Data.Introduction, img, "空闲").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert patient"})
		return
	}

	var doctorId int
	if err = models.DB.Raw("SELECT doctor_id FROM doctors WHERE username = ?", Data.Username).Scan(&doctorId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
		return
	}
	//创建一个排班基础模板数据
	if err = models.DB.Exec("INSERT INTO doctorschedule (doctor_id, data, timing, state) VALUES (?, ?, ?,?)", doctorId, models.GetDate(), 0, 1).Error; err != nil {

	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "创建成功", "data": Data, "code": 200})
}

// 删除医生
func (con DoctorController) Delete(c *gin.Context) {
	var Id models.Doctor
	if err := c.ShouldBindJSON(&Id); err != nil {
		c.JSON(401, gin.H{"error": "Invalid data format"})
		return
	}
	//scheduleId := c.Query("scheduleId")

	result := models.DB.Exec("DELETE FROM doctors WHERE doctor_id = ?", Id.DoctorId)
	if result.Error != nil {
		// 处理错误
		c.JSON(201, gin.H{"message": "删除失败", "code": 201})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功", "code": 200})
	return
}
