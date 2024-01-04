// main.go
// file currently useless was using it to begin with will use it again at end of users work
package main

import (
	"net/http"

	"revise-it/backend/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Connect to database
	configs.ConnectDB()

	//Define your routes and handlers here
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Flashcard App!",
		})
	})

	// Start the server
	router.Run(":8080")
}
