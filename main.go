package main

import (
	"1/Application"
	"github.com/gin-gonic/gin"
)

func main() {

	app := Application.NewApp()

	app.Gin.GET("/ping", func(c *gin.Context) {
		request := Application.NewRequest(c)

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

	app.Gin.Run()
}
