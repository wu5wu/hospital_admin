package models

type Appointment struct {
	AppointmentId int    `json:"appointment_id"`
	PatientId     int    `json:"patient_id"`
	ScheduleId    int    `json:"schedule_id"`
	Status        string `json:"status"`
	Created       string `json:"created"`
	Data          string
	StartTime     string
	EndTime       string
	Name          string
	Departments   string
	SignTime      string
	Number        int
}

func (Appointment) TableName() string {
	return "appointment"
}
