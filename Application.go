package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	Gin        *gin.Engine
	DB         *gorm.DB
	connection *sql.DB
}

func (req *Application) Share() {

}

func app() func() Application {
	return func() Application {
		var application Application
		application.Gin = gin.Default()
		connectionToDataBase(&application)
		return application

	}
}
