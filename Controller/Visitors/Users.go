package Visitors

import (
	"1/Application"
	"1/Models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	request := Application.NewRequest(c)
	//API To Create a User
	user := Models.User{
		UserName: "Abdelaziz Hamed",
		Email:    "AbdelazizHamed@gmail.com",
		Password: "0123456",
	}
	request.DB.Create(&user)
	request.Created(user)

}
