package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Course struct {
	ID   int
	Name string
}

type Person struct {
	ID      int
	Name    string
	RoleID  int
	Role    Role
	ClassID int
	Class   Class
}

type Role struct {
	ID   int
	Name string
}

type Class struct {
	ID       int
	PersonID int
	Person   []Person
	CourseID int
	Course   Course
	RoomID   int
	Room     Room
}

type Room struct {
	Id    int
	Floor int
}

func ConnectDB() (db *gorm.DB) {
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
