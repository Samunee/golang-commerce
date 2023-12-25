package seeders

import (
	"gocommerce/models"

	"github.com/jinzhu/gorm"
)

func Seed(db *gorm.DB) {
	// Seeder untuk Product Categories
	category1 := models.ProductCategory{Name: "Category 1"}
	category2 := models.ProductCategory{Name: "Category 2"}

	db.Create(&category1)
	db.Create(&category2)

	// Seeder untuk Users
	user1 := models.User{Username: "user1", Email: "user1@example.com", Password: "password1"}
	user2 := models.User{Username: "user2", Email: "user2@example.com", Password: "password2"}

	db.Create(&user1)
	db.Create(&user2)

	// Seeder untuk Products
	product1 := models.Product{Name: "Product 1", CategoryID: 1}
	product2 := models.Product{Name: "Product 2", CategoryID: 2}

	db.Create(&product1)
	db.Create(&product2)

	// Seeder untuk Transactions
	transaction1 := models.Transaction{UserID: user1.ID, Amount: 100.0}
	transaction2 := models.Transaction{UserID: user2.ID, Amount: 200.0}

	db.Create(&transaction1)
	db.Create(&transaction2)

	// Seeder untuk Transaction Items
	item1 := models.TransactionItem{TransactionID: transaction1.ID, ProductID: product1.ID, Quantity: 2}
	item2 := models.TransactionItem{TransactionID: transaction2.ID, ProductID: product2.ID, Quantity: 3}

	db.Create(&item1)
	db.Create(&item2)
}
