package models

// Doctor 医生表
type Doctor struct {
	DoctorId     int    `json:"DoctorId"`
	Username     string `json:"Username"`
	Password     string `json:"Password"`
	Name         string `json:"Name"`
	DepartmentId int    `json:"DepartmentId"` //科室id
	Department   string `json:"Department"`   //科室名称
	Introduction string `json:"Introduction"` //简历描述
	Price        int    `json:"Price"`        //预约价格
	Img          string `json:"Img"`
	State        string //医生状态
}

func (Doctor) TableName() string {
	return "doctors"
}
