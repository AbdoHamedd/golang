package Routes

import "project1/Application"

type RouterApp struct {
	*Application.Application
}

func (app *RouterApp) Routing() {
	app.visitorsRoutes()
	app.authsRoutes()
	app.adminsRoutes()
}
