package controllers

import (
	"fmt"
	"gin_vue/config"
	"gin_vue/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

type PatientController struct {
}
type doctorschedule struct {
	ScheduleId int
	DoctorId   int
	Data       time.Time
	StartTime  time.Time
	EndTime    time.Time
}

func (doctorschedule) TableName() string {
	return "doctorschedule"
}

// Login 注册
func (con PatientController) Login(c *gin.Context) {

	var Data models.Patient
	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	//if !tools.CaptchaVerify(tools.CaptchaData{
	//	CaptchaId: p.CaptchaId,
	//	Answer:    p.CaptchaValue,
	//}) {
	//	ResponseErr(c, CodeInvalidCaptcha)
	//	return
	//}

	if Data.Name == "" || Data.Password == "" || Data.Username == "" || Data.Identity == "" {
		c.JSON(http.StatusOK, gin.H{"error": "不能为空，请按要求进行输入"})
		return
	}
	if len(Data.Name) <= 2 || len(Data.Password) <= 5 || len(Data.Username) <= 5 || len(Data.Identity) != 18 {
		c.JSON(http.StatusOK, gin.H{"error": "信息有误， 请检查填入信息"})
		return
	}

	var mode string
	if err := models.DB.Raw("SELECT username FROM patient WHERE username = ?", Data.Username).Scan(&mode).Error; err != nil {
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
	Data.Img = "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png"
	// 执行插入操作
	if err := models.DB.Exec("INSERT INTO patient (username, password, name, identity, img) VALUES (?, ?, ?, ?, ?)", Data.Username, hashedPassword, Data.Name, Data.Identity, Data.Img).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert patient"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "Patient inserted successfully", "data": Data, "code": 200})
}

// Register 登录
func (con PatientController) Register(c *gin.Context) {
	var Data models.Patient

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
		c.JSON(http.StatusOK, gin.H{"error": "验证失败"})
		return
	}

	// 根据用户名查询数据库，验证用户信息
	var Patient models.Patient
	if err := models.DB.Raw("SELECT * FROM `patient` WHERE username = ?", Data.Username).Scan(&Patient).Error; err != nil {
		c.JSON(201, gin.H{"error": "没用该用户"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(Patient.Password), []byte(Data.Password)); err != nil {
		c.JSON(201, gin.H{"error": "密码错误"})
		return
	}
	//if Patient.Password != Data.Password {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
	//	return
	//}

	// 登录成功，生成 JWT Token
	tokenString, err := config.GenRegisteredClaims(Data.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 登录成功，返回成功响应和 Token
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString, "code": 200, "data": Patient})
}
func (con PatientController) Update(c *gin.Context) {

	//获取数据进行更新
	var updateData models.Patient
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	if updateData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
			return
		}
		updateData.Password = string(hashedPassword)
	}
	// 执行更新操作  利用gorm可以实现动态sql的效果，若字段为空，gorm会自动忽略不进行修改
	if err := models.DB.Model(&models.Patient{}).Where("patient_id = ?", updateData.PatientId).Updates(&updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient data"})
		return
	}

	// 返回更新成功的响应
	c.JSON(http.StatusOK, gin.H{"message": "Patient data updated successfully", "code": 200})
}

// FindDoctorsChedule 获取排班信息
func (con PatientController) FindDoctorsChedule(c *gin.Context) {

	//var results []map[string]interface{}
	var Data []models.DoctorsChedule
	models.DB.Model(&doctorschedule{}).Find(&Data)
	//var Data models.DoctorsChedule
	//models.DB.Find(&Data)

	c.JSON(http.StatusOK, gin.H{"date": Data})
}

// 头像更换
func (con PatientController) PatientLogin(c *gin.Context) {

	Id := c.PostForm("patient_id")
	file, err := c.FormFile("face")
	if err == nil {

		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[path.Ext(file.Filename)]; !ok {
			c.String(200, "文件不合法")
		}
		//创建文件保存目录
		day := models.GetDay()
		dir := "./static/images/" + day
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			fmt.Println("创建目录失败")
			return
		}

		//创建文件名进行保存
		unix := models.GetUnix()
		fileName := strconv.FormatInt(unix, 10) + path.Ext(file.Filename)
		dst := path.Join(dir, fileName)
		c.SaveUploadedFile(file, dst)
		var aaa int
		imgurl := "http://localhost:8080/patient/getimg?imageName=" + dst
		models.DB.Raw("UPDATE patient SET balance = ? WHERE patient_id = ?", imgurl, Id).Scan(aaa)

		c.JSON(200, gin.H{
			"path": dst,
		})
	}
}

// 找回密码，验证身份
func (con PatientController) ForgotPassword(c *gin.Context) {
	var Data models.Patient
	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	if Data.Identity == "" || Data.Username == "" {
		c.JSON(200, gin.H{"error": "不能为空", "code": 201})
		return
	}

	var Patient models.Patient
	if err := models.DB.Raw("SELECT * FROM `patient` WHERE username = ?", Data.Username).Scan(&Patient).Error; err != nil {
		c.JSON(200, gin.H{"error": "User not found", "code": 201})
		return
	}
	if Patient.Username == "" {
		c.JSON(200, gin.H{"error": "用户不存在", "code": 201})
		return
	}
	if Data.Identity != Patient.Identity {
		c.JSON(200, gin.H{"error": "身份证错误", "code": 201})
		return
	}
	c.JSON(200, gin.H{"message": "身份验证成功", "patient": Patient, "code": 200})
	return
}
