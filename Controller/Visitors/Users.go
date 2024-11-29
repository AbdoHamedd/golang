package Visitors

import (
	"github.com/gin-gonic/gin"
	"project1/Application"
	"project1/Models"
	Transformer "project1/Transformers/Visitors"
	"project1/Validations/Visitors"
)

func Register(c *gin.Context) {
	r := Application.NewRequest(c)
	var User Models.User
	//binding
	r.Context.ShouldBind(&User)
	//validate the requeset
	if r.ValidateRequest(Visitors.RegisterValidation(User)).Fails() {
		return
	}
	User.Token = User.Email
	User.Group = "user"
	r.DB.Create(&User)
	r.Created(Transformer.UserTransformer(User))
}

func Login(c *gin.Context) {
	r := Application.NewRequest(c)
	var User Models.User
	r.Context.ShouldBind(&User)
	if r.ValidateRequest(Visitors.LoginValidation(User)).Fails() {
		return
	}
	r.DB.Where("email = ?", User.Email).Where("password = ?", User.Password).First(&User)
	if User.ID == 0 {
		r.UserNoutFound()
		return
	}
	r.Ok(Transformer.UserTransformer(User))
}
