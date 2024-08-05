package models

type Department struct {
	DepartmentId int    `json:"DepartmentId"`
	Name         string `json:"Name"`
	Class        string `json:"Class"` //内科，外科....等
}
