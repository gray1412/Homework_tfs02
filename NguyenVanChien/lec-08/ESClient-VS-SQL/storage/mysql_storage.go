package storage

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

func ConnectDatabase() (db *gorm.DB) {
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"",
		"tcp",
		"localhost",
		"3306",
		"lab08",
	)
	db, err := gorm.Open("mysql", mysqlCredentials) // Open database

	if err != nil {
		panic("Failed to connect database") // Check connect to database
	}
	return
}

// Create table book
func CreateTableBook() {
	db := *ConnectDatabase()
	// defer db.Close()
	if (!db.HasTable(&Book{})) {
		db.CreateTable(&Book{})
	}
}
