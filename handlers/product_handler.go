package handlers

import (
	"gocommerce/models"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// contoh lama
// func ListProducts(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var products []models.Product
// 		db.Find(&products)
// 		c.JSON(200, products)
// 	}
// }

// contoh baru dengan goroutine

// @Summary Get a list of products
// @Description Retrieve a list of products from the database.
// @Produce json
// @Success 200 {array} Product
// @Router /products [get]
func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		var wg sync.WaitGroup

		// Menambahkan satu goroutine ke WaitGroup
		wg.Add(1)

		// Memulai goroutine untuk melakukan operasi yang membutuhkan waktu lama
		go func() {
			defer wg.Done() // Menandai bahwa goroutine telah selesai
			db.Find(&products)
		}()

		// Menunggu goroutine selesai
		wg.Wait()

		c.JSON(200, products)
	}
}

func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}
		c.JSON(200, product)
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Create(&input)
		c.JSON(201, input)
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}

		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Model(&product).Updates(input)
		c.JSON(200, product)
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}

		db.Delete(&product)
		c.JSON(200, gin.H{"message": "Product deleted"})
	}
}
