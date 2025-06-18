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
	//middleware se refiere a que las funciones que se ejecutan en la cadena de procesamiento de una solicitud HTTP
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
	//Inicia el servidor en el puerto
	router.Run(":" + port)
	//404 Not Found: Recurso no existe (ej: GET con ID inválido).
	//500 Internal Server Error: Error inesperado (ej: conexión DB fallida).
	//400 Bad Request: Datos malformados (validación fallida).
	//BindJSON: Verifica sintaxis JSON básica + tipos de datos.
	//validate.Struct: Valida reglas de negocio (campos requeridos, formatos, etc.).
}
