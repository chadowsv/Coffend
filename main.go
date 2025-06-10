package main

import (
	"coffend/database"
	"coffend/routes"
	"os"

	//Framework web usado para creacion del servidor y manejo de rutas
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	//Intento de obtencion del numero de puerto, pero por defecto sera 8000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	//Creacion de nueva instancia de Gin sin middleware por defecto
	router := gin.New()
	//Agregacion de middleware de logging para que se muetre en consola cada requesr que llega al servidor
	router.Use(gin.Logger())
	//Registro de las demas rutas
	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)
	//Inicializacion del servidor
	//Ruta para prueba mientras se termina de hacer los controladores para la ruta
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Servidor funcionando correctamente",
		})
	})
	router.Run(":" + port)
}
