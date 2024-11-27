package Visitors

import (
	"github.com/gin-gonic/gin"
	"project1/Application"
	"project1/Models"
)

func Register(c *gin.Context) {
	r := Application.NewREQUESTwithAUTN(c)
	var User Models.User
	r.Context.ShouldBind(&User)

}
