package storage

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json: "name"`
	Age  int    `json: "age"`
}

var Students = map[int]Student{}

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
	db, err := gorm.Open("mysql", mysqlCredentials) // Mở database

	if err != nil {
		panic("Failed to connect database") // Kiểm tra kết nối tới database
	}
	return
}

func CreateTableProduct() {
	db := *ConnectDatabase()
	// defer db.Close()
	if (!db.HasTable(&Student{})) {
		db.CreateTable(&Student{})
	}
}
