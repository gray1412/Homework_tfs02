package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Role struct {
	gorm.Model
	Role        string `json:"role"`
	Description string `json:"description"`
}
type Person struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Age       int    `json:"age"`
	Role      []Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Course struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Class struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Course      []Course `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Registration struct {
	gorm.Model
	Person []Person `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Class  []Class  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Lession struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Course      []Course `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Room struct {
	gorm.Model
	Floor  int      `json:"floor"`
	Course []Course `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Slot struct {
	gorm.Model
	Name      string    `json:"name"`
	StartTime string    `json:"startTime"`
	EndTime   string    `json:"endTime"`
	Class     []Class   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Lession   []Lession `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Room      []Room    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
