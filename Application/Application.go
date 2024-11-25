package Application

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

func NewApp() Application {
	app := App()
	aplication := app()
	return aplication
}

func (req *Application) Share() {

}

func App() func() Application {
	return func() Application {
		var application Application
		application.Gin = gin.Default()
		connectionToDataBase(&application)
		return application

	}
}
