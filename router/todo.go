package router

import (
	"net/http"

	"github.com/petersonsalme/rest-api-with-jwt/model"

	"github.com/gin-gonic/gin"
	"github.com/petersonsalme/rest-api-with-jwt/middleware"
	"github.com/petersonsalme/rest-api-with-jwt/redis"
)

// CreateTodo CreateTodo
func CreateTodo(c *gin.Context) {
	var td *model.Todo
	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	tokenAuth, err := middleware.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	userID, err := redis.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	td.UserID = userID

	c.JSON(http.StatusCreated, td)
}
