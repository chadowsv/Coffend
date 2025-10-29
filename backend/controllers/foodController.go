package controllers

import (
	"coffend/backend/database"
	"coffend/backend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var foods []models.Food

		if err := database.DB.WithContext(ctx).Preload("Menu").Find(&foods).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching food items"})
			return
		}

		c.JSON(http.StatusOK, foods)
	}
}

func GetFoodByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := c.Param("food_id")
		var food models.Food

		if err := database.DB.WithContext(ctx).First(&food, foodId).Error; err != nil {
			c.JSON(404, gin.H{"error": "Food Item not found."})
			return
		}
		c.JSON(200, food)
	}
}

func PostFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var newFood models.Food
		var menu models.Menu

		if err := c.BindJSON(&newFood); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(newFood); validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		if err := database.DB.WithContext(ctx).First(&menu, newFood.MenuID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu not found"})
			return
		}

		food := models.Food{
			Name:        newFood.Name,
			Description: newFood.Description,
			Price:       newFood.Price,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			MenuID:      newFood.MenuID,
		}

		if err := database.DB.WithContext(ctx).Create(&food).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create food"})
			return
		}

		c.JSON(http.StatusCreated, food)
	}
}

func PatchFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := c.Param("food_id")
		var inputFood models.Food
		var food models.Food

		if err := c.BindJSON(&inputFood); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(inputFood); validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		if err := database.DB.WithContext(ctx).First(&food, foodId).Error; err != nil {
			c.JSON(404, gin.H{"error": "Food not found."})
			return
		}

		updates := map[string]interface{}{
			"name":        inputFood.Name,
			"description": inputFood.Description,
			"price":       inputFood.Price,
			"menu_id":     inputFood.MenuID,
			"updated_at":  time.Now(),
		}

		if err := database.DB.WithContext(ctx).Model(&food).Updates(updates).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update food"})
			return
		}

		c.JSON(200, gin.H{"message": "Food updated successfully"})
	}
}

func DeleteFoodByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := c.Param("food_id")
		var food models.Food

		if err := database.DB.WithContext(ctx).Delete(&food, foodId).Error; err != nil {
			c.JSON(404, gin.H{"error": "Failed to delete food."})
			return
		}

		c.JSON(200, gin.H{"message": "Food eliminated successfully"})
	}
}
