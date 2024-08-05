package models

type Application struct {
	ApplicationId int
	DoctorId      int
	Data          string
	State         string
	Reason        string
	CTime         string
}

func (Application) TableName() string {
	return "application"
}
