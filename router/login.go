package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petersonsalme/rest-api-with-jwt/middleware"
	"github.com/petersonsalme/rest-api-with-jwt/model"
	"github.com/petersonsalme/rest-api-with-jwt/redis"
)

var user = model.User{
	ID:       1,
	Username: "username",
	Password: "password",
}

// Login should verify user credentials
func Login(c *gin.Context) {
	var u model.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	token, err := middleware.CreateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := redis.CreateAuth(user.ID, token)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}

// Logout Logout
func Logout(c *gin.Context) {
	au, err := middleware.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	deleted, delErr := redis.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	c.JSON(http.StatusOK, "Successfully logged out")
}
