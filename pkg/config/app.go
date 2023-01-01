package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "root:12345678@tcp(127.0.0.1:3399)/bookstore_db?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = database
}
func GetDb() *gorm.DB {
	return db
}
