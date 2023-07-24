package routes

import (
	"go-jwt/controllers"
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.Use(middleware.Authenticate())
	incommingRoutes.GET("/users", controllers.GetUsers())
	incommingRoutes.GET("/users/:user_id", controllers.GetUser())
}
