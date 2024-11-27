package Application

import (
	"database/sql"
	"fmt"
	"github.com/bykovme/gotrans"
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
		err := gotrans.InitLocales("./Public/Lang")
		if err != nil {
			fmt.Println("Error in Load Trans Files", err.Error())
		}
		return &application
	}
}

func NewApp() *Application {
	app := App()
	return app()
}

func (req *Application) Share() {

}
