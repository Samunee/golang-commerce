package handlers

import (
	"gocommerce/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		// Check if the username is already taken
		var existingUser models.User
		if err := db.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
			c.JSON(400, gin.H{"message": "Username already exists"})
			return
		}

		// Create the new user
		newUser := models.User{
			Username: input.Username,
			Password: input.Password, // You should hash the password before storing it
		}

		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}

		// Optionally, you can generate a token for the new user after registration.
		token, err := CreateToken(newUser.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}

		c.JSON(200, gin.H{"message": "User registered successfully", "token": token})
	}
}
