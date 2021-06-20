package Storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/studentmanage?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func CreateStudents() {
	db := Connect()
	if db.HasTable(&Students{}) == false {
		db.CreateTable(&Students{})
	}
}
