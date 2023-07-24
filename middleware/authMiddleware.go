package middleware

import (
	"go-jwt/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "token" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("firstName", claims.FirstName)
		c.Set("lastName", claims.LastName)
		c.Set("uid", claims.Uid)
		c.Set("userType", claims.UserType)
		c.Next()
	}
}
