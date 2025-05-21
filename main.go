package main

import (
	// "fmt"

	"github.com/binoymanoj/jwt-auth-go/controllers"
	"github.com/binoymanoj/jwt-auth-go/initializers"
	"github.com/binoymanoj/jwt-auth-go/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "JWT-Auth-GO's index page",
		})
	})

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// r.Run(":4000")

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	r.Run(":" + port)
}
