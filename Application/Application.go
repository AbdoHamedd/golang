package Application

import (
	"database/sql"
	"fmt"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"project1/Database"
	"project1/Database/Seeders"
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
		//1-Default Gin
		application.Gin = gin.Default()
		//2-Connect To DataBase
		connectionToDataBase(&application)
		//3-Load Trans;ate
		err := gotrans.InitLocales("./Public/Lang")
		if err != nil {
			fmt.Println("Error in Load Trans Files", err.Error())
		}
		//4-Return Application
		return &application
	}
}

// Need This For Handle Closure
func NewApp() *Application {
	app := App()
	return app()
}

// Need This Function To COnnect and Close connection with database
func (req *Application) Share() {

}

// For Migrate
func (app *Application) Migrate() {
	Database.Migrate(app.DB)
}

// For Seed
func (app *Application) Seed() {
	Seeders.Seeds(app.DB)
}
