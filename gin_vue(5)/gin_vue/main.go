package main

import (
	"gin_vue/middleware"
	"gin_vue/ruoters"
	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int
	Name string
	Pwd  string
}

func (User) TableName() string {
	return "user"
}



func main() {
	r := gin.Default()
	r.Use(middleware.LogMiddleWare())
	r.LoadHTMLGlob("templates/*")
	ruoters.IndexRuoteraInit(r)
	ruoters.LoginRuotersInit(r)
	ruoters.DoctorRuotersInit(r)
	ruoters.AdminRuotersInit(r)
	ruoters.AppointmentRuotersInit(r)

	r.Run()
}
