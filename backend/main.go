package main

import (
	"os"

	middleware "revise-it/backend/middleware"
	routes "revise-it/backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8006"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())
	routes.HealthcheckRoutes(router)
	routes.DeckRoutes(router)
	// put routes here that dont require authentication
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	// put routes here that require authentication
	router.Run(":" + port)
}
