package storage

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() (db *gorm.DB) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_crawl?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect Failed!")
		panic(err)
	} else {
		fmt.Println("Connect Successful!")
	}
	return
}
