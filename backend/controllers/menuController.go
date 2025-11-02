package controllers

import (
	"coffend/backend/database"
	"coffend/backend/models"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetAllMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menus []models.Menu

		if err := database.DB.WithContext(ctx).
			Preload("Foods").
			Find(&menus).Error; err != nil {
			c.JSON(500, gin.H{"error": "Error fetching menus"})
			return
		}

		c.JSON(200, menus)
	}
}

func GetMenuByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		menuId := c.Param("menu_id")
		var menu models.Menu

		if err := database.DB.WithContext(ctx).First(&menu, menuId).Error; err != nil {
			c.JSON(404, gin.H{"error": "Menu not found."})
			return
		}

		c.JSON(200, menu)
	}
}

func PostMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var newMenu models.Menu

		if err := c.BindJSON(&newMenu); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(newMenu); validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		menu := models.Menu{
			Name:       newMenu.Name,
			MenuStatus: newMenu.MenuStatus,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		if err := database.DB.WithContext(ctx).Create(&menu).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create menu"})
			return
		}

		c.JSON(201, menu)
	}

}

func UpdateMenuByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		menuId := c.Param("menu_id")
		var inputMenu models.Menu
		var menu models.Menu

		if err := c.BindJSON(&inputMenu); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if validationErr := validate.Struct(inputMenu); validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		if err := database.DB.WithContext(ctx).First(&menu, menuId).Error; err != nil {
			c.JSON(404, gin.H{"error": "Menu not found."})
			return
		}

		updates := map[string]interface{}{
			"name":        inputMenu.Name,
			"menu_status": inputMenu.MenuStatus,
			"updated_at":  time.Now(),
		}

		if err := database.DB.WithContext(ctx).Model(&menu).Updates(updates).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update menu"})
			return
		}

		c.JSON(200, gin.H{"message": "Menu updated successfully"})
	}
}

func DeleteMenuByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		menuId := c.Param("menu_id")
		var menu models.Menu

		if err := database.DB.WithContext(ctx).First(&menu, menuId).Error; err != nil {
			c.JSON(404, gin.H{"error": "Menu not found."})
			return
		}

		if err := database.DB.WithContext(ctx).Delete(&menu, menuId).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete menu."})
			return
		}
		c.JSON(200, gin.H{"message": "Menu deleted successfully"})
	}
}
