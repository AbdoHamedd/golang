package Visitors

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"project1/Application"
	"project1/Models"
)

func Register(c *gin.Context) {
	r := Application.NewREQUESTwithAUTN(c)
	var User Models.User
	//binding
	r.Context.ShouldBind(&User)
	err := validation.Errors{
		"name":     validation.Validate(User.UserName, validation.Required, validation.Length(2, 20)),
		"email":    validation.Validate(User.Email, validation.Required, is.Email, validation.Length(5, 20)),
		"password": validation.Validate(User.Password, validation.Required, validation.Length(5, 20)),
	}.Filter()
	if err != nil {
		r.BadRequest(err)
		return
	}
	r.Ok(User)
}
