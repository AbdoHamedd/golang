package Routes

import (
	"github.com/gin-gonic/gin/ginS"
	"project1/Controller/Visitors"
)

func (app *RouterApp) visitorsRoutes() {
	ginS.GET("Register", Visitors.Register)
}
