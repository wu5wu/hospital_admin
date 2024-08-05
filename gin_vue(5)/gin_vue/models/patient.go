package models

// Patient 患者表
type Patient struct {
	PatientId     int
	Username      string
	Password      string
	Name          string
	Identity      string
	Img           string
	Balance       int
	Phone         string
	AppointmentId int
}

func (Patient) TableName() string {
	return "patient"
}
