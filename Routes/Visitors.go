package Routes

import (
	"project1/Controller/Visitors"
)

func (app *RouterApp) visitorsRoutes() {
	app.Gin.POST("Register", Visitors.Register)
	app.Gin.POST("Login", Visitors.Login)
}
