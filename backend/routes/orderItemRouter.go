package routes

import (
	controller "coffend/backend/controllers"
	"coffend/backend/middleware"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", middleware.AuthRequired(), middleware.AdminOnly(), controller.GetOrderItems())
	incomingRoutes.GET("/orderItems-order/:order_id", middleware.AuthRequired(), controller.GetOrderItemsByOrder())
	incomingRoutes.GET("/orderItems/:order_item_id", middleware.AuthRequired(), controller.GetOrderItemByID())
	incomingRoutes.POST("/orderItems", middleware.AuthRequired(), controller.PostOrderItem())
	incomingRoutes.PATCH("/orderItems/:order_item_id", middleware.AuthRequired(), controller.PatchOrderItemByID())
	incomingRoutes.DELETE("/orderItems/:order_item_id", middleware.AuthRequired(), controller.DeleteOrderItemByID())
}
