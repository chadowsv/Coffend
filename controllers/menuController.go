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

type MenuCompleteViewFormat struct {
	MenuID       int    `json:"menu_id"`
	MenuName     string `json:"menuname"`
	MenuStatus   bool   `json:"menu_status" validate:"required"`
	FoodsDetails []models.FoodMenu
}

func GetMenuCom() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Obtener todos los menús
		queryMenu := "SELECT menu_id, name as menuname, menu_status FROM Menus"
		rows, err := database.DB.QueryContext(ctx, queryMenu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching menus"})
			return
		}
		defer rows.Close()

		var menus []MenuCompleteViewFormat

		for rows.Next() {
			var menu MenuCompleteViewFormat
			err := rows.Scan(&menu.MenuID, &menu.MenuName, &menu.MenuStatus)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning menu"})
				return
			}

			// Obtener los alimentos de este menú
			foodsItems, err := FoodsByMenu(menu.MenuID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching foods"})
				return
			}
			menu.FoodsDetails = foodsItems

			menus = append(menus, menu)
		}

		c.JSON(http.StatusOK, menus)
	}
}

// Ahora filtramos los alimentos por menú
func FoodsByMenu(menuID int) ([]models.FoodMenu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	query := `SELECT menu_id, name, description, price FROM Foods WHERE menu_id = @p1`
	rows, err := database.DB.QueryContext(ctx, query, menuID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foods []models.FoodMenu
	for rows.Next() {
		var food models.FoodMenu
		err := rows.Scan(&food.MenuID, &food.Name, &food.Description, &food.Price)
		if err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}
	return foods, nil
}

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Contexto creado para que no dure tanto
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		//Creacion de slice para manejar la lista de menus que tenemos
		var menus []models.Menu
		//Query para consultar los menus
		queryMenus := "SELECT * FROM Menus"
		//Se establece como resultado las filas al igual que el error que puede obtenerse de la ejecucion de la consulta
		rows, err := database.DB.QueryContext(ctx, queryMenus)
		//Si existe algun tipo de error, se envia error en la base de datos
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
		defer rows.Close()
		//Se recorren las filas de los resultados de las consultas
		for rows.Next() {
			//por cada fila recorrida se crea un nuevo menu de tipo models.menu
			var menu models.Menu
			// err va escaneando y guardando en los atributos del struct los datos recopilados
			err := rows.Scan(
				&menu.MenuID,
				&menu.Name,
				&menu.MenuStatus,
				&menu.Created_at,
				&menu.Updated_at,
			)
			//si se da algun error se evita la obtencion de los menus y se sale de la funcion
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse food"})
				return
			}
			//si todo sale bien se agrega al slice menus, el menu recien obtenido
			menus = append(menus, menu)
		}
		// nos devuelve con exito el HTTP rquest y regresa el slice de menus
		c.JSON(200, menus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Creacopm de una conexion a la base de datos que no dura siempre
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		//Cancelacion del contexto al terminar la funcion
		defer cancel()
		//Se obtiene el parametro de food_item_id que viene de la URL
		menuId := c.Param("menu_id")
		//se crea una variable de tipo FoodItem, aqui es donde se guardara el resultado de SELECT
		var menu models.Menu
		//Consulta de SQL Server para el resultado de un plato, usando ctx para que se cancele si demora mucho
		err := database.DB.QueryRowContext(ctx, "SELECT menu_id,name,menu_status,created_at,updated_at FROM Menus WHERE menu_id  = @p1", menuId).Scan(
			&menu.MenuID,
			&menu.Name,
			&menu.MenuStatus,
			&menu.Created_at,
			&menu.Updated_at,
		)
		//Scan hace que se copien los valores que vienen de la BDD dentro de los campos de foodItem
		//Si se encuentra con un error se detiene la ejecucion
		if err != nil {
			//Envío de una respuesta en formato JSON que significa que no se encontró en la BDD
			c.JSON(404, gin.H{"error": "Menu not created."})
			return
		}
		//Envio de una respuesta con formato JSON con código HTTP de éxito y el objeto foodItem
		c.JSON(200, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var newMenu models.Menu
		//Lectura del JSON de la solicitud, intenta hacer un objeto de tipo Menu
		//Revision de JSON con buen formato y compatible con estructura
		if err := c.BindJSON(&newMenu); err != nil {
			//Asignacion corta con verificacion de error para devolver un error si los formatos no coinciden
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		//Validaciones puestas en los structs
		var validate = validator.New()
		validationErr := validate.Struct(newMenu)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}
		//Se realiza la Query y se guarda el resultado de la consulta SQL en menu.MenuID, para guardar se necesita la direccion de la variable
		newMenu.Created_at = time.Now()
		newMenu.Updated_at = time.Now()
		//Con el scan se guarda el resultado de la columna en la variable menu.MenuID
		insertQuery := "INSERT INTO menus (name,menu_status,created_at,updated_at) VALUES (@p1,@p2,@p3,@p4);"
		_, insertErr := database.DB.ExecContext(ctx, insertQuery,
			newMenu.Name,
			newMenu.MenuStatus,
			newMenu.Created_at,
			newMenu.Updated_at,
		)
		if insertErr != nil {
			c.JSON(500, gin.H{"error": "Failed to create menu"})
			return
		}
		c.JSON(201, newMenu)
	}

}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu models.Menu
		//Revision del formato
		//En el espacio de memoria de menu se obtiene la fuente principal de informacion y no una copia
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		//Revision de las validaciones
		var validate = validator.New()
		validationErr := validate.Struct(menu)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}
		menu_id := c.Param("menu_id")
		menu.Updated_at = time.Now()
		updateQuery := "UPDATE Menus SET name=@p1, menu_status=@p2,updated_at=@p3 WHERE menu_id =@p4"
		_, updateErr := database.DB.ExecContext(ctx, updateQuery,
			menu.Name,
			menu.MenuStatus,
			menu.Updated_at,
			menu_id,
		)
		if updateErr != nil {
			c.JSON(500, gin.H{"error": "Failed to update de menu"})
			return
		}
		c.JSON(200, gin.H{"message": "Menu updated successfully"})
	}
}
