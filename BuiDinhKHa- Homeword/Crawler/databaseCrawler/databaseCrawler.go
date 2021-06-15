package databaseCrawler

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/crawler?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to the database!")
	}
	return db

}
func main() {
	fmt.Println("strart database...")
}
