package storage

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Role        string   `json:"role"`
	Description string   `json:"description"`
	Person      []Person `gorm:"foreignKey:RoleID;associationForeignKey:ID"`
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
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	ClassID   uint
	LessionID uint
	RoomID    uint
}
