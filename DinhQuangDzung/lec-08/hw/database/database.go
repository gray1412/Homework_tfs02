package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CsvLine struct {
	Type  string
	Title string
	Body  string
}

func ConnectSQL() (db *gorm.DB) {
	dsn := "root:Polarbear1011@/mysql_db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error when connect to db ", err)
		return
	}
	if err != nil {
		log.Fatal("error when auto migrate table ", err)
	}
	return db
}
