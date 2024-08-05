package models

type Admin struct {
	AdminId  int    `json:"AdminId"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}
