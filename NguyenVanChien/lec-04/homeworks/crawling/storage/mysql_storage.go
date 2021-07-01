package storage

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

type Film struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Rate string `db:"rate"`
}

type Item struct {
	Title string `db:"title"`
	Price string `db:"price"`
}

type ImageCraw struct {
	Src    string `json:"src"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}
type ItemCraw struct {
	ID    int64       `json:"id"`
	Title string      `json:"title"`
	Img   []ImageCraw `json:"images"`
	Price int64       `json:"price"`
}

type ListItemsCraw struct {
	ListItemsCraw []ItemCraw `json:"items"`
}

func ConvertIteamCrawltoItem(itemCraw ItemCraw) (item Item) {
	item.Title = itemCraw.Title
	item.Price = string(itemCraw.Price)
	return
}

func ConnectDatabase() (db *gorm.DB) {
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"",
		"tcp",
		"localhost",
		"3306",
		"testcrawl",
	)
	db, err := gorm.Open("mysql", mysqlCredentials) // Mở database

	if err != nil {
		panic("Failed to connect database") // Kiểm tra kết nối tới database
	}
	return
}
