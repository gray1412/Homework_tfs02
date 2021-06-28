package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func ConnectDB() *gorm.DB {
	dsn := "root:@/courseocg?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connected fail")
	}
	return db
}
