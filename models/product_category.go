package models

type ProductCategory struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	// Tambahkan kolom-kolom lain sesuai kebutuhan
}
