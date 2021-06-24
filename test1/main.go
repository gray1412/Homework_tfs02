package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Sale        int       `json:"sale"`
	Quantity    int       ` json:"qty"`
	Weight      float64   `json:"weight"`
	Description string    ` json:"desc"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time ` json:"updated_at"`
	IsRendered  bool      `json:"is_rendered"`
	BrandID     int       `json:"brand_id"`
}
type Brand struct {
	ID       int       ` json:"id" gorm:"primaryKey"`
	Name     string    ` json:"name"`
	Products []Product `json:"products" gorm:"foreignKey:BrandID;associationForeignKey:ID"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_crawl?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error when connect to db ", err)
		return
	}
	err = db.AutoMigrate(&Brand{}, &Product{})
	if err != nil {
		log.Fatal("error when auto migrate table ", err)
	}
}
