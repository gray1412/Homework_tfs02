package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqlCredentials = fmt.Sprintf(
	"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	"root",
	"keobng956",
	"tcp",
	"localhost",
	"3306",
	"tfs2",
)
var db, _ = gorm.Open("mysql", mysqlCredentials)


func ConnectDb() (db *gorm.DB) {
	db, err := gorm.Open("mysql", mysqlCredentials)
	if err != nil {
		panic("Couldn't connect to DB")
	}
	fmt.Println("Connected to DB...")
	return
}