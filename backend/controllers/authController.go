package controllers

import (
	"coffend/backend/database"
	"coffend/backend/models"
	"coffend/backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to hash password"})
			return
		}

		user.Password = hashedPassword
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		if err := database.DB.Create(&user).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(201, gin.H{"message": "User registered successfully"})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var user models.User

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		if !utils.CheckPassword(user.Password, input.Password) {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		accessToken, refreshToken, err := utils.GenerateToken(user.UserID, user.Role)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate tokens"})
			return
		}

		updates := map[string]interface{}{
			"token":         accessToken,
			"refresh_token": refreshToken,
			"updated_at":    time.Now(),
		}

		if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update user tokens"})
			return
		}

		c.JSON(200, gin.H{"message": "Login successful"})
	}
}
