package Visitors

import (
	"github.com/gin-gonic/gin"
	"project1/Application"
	"project1/Models"
	"project1/Validations/Visitors"
)

func Register(c *gin.Context) {
	r := Application.NewRequest(c)
	var User Models.User
	//binding
	r.Context.ShouldBind(&User)
	//validate the requeset
	r.ValidateRequest(Visitors.RegisterValidation(User))
	if r.Fails() {
		return
	}
	User.Token = User.Email
	User.Group = "user"
	r.DB.Create(&User)
	r.Created(User)
}
