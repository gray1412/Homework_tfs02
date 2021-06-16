package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func saveFilm() {
	db := Connect()
	if db.HasTable(&flimDb{}) == false {
		db.CreateTable(&flimDb{})
	}
}
func saveArrivals() {
	db := Connect()
	if db.HasTable(&Product{}) == false {
		db.CreateTable(&Product{})
	}
}

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
