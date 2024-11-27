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

// ( 1)
func App() func() *Application {
	return func() *Application {
		var application Application
		application.Gin = gin.Default()
		connectionToDataBase(&application)
		return &application

	}
}

func NewApp() *Application {
	app := App()
	return app()
}

func (req *Application) Share() {

}
