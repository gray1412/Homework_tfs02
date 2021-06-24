package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Price       float64   `  json:"price"`
	Sale        int       `json:"sale"`
	Quantity    int       ` json:"qty"`
	Weight      float64   `json:"weight"`
	Description string    `json:"desc"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time ` json:"updated_at"`
	IsRendered  bool      `json:"is_rendered"`
	BrandID     int       `json:"brand_id"`
}
type Brand struct {
	ID       int       `  json:"id" gorm:"primaryKey"`
	Name     string    ` json:"name"`
	Products []Product `json:"products" gorm:"foreignKey:BrandID;associationForeignKey:ID"`
}

func ConnectToDatabase() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:123456@/ocg_tfs_02?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database") // Kiểm tra kết nối tới databse
	}
	return
}

func main() {
	db := *ConnectToDatabase()
	defer db.Close()

	if (!db.HasTable(&Product{})) { //neu khong ton tai bang
		db.CreateTable(&Product{}) // thi tao bang
	} else {
		db.DropTable(&Product{})   // neu da ton tai bang thi xoa bang
		db.CreateTable(&Product{}) // sau do tao lai bang
	}

	if (!db.HasTable(&Brand{})) { //neu khong ton tai bang
		db.CreateTable(&Brand{}) // thi tao bang
	} else {
		db.DropTable(&Brand{})   // neu da ton tai bang thi xoa bang
		db.CreateTable(&Brand{}) // sau do tao lai bang
	}
}
