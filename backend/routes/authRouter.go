package routes

import (
	controller "coffend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register", controller.Register())
	incomingRoutes.POST("/login", controller.Login())
}
