package services

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const connectionString string = "admin:admin@tcp(127.0.0.1:3306)/CollegeSystems?charset=utf8mb4&parseTime=True&loc=Local"

func GetDB() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}
