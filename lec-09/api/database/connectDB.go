package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectDB() *gorm.DB {
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"123456",
		"tcp",
		"localhost",
		"3306",
		"project-tfs02",
	)
	db, err := gorm.Open("mysql", mysqlCredentials)
	if err != nil {
		fmt.Println(err)
		panic("fall to connect DB")
	}
	// defer db.Close()
	return db
}
