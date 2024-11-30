package main

import (
	"github.com/subosito/gotenv"
	"project1/Application"
	"project1/Routes"
)

func main() {

	gotenv.Load(".env")
	//git app/                  (1)
	app := Application.NewApp()

	//Migrate Project
	app.Migrate()

	//Seed Data
	app.Seed()

	//Close Connection App
	Application.CloseDatabaseConnection(app)

	//Start routing
	routerApp := Routes.RouterApp{app}
	routerApp.Routing()

	//Start Server app
	app.Gin.Run()
}
