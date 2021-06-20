package database

import (
	"tfs/tfs-api-mysql/storage"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func init(){
	db := *ConnectDB()
	defer db.Close()

	if (!db.HasTable(&storage.Person{})) {
		db.CreateTable(&storage.Person{})
	}
}
func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/ocg?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
