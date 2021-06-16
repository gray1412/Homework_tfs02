package ConnectDb

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connect Failed!")
		panic(err)
	}
	return db

}
