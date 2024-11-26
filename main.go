package main

import (
	"1/Application"
	"1/Models"
	"1/Routes"
)

func main() {

	//git app
	app := Application.NewApp()

	//Migrate Project
	app.DB.AutoMigrate(&Models.User{})

	//Close Connection App
	Application.CloseDatabaseConnection(&app)

	//Start routing
	routerApp := Routes.RouterApp{&app}
	routerApp.Routing()

	//Start Server app
	app.Gin.Run()
}
