package handlers

import (
	"gocommerce/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetTransactionWithItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var transaction models.Transaction
		if err := db.Preload("Items").First(&transaction, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Transaction not found"})
			return
		}

		c.JSON(200, transaction)
	}
}

func CreateTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Transaction
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		// Jika Anda ingin mengaitkan item transaksi, pastikan item-item tersebut ada dalam database
		for _, item := range input.Items {
			var product models.Product
			if err := db.First(&product, item.ProductID).Error; err != nil {
				c.JSON(400, gin.H{"message": "Invalid product ID"})
				return
			}
		}

		db.Create(&input)
		c.JSON(201, input)
	}
}
