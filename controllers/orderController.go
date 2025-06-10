package controllers

import (
	"coffend/database"
	"coffend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Contexto para que no dure tanto
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var orders []models.Order
		queryOrders := "SELECT * FROM Orders"
		rows, err := database.DB.QueryContext(ctx, queryOrders)
		if err != nil {
			c.JSON(500, gin.H{"error": "Database error"})
			return
		}
		defer rows.Close()
		for rows.Next() {
			var order models.Order
			err := rows.Scan(
				&order.OrderID,
				&order.OrderDate,
				&order.Created_at,
				&order.Updated_at,
				&order.TableID,
				&order.Total,
			)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to parse order"})
				return
			}
			orders = append(orders, order)
		}
		c.JSON(200, orders)
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		orderId := c.Param("order_id")
		var order models.Order
		queryOrder := "	SELECT * FROM Orders WHERE order_id= @p1"
		err := database.DB.QueryRowContext(ctx, queryOrder, orderId).Scan(
			&order.OrderID,
			&order.OrderDate,
			&order.Created_at,
			&order.Updated_at,
			&order.TableID,
			&order.Total,
		)
		if err != nil {
			c.JSON(404, gin.H{"error": "Order not created."})
			return
		}
		c.JSON(200, order)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var order models.Order
		var table models.Table
		if err := c.BindJSON(&order); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(order)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}
		err := database.DB.QueryRowContext(ctx, "SELECT table_id FROM Tables WHERE table_id = @p1", order.TableID).Scan(&table.TableID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Table not found"})
			return
		}
		order.OrderDate = time.Now()
		order.Created_at = time.Now()
		order.Updated_at = time.Now()
		insertQuery := "INSERT INTO Orders (order_date,created_at,updated_at,table_id,total) VALUES (@p1,@p2,@p3,@p4,@p5)"
		_, insertErr := database.DB.ExecContext(ctx, insertQuery,
			order.OrderDate,
			order.Created_at,
			order.Updated_at,
			order.TableID,
			order.Total,
		)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect create order"})
			return
		}
		c.JSON(http.StatusCreated, order)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var order models.Order
		var table models.Table

		// Validación del formato JSON
		if err := c.BindJSON(&order); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Validación con validator
		validationErr := validate.Struct(order)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		orderID := c.Param("order_id")

		// Verificar que la mesa existe
		err := database.DB.QueryRowContext(ctx, "SELECT table_id FROM Tables WHERE table_id = @p1", order.TableID).Scan(&table.TableID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Table not found"})
			return
		}

		order.Updated_at = time.Now()

		// Query para actualizar la orden
		updateQuery := `
			UPDATE Orders 
			SET order_date=@p1, created_at=@p2, updated_at=@p3, table_id=@p4, total=@p5 
			WHERE order_id=@p6
		`

		_, updateErr := database.DB.ExecContext(ctx, updateQuery,
			order.OrderDate,
			order.Created_at,
			order.Updated_at,
			order.TableID,
			order.Total,
			orderID,
		)

		if updateErr != nil {
			c.JSON(500, gin.H{"error": "Failed to update the order"})
			return
		}

		c.JSON(200, gin.H{"message": "Order updated successfully"})
	}
}
