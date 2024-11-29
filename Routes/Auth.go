package Routes

import "project1/Controller/Auths"

func (app *RouterApp) authsRoutes() {
	app.Gin.GET("/me", Auths.Me)

}
