package main

import (
	// "fmt"

	"net/http"
	"os"

	"github.com/binoymanoj/jwt-auth-go/controllers"
	"github.com/binoymanoj/jwt-auth-go/initializers"
	"github.com/binoymanoj/jwt-auth-go/middleware"
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
			"message": "Fireworks Ecom API",
		})
	})
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// grouping API with "/api" prefix route
	api := r.Group("/api")
	{
		// grouping API with "/api/auth" prefix route
		auth := api.Group("/auth")
		{
			auth.POST("/signup", controllers.SignUp)
			auth.POST("/login", controllers.Login)
			auth.GET("/validate", middleware.RequireAuth, controllers.Validate)
		}
	}

	// r.Run(":4000")

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	r.Run(":" + port)
}
