package routes

import (
	controller "coffend/backend/controllers"
	"coffend/backend/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", middleware.AuthRequired(), middleware.AdminOnly(), controller.GetAllOrders())
	incomingRoutes.GET("/orders/:order_id", middleware.AuthRequired(), controller.GetOrderByID())
	incomingRoutes.POST("/orders", middleware.AuthRequired(), controller.PostOrder())
	incomingRoutes.PATCH("/orders/:order_id", middleware.AuthRequired(), controller.PatchOrderByID())
	incomingRoutes.DELETE("/orders/:order_id", middleware.AuthRequired(), controller.DeleteOrderByID())
}
