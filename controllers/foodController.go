package controllers

import (
	"coffend/database"
	"coffend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Retorna un manejador HTTP
// Se encapsula en gin.Handler para integrarse con el enrutador
func GetFoods() gin.HandlerFunc {
	//Contexto que da acceso a la solicitud y respuesta, objeto de gin que contiene toda la informacion de la solicitud y respuesta
	return func(c *gin.Context) {
		//Evita que la consulta se quede colgada mucho tiempo
		//Manejo de deadlines, cancelaciones y pasar valores en operaciones concurrentes
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var foods []models.Food
		queryFoods := "SELECT * FROM Foods"
		rows, err := database.DB.QueryContext(ctx, queryFoods)
		if err != nil {
			//serializa datos a JSON y establece el header
			c.JSON(500, gin.H{"error": "Database error"})
			return
		}
		//Cierre de filas una vez procesadas, previene resource leaks
		defer rows.Close()
		//se recorre cada fila con el next
		for rows.Next() {
			var food models.Food
			err := rows.Scan(
				&food.FoodID,
				&food.Name,
				&food.Description,
				&food.Price,
				&food.Created_at,
				&food.Updated_at,
				&food.MenuID,
			)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to parse food"})
				return
			}
			foods = append(foods, food)
		}
		c.JSON(200, foods)
	}
}

// Declaracion de la funcion GetFoodItem que devuelve un handler o una funcion Gin para responder la solitcitud HTTP
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Creacopm de una conexion a la base de datos que no dura siempre
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		//Cancelacion del contexto al terminar la funcion
		defer cancel()
		//Se obtiene el parametro de food_item_id que viene de la URL
		foodId := c.Param("food_id")
		//se crea una variable de tipo FoodItem, aqui es donde se guardara el resultado de SELECT
		var food models.Food
		//Consulta de SQL Server para el resultado de un plato, usando ctx para que se cancele si demora mucho
		err := database.DB.QueryRowContext(ctx, "SELECT food_id, name, description, price, created_at, updated_at, menu_id FROM Foods WHERE food_id = @p1", foodId).Scan(
			&food.FoodID,
			&food.Name,
			&food.Description,
			&food.Price,
			&food.Created_at,
			&food.Updated_at,
			&food.MenuID,
		)
		//Scan hace que se copien los valores que vienen de la BDD dentro de los campos de foodItem
		//Si se encuentra con un error se detiene la ejecucion
		if err != nil {
			//Envío de una respuesta en formato JSON que significa que no se encontró en la BDD
			c.JSON(404, gin.H{"error": "Food Item not created."})
			return
		}
		//Envio de una respuesta con formato JSON con código HTTP de éxito y el objeto foodItem
		c.JSON(200, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		//Se usa el menu id para crear un plato
		var menu models.Menu
		var newFood models.Food
		//Lectura del JSON de la solicitud, intenta hacer un objeto de tipo Fodd
		//Revision de JSON con buen formato y compatible con estructura
		if err := c.BindJSON(&newFood); err != nil {
			//Asignacion corta con verificacion de error para devolver un error si los formatos no coinciden
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		//Validaciones puestas en los structs
		validationErr := validate.Struct(newFood)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}
		//Se realiza la Query y se guarda el resultado de la consulta SQL en menu.MenuID, para guardar se necesita la direccion de la variable
		err := database.DB.QueryRowContext(ctx, "SELECT menu_id FROM Menus WHERE menu_id = @p1", newFood.MenuID).Scan(&menu.MenuID)
		//Scan toma los valores de la fila de resultado y los asigna a las variables
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu not found"})
			return
		}
		newFood.Created_at = time.Now()
		newFood.Updated_at = time.Now()
		//Se accede al valor que esta en esa direccion de memoria
		//Con el scan se guarda el resultado de la columna en la variable menu.MenuID
		insertQuery := "INSERT INTO	Foods (name,description,price, created_at, updated_at,menu_id) VALUES (@p1,@p2,@p3,@p4,@p5,@p6);"
		//Se ignora el resultado que diria la cantidad de filas afectadas y nos interesa el error
		_, insertErr := database.DB.ExecContext(ctx, insertQuery,
			newFood.Name,
			newFood.Description,
			newFood.Price,
			newFood.Created_at,
			newFood.Updated_at,
			newFood.MenuID,
		)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect create food"})
			return
		}
		c.JSON(http.StatusCreated, newFood)
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var food models.Food
		var menu models.Menu
		//Revision del formato
		if err := c.BindJSON(&food); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		//Revision de las validaciones
		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}
		food_id := c.Param("food_id")
		err := database.DB.QueryRowContext(ctx, "SELECT menu_id FROM Menus WHERE menu_id = @p1", food.MenuID).Scan(&menu.MenuID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Food not found"})
			return
		}
		food.Updated_at = time.Now()
		updateQuery := "UPDATE Foods SET name=@p1,description=@p2,price=@p3, updated_at=@p4,menu_id=@p5 WHERE food_id=@p6"
		_, updateErr := database.DB.ExecContext(ctx, updateQuery,
			food.Name,
			food.Description,
			food.Price,
			food.Updated_at,
			food.MenuID,
			food_id,
		)
		if updateErr != nil {
			c.JSON(500, gin.H{"error": "Failed to update de food"})
			return
		}
		c.JSON(200, gin.H{"message": "Food updated successfully"})
	}
}

// Delete Food based on food_id
func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		food_id := c.Param("food_id")
		deleteQuery := "DELETE FROM Foods WHERE food_id=@p1"
		_, deleteErr := database.DB.ExecContext(ctx, deleteQuery, food_id)
		if deleteErr != nil {
			c.JSON(500, gin.H{"error": "Failed to delete food"})
			return
		}
		c.JSON(200, gin.H{"message": "Food eliminated successfully"})
	}
}
