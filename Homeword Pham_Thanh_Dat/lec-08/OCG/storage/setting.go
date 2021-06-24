package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/ocg?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func CreateRole() {
	db := Connect()
	if db.HasTable(&Role{}) == false {
		db.CreateTable(&Role{})
	}
}
func CreateCourse() {
	db := Connect()
	if db.HasTable(&Course{}) == false {
		db.CreateTable(&Course{})
	}
}
func CreateRoom() {
	db := Connect()
	if db.HasTable(&Room{}) == false {
		db.CreateTable(&Room{})
	}
}
func CreateLesstion() {
	db := Connect()
	if db.HasTable(&Lession{}) == false {
		db.CreateTable(&Lession{})
	}
}
func CreatePerson() {
	db := Connect()
	if db.HasTable(&Person{}) == false {
		db.CreateTable(&Person{})
	}
}
func CreateClass() {
	db := Connect()
	if db.HasTable(&Class{}) == false {
		db.CreateTable(&Class{})
	}
}
func CreateRegistration() {
	db := Connect()
	if db.HasTable(&Registration{}) == false {
		db.CreateTable(&Registration{})
	}
}
func CreateSlot() {
	db := Connect()
	if db.HasTable(&Slot{}) == false {
		db.CreateTable(&Slot{})
	}
}
