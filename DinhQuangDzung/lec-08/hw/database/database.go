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

type Role struct {
	gorm.Model
	Role        string   `json:"role"`
	Description string   `json:"description"`
	Person      []Person `gorm:"foreignKey:RoleID"`
}
type Person struct {
	gorm.Model
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	Address      string         `json:"address"`
	Email        string         `json:"email"`
	Phone        int            `json:"phone"`
	Age          int            `json:"age"`
	RoleID       int            `json:"role"`
	Registration []Registration `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Course struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Class       []Class `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Class struct {
	gorm.Model
	Name         string `json:"name"`
	Description  string `json:"description"`
	CourseID     uint
	Registration []Registration `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Slot         []Slot         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Registration struct {
	gorm.Model
	PersonID uint
	ClassID  uint
}
type Lession struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	CourseID    uint
	Slot        []Slot `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Room struct {
	gorm.Model
	Floor    int `json:"floor"`
	CourseID uint
	Slot     []Slot `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Slot struct {
	gorm.Model
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	ClassID   uint
	LessionID uint
	RoomID    uint
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
