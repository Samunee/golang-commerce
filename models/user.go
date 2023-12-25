package models

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Tambahkan kolom-kolom lain sesuai kebutuhan
}
