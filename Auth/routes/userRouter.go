package routes

import (
	"github.com/Auth/controllers"
	"github.com/Auth/middleware"
	"github.com/gin-gonic/gin"
)

//after login the user will have a token because of which we are using
//protected routes and using this middleware layer
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
