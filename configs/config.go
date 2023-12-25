package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:Password@tcp(127.0.0.1:3306)/ecomm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	DB = db
	return DB, nil
}
