package routes

import (
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.Use(middleware.Authenticate())

}
