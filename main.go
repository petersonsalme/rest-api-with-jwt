package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/petersonsalme/rest-api-with-jwt/middleware"

	"github.com/petersonsalme/rest-api-with-jwt/redis"
	"github.com/petersonsalme/rest-api-with-jwt/router"

	"github.com/gin-gonic/gin"
)

var routerEngine *gin.Engine

func init() {
	// .env is optional: variables can also come from the environment
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	redis.Connect()

	routerEngine = gin.Default()
}

func main() {
	routerEngine.POST("/login", router.Login)
	routerEngine.POST("/logout", middleware.TokenAuthMiddleware(), router.Logout)
	routerEngine.POST("/token/refresh", middleware.Refresh)

	routerEngine.POST("/todo", middleware.TokenAuthMiddleware(), router.CreateTodo)

	log.Fatal(routerEngine.Run(":8080"))
}
