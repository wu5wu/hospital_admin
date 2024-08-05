package controllers

import (
	"fmt"
	"gin_vue/config"
	"gin_vue/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AdminController struct {
}

// Register 登录
func (con AdminController) Register(c *gin.Context) {

	var Data models.Admin
	if err := c.ShouldBindBodyWith(&Data, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	var param CaptchaData
	if err := c.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	fmt.Println(param)
	if !CaptchaVerify(param) {
		c.JSON(http.StatusOK, gin.H{"error": "验证码错误"})
		return
	}

	// 根据用户名查询数据库，验证用户信息
	var admin models.Admin
	if err := models.DB.Where("username = ?", Data.Username).First(&admin).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "账户不存在"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(Data.Password)); err != nil {
		c.JSON(200, gin.H{"error": "密码错误"})
		return
	}
	// 登录成功，生成 JWT Token
	tokenString, err := config.GenRegisteredClaims(Data.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 登录成功，返回成功响应和 Token
	c.JSON(200, gin.H{"message": "Login successful", "token": tokenString, "data": admin, "code": 200})
}

// UpdateDepartment 科室信息进行数据修改更新
func (con AdminController) UpdateDepartment(c *gin.Context) {

	//获取数据进行更新
	var updateData models.Department
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	var mode string
	if err := models.DB.Raw("SELECT name FROM departments WHERE name = ?", updateData.Name).Scan(&mode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
		return
	}

	if mode == updateData.Name {
		c.JSON(http.StatusOK, gin.H{"error": "该科室已存在"})
		return
	}

	// 执行更新操作  利用gorm可以实现动态sql的效果，若字段为空，gorm会自动忽略不进行修改
	if err := models.DB.Model(&models.Department{}).Where("department_id = ?", updateData.DepartmentId).Updates(&updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient data"})
		return
	}

	// 返回更新成功的响应
	c.JSON(200, gin.H{"message": "修改成功", "code": 200})
}

// AddDepartment 进行科室的增加
func (con AdminController) AddDepartment(c *gin.Context) {
	var updateData models.Department
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	var mode string
	if err := models.DB.Raw("SELECT name FROM departments WHERE name = ?", updateData.Name).Scan(&mode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
		return
	}

	if mode == updateData.Name {
		c.JSON(http.StatusOK, gin.H{"error": "该科室已存在"})
		return
	}

	// 执行插入操作
	if err := models.DB.Exec("INSERT INTO departments (name, class) VALUES (?, ?)", updateData.Name, updateData.Class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert patient"})
		return
	}

	// 返回成功响应
	c.JSON(200, gin.H{"message": "科室创建 successfully", "code": 200})
}

// 科室的删除
func (con AdminController) DeleteDepartment(c *gin.Context) {
	var Data models.Department
	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(401, gin.H{"error": "Invalid data format"})
		return
	}

	// 执行删除操作
	if err := models.DB.Delete(&models.Department{}, Data.DepartmentId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient data"})
		return
	}

	// 返回更新成功的响应
	c.JSON(200, gin.H{"message": "删除成功", "code": 200})

}

// 科室的批量删除
func (con AdminController) DeleteDepartments(c *gin.Context) {
	var DepartmentIds []int
	if err := c.ShouldBindJSON(&DepartmentIds); err != nil {
		c.JSON(401, gin.H{"error": "Invalid data format"})
		return
	}

	for _, id := range DepartmentIds {
		if err := models.DB.Delete(&models.Department{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient data"})
			return
		}
	}
	c.JSON(200, gin.H{"message": "批量删除成功", "code": 200})

}

// FindDepartment 按名称进行查找
func (con AdminController) FindDepartment(c *gin.Context) {
	var Data models.Department
	var Datas []models.Department

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

	if Data.Name != "" {
		models.DB.Raw("SELECT * FROM departments WHERE name LIKE ?", "%"+Data.Name+"%").Count(&Paging.Total)
		if err := models.DB.Raw("SELECT * FROM departments WHERE name LIKE ? LIMIT ? OFFSET ?", "%"+Data.Name+"%", Paging.Limit, skip).Scan(&Datas).Error; err != nil {
			c.JSON(401, gin.H{"error": "Failed to check existing username"})
			return
		}
		if Datas == nil {
			c.JSON(200, gin.H{"error": "不存在"})
			return
		}
		c.JSON(200, gin.H{"message": Datas, "code": 200, "Total": Paging.Total})
		return
	} else {
		models.DB.Raw("SELECT * FROM departments ").Count(&Paging.Total)
		if err := models.DB.Raw("SELECT * FROM departments  LIMIT ? OFFSET ?", Paging.Limit, skip).Scan(&Datas).Error; err != nil {
			c.JSON(401, gin.H{"error": "Failed to check existing username"})
			return
		}
		c.JSON(200, gin.H{"message": Datas, "code": 200, "Total": Paging.Total})
	}

}

// FindAllDepartment 查找全部进行遍历
func (con AdminController) FindAllDepartment(c *gin.Context) {

	var Data []models.Department

	models.DB.Raw("SELECT * FROM departments ").Scan(&Data)

	c.JSON(200, gin.H{"message": Data})

	return
}

// 获取所有医生和对应的科室信息
func findAllDoctor() []models.Doctor {
	var Data []models.Doctor
	models.DB.Raw("SELECT doctor_id,department_id FROM doctors ").Scan(&Data)

	return Data
}

// FindAllDoctor 获取所有的医生信息
func (con AdminController) FindAllDoctor(c *gin.Context) {
	var Datas models.Doctor
	if err := c.ShouldBindBodyWith(&Datas, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	var Paging models.Paging
	if err := c.ShouldBindBodyWith(&Paging, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	var Data []models.Doctor

	if Datas.Name != "" && Datas.DepartmentId == 0 { //按姓名查询
		skip := (Paging.Page - 1) * Paging.Limit
		models.DB.Raw("SELECT a.doctor_id,a.username,a.`name`,a.introduction,a.price,b.`name` as  Department FROM doctors a , departments b WHERE a.department_id = b.department_id AND a.`name` LIKE ? ", "%"+Datas.Name+"%").Count(&Paging.Total)
		if int(Paging.Total) < skip {
			c.JSON(200, gin.H{"message": "当前页面数过大", "code": 201})
		}
		if err := models.DB.Raw("SELECT a.doctor_id,a.username,a.`name`,a.introduction,a.price,b.`name` as  Department ,a.img , b.department_id , a.state FROM doctors a , departments b WHERE a.department_id = b.department_id AND a.`name` LIKE ? LIMIT ? OFFSET ?", "%"+Datas.Name+"%", Paging.Limit, skip).Scan(&Data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
			return
		}
	} else if Datas.Name != "" && Datas.DepartmentId != 0 { //姓名加科室进行关联查询
		skip := (Paging.Page - 1) * Paging.Limit
		models.DB.Raw("SELECT a.doctor_id,a.username,a.`name`,a.introduction,a.price,b.`name` as  Department FROM doctors a , departments b WHERE a.department_id = b.department_id AND a.`name` LIKE ? AND a.department_id = ?", "%"+Datas.Name+"%", Datas.DepartmentId).Count(&Paging.Total)
		if int(Paging.Total) < skip {
			c.JSON(200, gin.H{"message": "当前页面数过大", "code": 201})
		}
		if err := models.DB.Raw("SELECT a.doctor_id,a.username,a.`name`,a.introduction,a.price,b.`name` as  Department ,a.img , b.department_id , a.state FROM doctors a , departments b WHERE a.department_id = b.department_id AND a.`name` LIKE ? AND a.department_id = ? LIMIT ? OFFSET ?", "%"+Datas.Name+"%", Datas.DepartmentId, Paging.Limit, skip).Scan(&Data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
			return
		}
	} else if Datas.Name == "" && Datas.DepartmentId != 0 { //按科室查
		skip := (Paging.Page - 1) * Paging.Limit
		models.DB.Raw("SELECT a.doctor_id,a.username,a.`name`,a.introduction,a.price,b.`name` as  Department FROM doctors a , departments b WHERE a.department_id = b.department_id AND a.department_id = ? ", Datas.DepartmentId).Count(&Paging.Total)
		if int(Paging.Total) < skip {
			c.JSON(200, gin.H{"message": "当前页面数过大", "code": 201})
		}
		if err := models.DB.Raw("SELECT a.doctor_id,a.username,a.`name`,a.introduction,a.price,b.`name` as  Department ,a.img , b.department_id , a.state FROM doctors a , departments b WHERE a.department_id = b.department_id AND a.department_id = ? LIMIT ? OFFSET ?", Datas.DepartmentId, Paging.Limit, skip).Scan(&Data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
			return
		}
	} else { //显示全部
		skip := (Paging.Page - 1) * Paging.Limit
		models.DB.Raw("SELECT a.doctor_id,a.username,a.`name`,a.introduction,a.price,b.`name` as  Department FROM doctors a , departments b WHERE a.department_id = b.department_id ").Count(&Paging.Total)
		if int(Paging.Total) < skip {
			c.JSON(200, gin.H{"message": "当前页面数过大", "code": 201})
		}
		if err := models.DB.Raw("SELECT a.doctor_id,a.username,a.`name`,a.introduction,a.price,b.`name` as  Department ,a.img , b.department_id , a.state FROM doctors a , departments b WHERE a.department_id = b.department_id LIMIT ? OFFSET ?", Paging.Limit, skip).Scan(&Data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
			return
		}
	}

	c.JSON(200, gin.H{"message": Data, "code": 200, "Total": Paging.Total})
	return
}

//// FindAllDoctor 进行当天排班的数据创建,默认空闲
//func (con AdminController) AddDoctorsChedule(c *gin.Context) {
//	// 获取当前日期
//	currentDate := time.Now().Format("2006-01-02")
//
//	var Data = findAllDoctor()
//	for _, value := range Data {
//		// 执行插入操作
//		if err := models.DB.Exec("INSERT INTO doctorschedule (doctor_id, data, timing, state) VALUES (?, ?, ?,?)", value.DoctorId, currentDate, 0, 0).Error; err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert patient"})
//			return
//		}
//		if err := models.DB.Exec("INSERT INTO doctorschedule (doctor_id, data,timing, state) VALUES (?, ?, ?,?)", value.DoctorId, currentDate, 1, 0).Error; err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert patient"})
//			return
//		}
//		// 返回成功响应
//	}
//	c.JSON(200, gin.H{"message": Data})
//	return
//}

// 进行当天排班的数据创建,默认空闲
func addDoctorsChedule(currentDate string) {

	var Data = findAllDoctor() //获取所有的医生
	for _, value := range Data {
		// 执行插入操作
		if err := models.DB.Exec("INSERT INTO doctorschedule (doctor_id, data, timing, state) VALUES (?, ?, ?,?)", value.DoctorId, currentDate, 0, 0).Error; err != nil {

		}
		if err := models.DB.Exec("INSERT INTO doctorschedule (doctor_id, data,timing, state) VALUES (?, ?, ?,?)", value.DoctorId, currentDate, 1, 0).Error; err != nil {

		}
		// 返回成功响应
	}
}

// 设置初始初始数据,默认空闲
func addDoctorsChedule1(currentDate string) {

	var Data = findAllDoctor()
	for _, value := range Data {
		// 执行插入操作
		if err := models.DB.Exec("INSERT INTO doctorschedule (doctor_id, data, timing, state) VALUES (?, ?, ?,?)", value.DoctorId, currentDate, 0, 1).Error; err != nil {

		}
		if err := models.DB.Exec("INSERT INTO doctorschedule (doctor_id, data,timing, state) VALUES (?, ?, ?,?)", value.DoctorId, currentDate, 1, 1).Error; err != nil {

		}
		// 返回成功响应
	}
}

// 检查对应日期是否已有排班，如果没有，则进行创建
func examineDoctorsChedule(date string) {
	var Data models.Doctor
	models.DB.Raw("SELECT doctor_id FROM doctorschedule WHERE data =? ", date).First(&Data)
	if Data.DoctorId == 0 {
		//判断出当前日期没有数据，于是调用创建排班数据的函数
		addDoctorsChedule(date)
	} else {
		fmt.Println(date, "已有数据")
	}

}

// 获取当天对应科室工时最少的医生，并返回其id
func getDoctor(date string, r int) int {
	var Data models.Doctor
	models.DB.Raw("SELECT d.doctor_id FROM doctors d JOIN (SELECT doctor_id, SUM(TIME_TO_SEC(TIMEDIFF(end_time, start_time))) AS 总工作秒数 FROM doctorschedule WHERE data BETWEEN DATE_SUB(?, INTERVAL 7 DAY) AND ? AND state = 1 GROUP BY doctor_id) s ON d.doctor_id = s.doctor_id WHERE d.department_id = ? ORDER BY  s.总工作秒数 ASC", date, date, r).First(&Data)
	return Data.DoctorId
}

// 修改当天医生的数据
func setDoctorsChedule0(date string, timing int, doctor_id int) {
	var users int
	models.DB.Raw("UPDATE doctorschedule SET state = 1, start_time = '09:00:00',end_time = '12:00:00', number = 25 WHERE doctor_id = ? AND data = ? AND timing = ?", doctor_id, date, timing).Scan(users)
}
func setDoctorsChedule1(date string, timing int, doctor_id int) {
	var users int
	models.DB.Raw("UPDATE doctorschedule SET state = 1, start_time = '14:00:00',end_time = '18:00:00',number = 25 WHERE doctor_id = ? AND data = ? AND timing = ?", doctor_id, date, timing).Scan(users)
}

// 检查当前科室是否有人值班
func findDoctorsChedule(date string, timing int, department int) int {
	var Data models.Doctor
	models.DB.Raw("SELECT d.doctor_id FROM doctors d JOIN ( SELECT * FROM doctorschedule WHERE data =? AND state = 1 AND timing = ?) s ON d.doctor_id = s.doctor_id WHERE department_id = ?", date, timing, department).First(&Data)
	return Data.DoctorId
}

// 获取所有科室id
func alldepartment() []int {
	var Data []int
	models.DB.Raw("SELECT department_id FROM departments").Scan(&Data)
	return Data
}

func deleteDoctorsChedule1() {
	result := models.DB.Exec("DELETE FROM doctorschedule WHERE state = ?", 0)
	if result.Error != nil {
		// 处理错误
	}
}

// 智能排班
func intelligentScheduling() {
	//获取后七天的日期,加今天一共八天
	var dataArr = futureData()
	// 检查后七日日期是否已有排班，如果没有，则进行创建
	for _, vale := range dataArr {
		examineDoctorsChedule(vale) //调用判断创建函数
	}

	//数据初始化成功

	// 检查当前科室是否有人值班
	var departmentArr = alldepartment() //departmentId获取所有科室id
	for _, data := range dataArr {
		for _, departmentId := range departmentArr {
			if findDoctorsChedule(data, 0, departmentId) == 0 { //上午是否有人值班
				// 获取当天对应科室前七天工时最少的医生，并返回其id
				var DoctorId = getDoctor(data, departmentId)
				// 修改当天医生的数据
				setDoctorsChedule0(data, 0, DoctorId)
				setDoctorsChedule1(data, 1, DoctorId)
			}
			if findDoctorsChedule(data, 1, departmentId) == 0 { //下午是否有人值班
				var DoctorId = getDoctor(data, departmentId)
				setDoctorsChedule1(data, 1, DoctorId)
			}
		}

	}
	deleteDoctorsChedule1() //删除多余的空白数据项
}

func futureData() []string {
	// 获取当前日期
	currentTime := time.Now()
	var dataArr []string
	dataArr = append(dataArr, currentTime.Format("2006-01-02"))
	// 获取后七天的日期,加今天一共八天
	for i := 1; i <= 7; i++ {
		// 使用 AddDate 方法获取后一天的日期
		nextDay := currentTime.AddDate(0, 0, i)
		// 存入后七天的日期
		dataArr = append(dataArr, nextDay.Format("2006-01-02"))
	}
	return dataArr
}

func (con AdminController) FindAllDoctorsChedule(c *gin.Context) {
	var Paging models.Paging
	var Data []models.DoctorsChedule
	if err := c.ShouldBindJSON(&Paging); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	skip := (Paging.Page - 1) * Paging.Limit

	table := models.DB.Raw("SELECT a.*,b.name,c.name as departments FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id ").Scan(&Data)
	var err = models.DB.Table("(?) as a ", table).Count(&Paging.Total).Limit(Paging.Limit).Offset(skip).Scan(&Data).Error

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		c.JSON(200, gin.H{"message": Data, "sda": &Paging.Total})
	}

	return

}

// IntelligentScheduling 智能排班
func (con AdminController) IntelligentScheduling(c *gin.Context) {

	//var total int64
	intelligentScheduling()

	c.JSON(200, gin.H{"message": "排班成功", "code": 200})

	return
}

//type DoctorsChedule struct {
//	ScheduleId int
//	DoctorId   int
//	Data       string
//	Timing     int
//	StartTime  string
//	EndTime    string
//	State      int
//	doctor     Doctor
//}

//func (DoctorsChedule) TableName() string {
//	return "doctorschedule"
//}

//type Doctor struct {
//	DoctorId     int    `json:"DoctorId"`
//	Username     string `json:"Username"`
//	Password     string `json:"Password"`
//	Name         string `json:"Name"`
//	DepartmentId int    `json:"DepartmentId"` //科室id
//	Department   string `json:"Department"`   //科室名称
//	Introduction string `json:"Introduction"` //简历描述
//
//}
//
//func (Doctor) TableName() string {
//	return "doctors"
//}
//
//type Department struct {
//	DepartmentId int    `json:"DepartmentId"`
//	Name         string `json:"Name"`
//	Class        string `json:"Class"` //内科，外科....等
//}

func (con AdminController) GetDoctorsChedule(c *gin.Context) {
	var Paging models.Paging
	var Data []models.DoctorsChedule
	if err := c.ShouldBindJSON(&Paging); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	skip := (Paging.Page - 1) * Paging.Limit
	models.DB.Raw("SELECT a.*,b.name,c.name as departments FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id   LIMIT ? OFFSET ?", Paging.Limit, skip).Scan(&Data)
	models.DB.Raw("SELECT a.*,b.name,c.name as departments FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id ").Count(&Paging.Total)

	c.JSON(200, gin.H{"message": Data, "Total": Paging.Total})
	return

}

// 进行排班的修改
func (con AdminController) SetDoctorsChedule(c *gin.Context) {
	//获取数据进行更新
	var updateData models.DoctorsChedule
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	fmt.Println("获取到的数据", updateData)
	// 执行更新操作  利用gorm可以实现动态sql的效果，若字段为空，gorm会自动忽略不进行修改
	if err := models.DB.Model(&models.DoctorsChedule{}).Where("schedule_id = ?", updateData.ScheduleId).Updates(&updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient data"})
		return
	}
	// 返回更新成功的响应
	c.JSON(200, gin.H{"message": "Patient data updated successfully", "code": 200})

}

// 进行排班表的查询
func (con AdminController) FinfDoctorsChedule(c *gin.Context) {
	var updateData models.DoctorsChedule
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

	if updateData.Name != "" && updateData.Data != "" {

		models.DB.Raw("SELECT a.*,b.name,c.name as departments ,a.number,b.price FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id  And b.name LIKE ? and a.data = ? LIMIT ? OFFSET ?", "%"+updateData.Name+"%", updateData.Data, Paging.Limit, skip).Scan(&Data)
		models.DB.Raw("SELECT a.*,b.name,c.name as departments FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id  And b.name LIKE ? and a.data = ? ", "%"+updateData.Name+"%", updateData.Data).Count(&Paging.Total)

		c.JSON(200, gin.H{"message": Data, "Total": Paging.Total})
		return

	} else if updateData.Name != "" && updateData.Data == "" {

		models.DB.Raw("SELECT a.*,b.name,c.name as departments,a.number,b.price FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id AND b.name LIKE ? LIMIT ? OFFSET ?", "%"+updateData.Name+"%", Paging.Limit, skip).Scan(&Data)
		models.DB.Raw("SELECT a.*,b.name,c.name as departments FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id AND b.name LIKE ? ", "%"+updateData.Name+"%").Count(&Paging.Total)

		c.JSON(200, gin.H{"message": Data, "Total": Paging.Total})
		return

	} else if updateData.Name == "" && updateData.Data != "" {

		models.DB.Raw("SELECT a.*,b.name,c.name as departments,a.number,b.price FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id AND a.data = ? LIMIT ? OFFSET ?", updateData.Data, Paging.Limit, skip).Scan(&Data)
		models.DB.Raw("SELECT a.*,b.name,c.name as departments FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id AND a.data = ? ", updateData.Data).Count(&Paging.Total)

		c.JSON(200, gin.H{"message": Data, "Total": Paging.Total})
		return

	} else if updateData.Name == "" && updateData.Data == "" {

		models.DB.Raw("SELECT a.*,b.name,c.name as departments,a.number,b.price FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id  order by a.data desc LIMIT ? OFFSET ?", Paging.Limit, skip).Scan(&Data)
		models.DB.Raw("SELECT a.*,b.name,c.name as departments FROM doctorschedule a,doctors b , departments c WHERE a.doctor_id = b.doctor_id AND b.department_id = c.department_id ").Count(&Paging.Total)

		c.JSON(200, gin.H{"message": Data, "Total": Paging.Total})
		return

	} else {
		c.JSON(201, gin.H{"message": "查询失败", "code": 201})
		return
	}

	return
}

// 进行排班表的删除
func (con AdminController) DeleteDoctorsChedule(c *gin.Context) {
	var scheduleId int
	if err := c.ShouldBindJSON(&scheduleId); err != nil {
		c.JSON(401, gin.H{"error": "Invalid data format"})
		return
	}
	result := models.DB.Exec("DELETE FROM doctorschedule WHERE schedule_id = ?", scheduleId)
	if result.Error != nil {
		// 处理错误
		c.JSON(201, gin.H{"message": "删除失败", "code": 201})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功", "code": 200})
	return
}

// 进行批量排班表的删除
func (con AdminController) DeleteDoctorsChedules(c *gin.Context) {
	var scheduleId []int
	if err := c.ShouldBindJSON(&scheduleId); err != nil {
		c.JSON(401, gin.H{"error": "Invalid data format"})
		return
	}
	for _, id := range scheduleId {
		result := models.DB.Exec("DELETE FROM doctorschedule WHERE schedule_id = ?", id)
		if result.Error != nil {
			// 处理错误
			c.JSON(201, gin.H{"message": "删除失败", "code": 201})
			return
		}
	}

	c.JSON(200, gin.H{"message": "批量删除成功", "code": 200})
	return
}

// 进行排班表的插入
func (con AdminController) AddDoctorsChedule(c *gin.Context) {

	var Data models.DoctorsChedule
	var doctor_id int
	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	fmt.Println(Data)

	models.DB.Raw("SELECT doctor_id  FROM doctors  WHERE  name = ? ", Data.Name).First(&doctor_id)
	if doctor_id == 0 {
		c.JSON(201, gin.H{"message": "医生不存在", "code": 201})
		return
	}
	if err := models.DB.Exec("INSERT INTO doctorschedule (doctor_id, data, timing,start_time,end_time, state,number) VALUES (?, ?, ?,?,?,?,?)", doctor_id, Data.Data, Data.Timing, Data.StartTime, Data.EndTime, 1, 25).Error; err != nil {
		c.JSON(201, gin.H{"message": "添加失败", "code": 201})
		return
	}
	c.JSON(201, gin.H{"message": "添加成功", "code": 200})
	return
}

// 获取用户列表
func (con AdminController) PatientList(c *gin.Context) {
	var Datas models.Patient
	if err := c.ShouldBindBodyWith(&Datas, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	var Paging models.Paging
	if err := c.ShouldBindBodyWith(&Paging, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	var Data []models.Patient
	skip := (Paging.Page - 1) * Paging.Limit

	if Datas.Name != "" { //根据姓名进行查

		models.DB.Raw("SELECT * FROM patient WHERE `name` LIKE ?", "%"+Datas.Name+"%").Count(&Paging.Total)
		if err := models.DB.Raw("SELECT * FROM patient WHERE `name` LIKE ? LIMIT ? OFFSET ?", "%"+Datas.Name+"%", Paging.Limit, skip).Scan(&Data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
			return
		}
	} else { //获取全部
		models.DB.Raw("SELECT * FROM patient").Count(&Paging.Total)
		if err := models.DB.Raw("SELECT * FROM patient  LIMIT ? OFFSET ?", Paging.Limit, skip).Scan(&Data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
			return
		}
	}
	c.JSON(200, gin.H{"message": Data, "code": 200, "Total": Paging.Total})
	return
}

// 删除用户
func (con AdminController) DeletePatient(c *gin.Context) {
	var Datas models.Patient
	if err := c.ShouldBindBodyWith(&Datas, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	result := models.DB.Exec("DELETE FROM patient WHERE patient_id = ?", Datas.PatientId)
	if result.Error != nil {
		// 处理错误
		c.JSON(201, gin.H{"message": "删除失败", "code": 201})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功", "code": 200})
	return
}
