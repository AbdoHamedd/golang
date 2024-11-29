package Auths

import (
	"github.com/gin-gonic/gin"
	"project1/Application"
	Transformer "project1/Transformers/Visitors"
)

func Me(c *gin.Context) {
	request, auth := Application.AuthRequest(c)
	if !auth {

		return
	}
	request.Ok(Transformer.UserTransformer(*request.User))
}
