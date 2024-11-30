package main

import (
	"project1/Application"
	"project1/Routes"
)

func main() {

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
