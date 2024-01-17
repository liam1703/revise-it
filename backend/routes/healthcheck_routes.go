package routes

import (
	controller "revise-it/backend/controllers"

	"github.com/gin-gonic/gin"
)

func HealthcheckRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/healthcheck", controller.Health())
}
