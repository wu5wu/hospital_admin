package models

// DoctorsChedule 排班表
type DoctorsChedule struct {
	ScheduleId   int
	DoctorId     int
	Data         string
	Timing       int
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Name         string `json:"Name"`
	Introduction string `json:"Introduction"` //简历描述
	Departments  string `json:"Departments"`
	DepartmentId int    `json:"DepartmentId"`
	Class        string `json:"Class"`
	Price        int    `json:"Price"`
	Number       int    `json:"Number"`
	Img          string
}

func (DoctorsChedule) TableName() string {
	return "doctorschedule"
}
