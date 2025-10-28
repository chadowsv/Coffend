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

		if err := database.DB.WithContext(ctx).Find(&orderItems).Error; err != nil {
			c.JSON(500, gin.H{"error": "Error fetching order items"})
			return
		}

		c.JSON(200, orderItems)
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderIdStr := c.Param("order_id")
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid order ID"})
			return
		}

		var orderItems []models.OrderItem

		if err := database.DB.WithContext(ctx).Preload("Food").Where("order_id = ?", orderId).Find(&orderItems).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch order items"})
			return
		}

		if len(orderItems) == 0 {
			c.JSON(404, gin.H{"error": "No order items found for this order"})
			return
		}

		c.JSON(200, orderItems)
	}
}

func GetOrderItemByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderItemId := c.Param("order_item_id")
		var orderItem models.OrderItem

		if err := database.DB.WithContext(ctx).First(&orderItem, orderItemId).Error; err != nil {
			c.JSON(404, gin.H{"error": "Order not found."})
			return
		}

		c.JSON(200, orderItem)
	}
}

func PostOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var newOrderItems []models.OrderItem

		if err := c.BindJSON(&newOrderItems); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.WithContext(ctx).Create(&newOrderItems).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order item"})
			return
		}

		c.JSON(200, gin.H{"message": "Order Item successfully created"})
	}
}

func PatchOrderItemByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var inputOrderItem models.OrderItem
		var orderItem models.OrderItem

		orderItemId := c.Param("order_item_id")
		if err := c.BindJSON(&inputOrderItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(inputOrderItem); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		if err := database.DB.WithContext(ctx).First(&orderItem, orderItemId).Error; err != nil {
			c.JSON(404, gin.H{"error": "Order item not found"})
			return
		}

		updates := map[string]interface{}{
			"quantity":   inputOrderItem.Quantity,
			"unit_price": inputOrderItem.UnitPrice,
			"updated_at": time.Now(),
			"order_id":   inputOrderItem.OrderID,
			"food_id":    inputOrderItem.FoodID,
		}

		if err := database.DB.WithContext(ctx).Model(&orderItem).Updates(updates).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update order item"})
			return
		}

		c.JSON(200, gin.H{"message": "Order item updated Successfully"})
	}
}

func DeleteOrderItemByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderItemId := c.Param("order_item_id")

		if err := database.DB.WithContext(ctx).Delete(&models.OrderItem{}, orderItemId).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete order item"})
			return
		}

		c.JSON(200, gin.H{"message": "Order item deleted successfully"})
	}
}
