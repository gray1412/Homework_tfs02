package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456@/test_crawl?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func CreateReview() {
	db := Connect()
	if !db.HasTable(&Review{}) {
		db.CreateTable(&Review{})
	}
}
