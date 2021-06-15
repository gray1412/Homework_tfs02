package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func ConnectDatabase() (db *gorm.DB) {

	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"123456",
		"tcp",
		"localhost",
		"3306",
		"crawl",
	)
	db, err := gorm.Open("mysql", mysqlCredentials)
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to the database!")
	}
	defer db.Close()
	return
}
