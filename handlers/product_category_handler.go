package handlers

import (
	"gocommerce/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListProductCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []models.ProductCategory
		db.Find(&categories)
		c.JSON(200, categories)
	}
}

func GetProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var category models.ProductCategory
		if err := db.First(&category, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Category not found"})
			return
		}
		c.JSON(200, category)
	}
}

func CreateProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.ProductCategory
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		if input.Name == "" {
			c.JSON(400, gin.H{"message": "Name is required"})
			return
		}

		db.Create(&input)
		c.JSON(201, input)
	}
}

func UpdateProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var category models.ProductCategory
		if err := db.First(&category, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Category not found"})
			return
		}

		var input models.ProductCategory
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Model(&category).Updates(input)
		c.JSON(200, category)
	}
}

func DeleteProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var category models.ProductCategory
		if err := db.First(&category, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Category not found"})
			return
		}

		db.Delete(&category)
		c.JSON(200, gin.H{"message": "Category deleted"})
	}
}
