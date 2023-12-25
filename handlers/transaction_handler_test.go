package handlers

import (
	"gocommerce/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func CalculateTotal(items []models.TransactionItem) float64 {
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func TestCalculateTotal(t *testing.T) {
	items := []models.TransactionItem{
		{Price: 10.0, Quantity: 2},
		{Price: 5.0, Quantity: 3},
	}

	total := CalculateTotal(items)

	// Menggunakan pustaka assert untuk memeriksa hasil
	assert.Equal(t, 35.0, total, "Total should be 35")
}

func TestGetTransactionWithItemsIntegration(t *testing.T) {

	tx := db.Begin()
	defer tx.Rollback()
	// Set up router
	router := gin.Default()
	router.GET("/transactions/:id", GetTransactionWithItems(db))

	// Membuat request HTTP simulasi
	req, _ := http.NewRequest("GET", "/transactions/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Memeriksa status kode
	assert.Equal(t, http.StatusOK, w.Code)

	// Memeriksa isi tanggapan
	assert.Contains(t, w.Body.String(), `"id":1`)

	tx.Rollback()
}
