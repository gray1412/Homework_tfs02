package storage

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/ocg?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}

func CreateTable() {
	db := Connect()
	db.AutoMigrate(&Role{}, &Person{}, &Room{}, &Course{}, &Lession{}, &Person{}, &Class{}, &Registration{}, &Slot{})
}
