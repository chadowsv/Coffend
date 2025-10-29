package controllers

import (
	"coffend/backend/database"
	"coffend/backend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var tables []models.Table

		if err := database.DB.WithContext(ctx).Find(&tables).Error; err != nil {
			c.JSON(500, gin.H{"error": "Error fetching tables"})
			return
		}

		c.JSON(200, tables)
	}
}

func GetTableByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		tableId := c.Param("table_id")
		var table models.Table

		if err := database.DB.WithContext(ctx).First(&table, tableId).Error; err != nil {
			c.JSON(404, gin.H{"error": "Table not found."})
			return
		}

		c.JSON(200, table)
	}
}

func PostTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var newTable models.Table

		if err := c.BindJSON(&newTable); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(newTable); validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		table := models.Table{
			NumberGuests: newTable.NumberGuests,
			Status:       newTable.Status,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		if err := database.DB.WithContext(ctx).Create(&table).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create table"})
			return
		}

		c.JSON(http.StatusCreated, table)
	}
}

func PatchTableByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		tableID := c.Param("table_id")
		var inputTable models.Table
		var table models.Table

		if err := c.BindJSON(&inputTable); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(inputTable); validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		if err := database.DB.WithContext(ctx).First(&table, tableID).Error; err != nil {
			c.JSON(404, gin.H{"error": "Table not found."})
			return
		}

		updates := map[string]interface{}{
			"number_guests": table.NumberGuests,
			"status":        table.Status,
			"updated_at":    time.Now(),
		}

		if err := database.DB.WithContext(ctx).Model(&table).Updates(updates).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update table"})
			return
		}

		c.JSON(200, gin.H{
			"message":  "Table updated successfully",
			"table_id": tableID,
			"changes": gin.H{
				"number_guests": table.NumberGuests,
				"status":        table.Status,
			},
		})
	}
}

func DeleteTableByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		tableID := c.Param("table_id")

		if err := database.DB.WithContext(ctx).Delete(&models.Table{}, tableID).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete table"})
			return
		}

		c.JSON(200, gin.H{"message": "Table deleted successfully"})
	}
}
