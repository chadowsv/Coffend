package routes

import (
	controller "coffend/controllers"

	"github.com/gin-gonic/gin"
)

// Controller para usar funciones que están en la carpeta controllers
// Usar framework Gin que maneja las rutas y peticiones HTTP
func FoodRoutes(incomingRoutes *gin.Engine) {
	//Recibe el parámetro incomingRoutes de tipo *gin.Engine
	//Obtención de todos los platos
	incomingRoutes.GET("/foods", controller.GetFoods())
	//Obtención de un plato específico
	incomingRoutes.GET("/foods/:food_id", controller.GetFood())
	//Creación de un nuevo plato
	incomingRoutes.POST("/foods", controller.CreateFood())
	//Actualización parcial de un plato existente
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
	//Eliminacion de plato en base a food_id
	incomingRoutes.DELETE("/foods/:food_id", controller.DeleteFood())
}
