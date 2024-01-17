package routes

import (
	controller "revise-it/backend/controllers"

	"github.com/gin-gonic/gin"
)

func DeckRoutes(incomingRoutes *gin.Engine) {
	//add deck route
	incomingRoutes.POST("/add-deck", controller.AddDeck())
}
