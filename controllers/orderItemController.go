package controllers

import (
	"coffend/database"
	"coffend/models"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var orderItems []models.OrderItem
		queryOrderItems := "SELECT * FROM OrderItems"
		rows, err := database.DB.QueryContext(ctx, queryOrderItems)
		if err != nil {
			c.JSON(500, gin.H{"error": "Database error"})
			return
		}
		defer rows.Close()
		for rows.Next() {
			var orderItem models.OrderItem
			err := rows.Scan(
				&orderItem.OrderItemID,
				&orderItem.Quantity,
				&orderItem.UnitPrice,
				&orderItem.Created_at,
				&orderItem.Updated_at,
				&orderItem.OrderID,
				&orderItem.FoodID,
			)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to parse order item"})
				return
			}
			orderItems = append(orderItems, orderItem)
		}
		c.JSON(200, orderItems)
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderIdStr := c.Param("order_id")
		//Conversion de lo que da el JSON, a manera de string para la funcion convertirlo en int como pide la funcion
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid order ID"})
			return
		}
		orderItems, err := ItemsByOrder(orderId)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch order items"})
			return
		}
		c.JSON(200, orderItems)
	}
}
func ItemsByOrder(id int) (OrderItems []models.OrderItemExtended, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	query := `SELECT 
			oi.order_item_id,
			oi.order_id,
			oi.food_id,
			f.name AS food_name,
			oi.quantity,
			oi.unit_price,
			(oi.quantity * oi.unit_price) AS subtotal,
			oi.created_at,
			oi.updated_at
		FROM OrderItems oi
		INNER JOIN Foods f ON oi.food_id = f.food_id
		WHERE oi.order_id = @p1`
	rows, err := database.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.OrderItemExtended
	for rows.Next() {
		var item models.OrderItemExtended
		err := rows.Scan(
			&item.OrderItemID,
			&item.OrderID,
			&item.FoodID,
			&item.FoodName,
			&item.Quantity,
			&item.UnitPrice,
			&item.Subtotal,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		orderItemId := c.Param("order_item_id")
		var orderItem models.OrderItem
		queryOrderItem := "	SELECT * FROM OrderItems WHERE order_item_id= @p1"
		err := database.DB.QueryRowContext(ctx, queryOrderItem, orderItemId).Scan(
			&orderItem.OrderItemID,
			&orderItem.Quantity,
			&orderItem.UnitPrice,
			&orderItem.Created_at,
			&orderItem.Updated_at,
			&orderItem.OrderID,
			&orderItem.FoodID,
		)
		if err != nil {
			c.JSON(404, gin.H{"error": "Order Item not found."})
			return
		}
		c.JSON(200, orderItem)
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var orderItem models.OrderItem
		if err := c.BindJSON(&orderItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(orderItem)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		orderItem.Created_at = time.Now()
		orderItem.Updated_at = time.Now()
		query := "INSERT INTO OrderItems (quantity,unit_price,created_at,updated_at,order_id,food_id) VALUES (@p1,@p2,@p3,@p4,@p5,@p6)"
		_, err := database.DB.ExecContext(ctx, query,
			orderItem.Quantity,
			orderItem.UnitPrice,
			orderItem.Created_at,
			orderItem.Updated_at,
			orderItem.OrderID,
			orderItem.FoodID,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create Order Item"})
			return
		}

		c.JSON(200, gin.H{"message": "Order Item successfully"})
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var orderItem models.OrderItem
		orderItemId := c.Param("order_item_id")
		if err := c.BindJSON(&orderItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(orderItem)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		orderItem.Updated_at = time.Now()
		updateQuery := "UPDATE OrderItems SET quantity=@p1, unit_price=@p2,updated_at=@p3,order_id=@p4,food_id=@p5 WHERE order_item_id=@p6"
		_, erro := database.DB.ExecContext(ctx, updateQuery,
			orderItem.Quantity,
			orderItem.UnitPrice,
			orderItem.Updated_at,
			orderItem.OrderID,
			orderItem.FoodID,
			orderItemId,
		)
		if erro != nil {
			c.JSON(500, gin.H{"error": "Failes to update the order item"})
			return
		}
		c.JSON(200, gin.H{"message": "Order item updated Successfully"})
	}
}
