package Auths

import (
	"github.com/gin-gonic/gin"
	"project1/Application"
)

func CreateUser(c *gin.Context) {
	request, auth := Application.AuthRequest(c)
	if !auth {
		request.NotAuth()
		return
	}
	request.Ok(request.User)
}
