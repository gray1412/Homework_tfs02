package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Review struct {
	gorm.Model
	Rate  int    `json:"rate"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
