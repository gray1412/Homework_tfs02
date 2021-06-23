package storage

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json: "name"`
	Price       string `json: "price"`
	Img         string `json:"img"`
	Description string `json:"description"`
}

var Products = map[int]Product{}

func ConnectDatabase() (db *gorm.DB) {
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"",
		"tcp",
		"localhost",
		"3306",
		"testproducts",
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
	if (!db.HasTable(&Product{})) {
		db.CreateTable(&Product{})
	} else {
		db.DropTable(&Product{})
		db.CreateTable(&Product{})
	}
}
