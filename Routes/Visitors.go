package Routes

import "1/Controller/Visitors"

func (app *RouterApp) visitorsRoutes() {
	app.Gin.GET("/create-user", Visitors.CreateUser)
}
