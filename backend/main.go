// main.go
// file currently useless was using it to begin with will use it again at end of users work
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Connect to database
	// configs.ConnectDB()

	// Start the server
	router.Run(":8080")
}
