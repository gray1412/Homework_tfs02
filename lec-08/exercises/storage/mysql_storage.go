package storage

import "github.com/jinzhu/gorm"

type People struct {
	gorm.Model
	Name      string `json:"name" gorm:"index"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"addr"`
	RoleRefer uint   `json:"roleId"`
}

type PersonRole struct {
	RoleId      uint     `gorm:"primaryKey"`
	PersonRole  string   `json:"role"`
	Description string   `json:"desc"`
	People      []People `json:"people" gorm:"foreignKey:RoleRefer;associationForeignKey:RoleId"`
}

type Room struct {
	RoomId uint `gorm:"primaryKey"`
	Floor  int  `json:"floor"`
}
