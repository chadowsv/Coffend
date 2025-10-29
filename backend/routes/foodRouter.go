package routes

import (
	controller "coffend/backend/controllers"
	"coffend/backend/middleware"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetAllFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFoodByID())
	incomingRoutes.POST("/foods", middleware.AuthRequired(), middleware.AdminOnly(), controller.PostFood())
	incomingRoutes.PATCH("/foods/:food_id", middleware.AuthRequired(), middleware.AdminOnly(), controller.PatchFood())
	incomingRoutes.DELETE("/foods/:food_id", middleware.AuthRequired(), middleware.AdminOnly(), controller.DeleteFoodByID())
}
