package main

import (
	"coffend/backend/database"
	"coffend/backend/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Servidor funcionando correctamente",
		})
	})

	router.Run(":" + port)
}
