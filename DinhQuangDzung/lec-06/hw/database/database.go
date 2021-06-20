package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	Id   int    `json:"id"  db:"id"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age"  db:"age"`
}

func ConnectDb() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:Polarbear1011@/mysql_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Couldn't connect to DB")
	}
	fmt.Println("Connected to DB...")
	return
}
