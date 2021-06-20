package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	Title string `json:"title"`
	Price int    `json:"price"`
}

type Products struct {
	Products []Product `json:"products"`
}

type Movie struct {
	Name    string `json:"name" db:"name"`
	Ranking string `json:"ranking" db:"ranking"`
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Polarbear1011@/mysql_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.CreateTable(&Movie{})
	db.CreateTable(&Product{})
	return db
}
