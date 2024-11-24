package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	app := app()
	application := app()

	application.Gin.GET("/ping", func(c *gin.Context) {
		request := newRequest(c)

		request.Ok(gin.H{
			"message": "Hi All",
		})

		//request.Response(200, gin.H{
		//	"message": "Hi All",
		//})

		//request.context.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
	})

	application.Gin.Run()
}
