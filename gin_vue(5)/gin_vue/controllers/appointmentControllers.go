package controllers

import (
	"fmt"
	"gin_vue/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type AppointmentControllers struct {
}

type deal struct {
	Balance int
	Price   int
	Status  string
	Number  int
	Data    string
}

// AddDepartment 进行预约的增加
func (con AppointmentControllers) AddAppointment(c *gin.Context) {

	var deals deal
	//var deals2 deal
	var updateData models.Appointment
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	updateData.Status = "已预约"
	// 执行插入操作
	//查询余额和价格数据传入结构体
	models.DB.Raw("SELECT balance FROM patient WHERE patient_id = ?", updateData.PatientId).Scan(&deals)
	models.DB.Raw("SELECT b.price FROM doctorschedule a, doctors b WHERE a.doctor_id = b.doctor_id AND a.schedule_id = ?", updateData.ScheduleId).Scan(&deals)

	models.DB.Raw("SELECT number FROM doctorschedule WHERE schedule_id = ?", updateData.ScheduleId).Scan(&deals)
	if deals.Number <= 0 {
		c.JSON(200, gin.H{"message": "预约已满", "code": 201})
		return
	}

	if deals.Balance >= deals.Price {
		var aaa int
		balance := deals.Balance - deals.Price

		//更新余额
		models.DB.Raw("UPDATE patient SET balance = ? WHERE patient_id = ?", balance, updateData.PatientId).Scan(aaa)
		//更新剩余可预约人数
		models.DB.Raw("UPDATE doctorschedule SET number = ? WHERE schedule_id = ?", deals.Number-1, updateData.ScheduleId).Scan(aaa)
		if err := models.DB.Exec("INSERT INTO appointment (patient_id, schedule_id,status,created) VALUES (?, ?,?,?)", updateData.PatientId, updateData.ScheduleId, updateData.Status, updateData.Created).Error; err != nil {
			c.JSON(200, gin.H{"message": err, "code": 201})
			return
		}

		c.JSON(200, gin.H{"message": "预约成功", "code": 200, "balance": balance})
		return
	} else if deals.Balance < deals.Price {
		c.JSON(200, gin.H{"message": "余额不足", "code": 201})
		return
	} else {
		c.JSON(200, gin.H{"message": "有问题", "code": 201})
		return
	}
	// 返回成功响应
	c.JSON(200, gin.H{"message": "预约 successfully", "code": 200})
}

// 获取预约列表
func (con AppointmentControllers) GetAppointment(c *gin.Context) {
	var updateData models.Patient
	var appointmentData []models.Appointment

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

	if updateData.PatientId != 0 {
		models.DB.Raw("SELECT a.appointment_id,a.created,b.`data`,b.start_time,b.end_time,c.`name`,d.`name` as departments ,a.status FROM appointment a , doctorschedule b ,doctors c,departments d  WHERE a.patient_id = ? AND a.schedule_id = b.schedule_id AND c.doctor_id = b.doctor_id AND d.department_id = c.department_id", updateData.PatientId).Count(&Paging.Total)
		models.DB.Raw("SELECT a.appointment_id,a.created,b.`data`,b.start_time,b.end_time,c.`name`,d.`name` as departments ,a.status FROM appointment a , doctorschedule b ,doctors c,departments d  WHERE a.patient_id = ? AND a.schedule_id = b.schedule_id AND c.doctor_id = b.doctor_id AND d.department_id = c.department_id order by a.created desc  LIMIT ? OFFSET ?", updateData.PatientId, Paging.Limit, skip).Scan(&appointmentData)

	} else {
		models.DB.Raw("SELECT a.appointment_id,a.created,b.`data`,b.start_time,b.end_time,c.`name`,d.`name` as departments ,a.status FROM appointment a , doctorschedule b ,doctors c,departments d WHERE  a.schedule_id = b.schedule_id AND c.doctor_id = b.doctor_id AND d.department_id = c.department_id").Count(&Paging.Total)
		models.DB.Raw("SELECT a.appointment_id,a.created,b.`data`,b.start_time,b.end_time,c.`name`,d.`name` as departments ,a.status FROM appointment a , doctorschedule b ,doctors c,departments d WHERE  a.schedule_id = b.schedule_id AND c.doctor_id = b.doctor_id AND d.department_id = c.department_id order by a.created desc LIMIT ? OFFSET ?", Paging.Limit, skip).Scan(&appointmentData)

	}
	c.JSON(200, gin.H{"message": appointmentData, "code": 200, "Total": Paging.Total})

}

// 删除预约表数据
func (con AppointmentControllers) DeleteAppointment(c *gin.Context) {
	var Id models.Appointment

	if err := c.ShouldBindJSON(&Id); err != nil {
		c.JSON(401, gin.H{"error": "Invalid data format"})
		return
	}
	result := models.DB.Exec("DELETE FROM appointment WHERE appointment_id = ?", Id.AppointmentId)
	if result.Error != nil {
		// 处理错误
		c.JSON(201, gin.H{"message": "删除失败", "code": 201})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功", "code": 200})
	return
}

// 取消预约（修改状态）
func (con AppointmentControllers) SetAppointment(c *gin.Context) {
	var deals deal
	var data models.Appointment

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(401, gin.H{"error": "Invalid data format"})
		return
	}

	models.DB.Raw("SELECT b.`data` FROM appointment a ,doctorschedule b WHERE a.schedule_id = b.schedule_id and a.appointment_id=?", data.AppointmentId).Scan(&deals)

	models.DB.Raw("SELECT `status` FROM `appointment` WHERE appointment_id = ?", data.AppointmentId).Scan(&deals)
	fmt.Println(deals.Status)
	if deals.Status == "已预约" {

		if data.Status == "已签到" && deals.Data == models.GetDate() {
			result := models.DB.Exec("UPDATE appointment SET `status` = ? ,sign_time = ?  WHERE appointment_id = ?", data.Status, data.SignTime, data.AppointmentId)
			if result.Error != nil {
				// 处理错误
				c.JSON(201, gin.H{"message": "修改失败", "code": 201})
				return
			}
			c.JSON(200, gin.H{"message": "签到成功", "code": 200})
			return
		} else if data.Status == "已签到" && deals.Data != models.GetDate() {
			c.JSON(201, gin.H{"message": "当天无预约", "code": 201})
			return
		}

		result := models.DB.Exec("UPDATE appointment SET `status` = ?  WHERE appointment_id = ?", data.Status, data.AppointmentId)
		if result.Error != nil {
			// 处理错误
			c.JSON(201, gin.H{"message": "修改失败", "code": 201})
			return
		}
		models.DB.Raw("SELECT balance FROM patient WHERE patient_id = ?", data.PatientId).Scan(&deals)
		models.DB.Raw("SELECT b.price FROM doctorschedule a, doctors b,appointment c WHERE a.doctor_id = b.doctor_id AND a.schedule_id = c.schedule_id AND c.appointment_id = ?", data.AppointmentId).Scan(&deals)
		var Number int
		balance := deals.Balance + deals.Price
		models.DB.Raw("UPDATE patient SET balance = ? WHERE patient_id = ?", balance, data.PatientId).Scan(Number)
		models.DB.Raw("SELECT b.number FROM appointment a , doctorschedule b  WHERE a.schedule_id = b.schedule_id AND a.appointment_id = ?", data.AppointmentId).Scan(&deals)

		result = models.DB.Exec("UPDATE doctorschedule a , appointment b SET a.number = ? WHERE a.schedule_id = b.schedule_id AND b.appointment_id = ?", deals.Number+1, data.AppointmentId)
		if result.Error != nil {
			// 处理错误
			c.JSON(201, gin.H{"message": "修改失败", "code": 201})
			return
		}
		c.JSON(200, gin.H{"message": "取消成功", "code": 200, "balance": balance})
		return

	} else if deals.Status == "已取消" {
		c.JSON(200, gin.H{"message": "已取消的预约，不能进行操作", "code": 201})
		return
	} else if deals.Status == "已签到" {
		if data.Status == "已签到" {
			c.JSON(200, gin.H{"message": "请不要重复签到", "code": 201})
		} else {
			c.JSON(200, gin.H{"message": "已签到，不能取消", "code": 201})
		}

		return
	} else if deals.Status == "已完成" {
		c.JSON(200, gin.H{"message": "已完成，不能进行操作", "code": 201})
		return
	} else {
		c.JSON(200, gin.H{"message": "当前状态有问题", "code": 201})
	}

}

func (con AppointmentControllers) GetDate(c *gin.Context) {

	c.JSON(200, gin.H{"message": futureData(), "code": 200})
	return
}

// 获取医生值班表
func (con AppointmentControllers) GetDoctorsChedule(c *gin.Context) {
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

	if updateData.DepartmentId != 0 && updateData.Data != "" { //按日期和科室进行查

		models.DB.Raw("SELECT c.schedule_id,c.`data`,d.`name`,d.introduction,c.timing,c.start_time,c.end_time,a.class,a.`name` as departments ,d.price ,c.number,d.img FROM departments a, doctors d,doctorschedule c WHERE c.data = ? AND d.doctor_id = c.doctor_id AND d.department_id = ? AND a.department_id = d.department_id", updateData.Data, updateData.DepartmentId).Scan(&Data)

		c.JSON(200, gin.H{"message": Data, "Total": Paging.Total})
		return

	} else if updateData.DepartmentId != 0 && updateData.Data == "" { //按科室进行查

		models.DB.Raw("SELECT c.schedule_id,c.`data`,d.`name`,d.introduction,c.timing,c.start_time,c.end_time,a.class,a.`name` as departments,d.price ,c.number,d.img FROM departments a, doctors d,doctorschedule c WHERE d.doctor_id = c.doctor_id AND d.department_id = ? AND a.department_id = d.department_id LIMIT ? OFFSET ?", updateData.DepartmentId, Paging.Limit, skip).Scan(&Data)
		models.DB.Raw("SELECT c.schedule_id,d.`name`,d.introduction,c.timing,c.start_time,c.end_time,a.class,a.`name` FROM departments a, doctors d,doctorschedule c WHERE d.doctor_id = c.doctor_id AND d.department_id = ? AND a.department_id = d.department_id", updateData.DepartmentId).Count(&Paging.Total)

		c.JSON(200, gin.H{"message": Data, "Total": Paging.Total})
		return

	} else if updateData.DepartmentId == 0 && updateData.Data != "" { //按日期查询

		models.DB.Raw("SELECT c.schedule_id,c.`data`,d.`name`,d.introduction,c.timing,c.start_time,c.end_time,a.class,a.`name` as departments,d.price,c.number ,d.img FROM departments a, doctors d,doctorschedule c WHERE c.data = ? AND d.doctor_id = c.doctor_id AND a.department_id = d.department_id LIMIT ? OFFSET ?", updateData.Data, Paging.Limit, skip).Scan(&Data)
		models.DB.Raw("SELECT c.schedule_id,c.`data`,d.`name`,d.introduction,c.timing,c.start_time,c.end_time,a.class,a.`name` FROM departments a, doctors d,doctorschedule c WHERE c.data = ? AND d.doctor_id = c.doctor_id AND a.department_id = d.department_id ", updateData.Data).Count(&Paging.Total)

		c.JSON(200, gin.H{"message": Data, "Total": Paging.Total})
		return

	} else {
		c.JSON(201, gin.H{"message": "查询失败", "code": 201})
		return
	}

	return
}
