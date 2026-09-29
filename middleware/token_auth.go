package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware TokenAuthMiddleware
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
