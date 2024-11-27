package main

import (
	"fmt"
	"github.com/bykovme/gotrans"
	"project1/Application"
	"project1/Models"
	"project1/Routes"
)

func main() {

	gotrans.InitLocales("./Public/Lang") //  Path to the folder with localization files
	gotrans.SetDefaultLocale("re")       // Setting default locale

	fmt.Println(gotrans.T("hello_world"))

	//git app/                  (1)
	app := Application.NewApp()

	//Migrate Project
	app.DB.AutoMigrate(&Models.User{})

	//Close Connection App
	Application.CloseDatabaseConnection(app)

	//Start routing
	routerApp := Routes.RouterApp{app}
	routerApp.Routing()

	//Start Server app
	app.Gin.Run()
}
