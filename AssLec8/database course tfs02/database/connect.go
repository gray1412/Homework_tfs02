package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectToDatabase() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:ngochd246@/tfs02?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database") // Kiểm tra kết nối tới databse
	}
	return
}
