package routes

import (
	controller "coffend/backend/controllers"
	"coffend/backend/middleware"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", middleware.AuthRequired(), controller.GetAllTables())
	incomingRoutes.GET("/tables/:table_id", middleware.AuthRequired(), controller.GetTableByID())
	incomingRoutes.POST("/tables", middleware.AuthRequired(), middleware.AdminOnly(), controller.PostTable())
	incomingRoutes.PATCH("/tables/:table_id", middleware.AuthRequired(), middleware.AdminOnly(), controller.PatchTableByID())
	incomingRoutes.DELETE("/tables/:table_id", middleware.AuthRequired(), middleware.AdminOnly(), controller.DeleteTableByID())
}
