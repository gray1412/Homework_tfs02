package connectDB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectToDatabase() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:anhvahieu2k@tcp(127.0.0.1:3306)/TestDB?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return
}
