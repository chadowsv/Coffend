package routes

import (
	controller "coffend/backend/controllers"
	"coffend/backend/middleware"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetAllMenus())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenuByID())
	incomingRoutes.POST("/menus", middleware.AuthRequired(), middleware.AdminOnly(), controller.PostMenu())
	incomingRoutes.PATCH("/menus/:menu_id", middleware.AuthRequired(), middleware.AdminOnly(), controller.UpdateMenuByID())
	incomingRoutes.DELETE("/menus/:menu_id", middleware.AuthRequired(), middleware.AdminOnly(), controller.DeleteMenuByID())
}
