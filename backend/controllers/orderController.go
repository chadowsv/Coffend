package controllers

import (
	"coffend/backend/database"
	"coffend/backend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var orders []models.Order

		if err := database.DB.WithContext(ctx).Find(&orders).Error; err != nil {
			c.JSON(500, gin.H{"error": "Error fetching orders"})
			return
		}

		c.JSON(200, orders)
	}
}

func GetOrderByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderId := c.Param("order_id")
		var order models.Order

		if err := database.DB.WithContext(ctx).First(&order, orderId).Error; err != nil {
			c.JSON(404, gin.H{"error": "Order not found."})
			return
		}

		c.JSON(200, order)
	}
}

func PostOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var newOrder models.Order

		if err := c.BindJSON(&newOrder); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(newOrder); validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		if tableExists(ctx, *newOrder.TableID) {
			c.JSON(400, gin.H{"error": "Table not found"})
			return
		}

		order := models.Order{
			OrderDate: time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			TableID:   newOrder.TableID,
			Total:     newOrder.Total,
		}

		if err := database.DB.WithContext(ctx).Create(&order).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create order"})
			return
		}

		c.JSON(http.StatusCreated, order)
	}
}

func PatchOrderByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderID := c.Param("order_id")
		var inputOrder models.Order
		var order models.Order

		if err := c.BindJSON(&inputOrder); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(inputOrder)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		if !tableExists(ctx, *inputOrder.TableID) {
			c.JSON(400, gin.H{"error": "Table not found"})
			return
		}

		if err := database.DB.WithContext(ctx).First(&order, orderID).Error; err != nil {
			c.JSON(404, gin.H{"error": "Order not found."})
			return
		}

		updates := map[string]interface{}{
			"order_date": inputOrder.OrderDate,
			"table_id":   inputOrder.TableID,
			"total":      inputOrder.Total,
			"updated_at": time.Now(),
		}

		if err := database.DB.WithContext(ctx).Model(&order).Updates(updates).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update the order"})
			return
		}

		c.JSON(200, gin.H{"message": "Order updated successfully"})
	}
}

func DeleteOrderByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderID := c.Param("order_id")

		if err := database.DB.WithContext(ctx).Delete(&models.Order{}, orderID).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete the order"})
			return
		}

		c.JSON(200, gin.H{"message": "Order deleted successfully"})
	}
}

func tableExists(ctx context.Context, tableID int) bool {
	var table models.Table
	if err := database.DB.WithContext(ctx).First(&table, tableID).Error; err != nil {
		return false
	}
	return true
}
