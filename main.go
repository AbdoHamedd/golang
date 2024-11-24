package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	app := app()
	application := app()

	application.Gin.GET("/ping", func(c *gin.Context) {
		request := newRequest(c)

		request.closeDatabaseConnection()
		request.context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	application.Gin.Run()
}
