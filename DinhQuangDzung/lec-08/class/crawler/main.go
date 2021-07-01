package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type HTML struct {
	Url       string
	Content   string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Polarbear1011@/mysql_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Crawl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Couldn't get url...")
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	_ = ioutil.WriteFile("./output.html", body, 0644)
}

func main() {
	db := ConnectDB()
	defer db.Close()
	Crawl("https://vnexpress.net/nhac-si-phu-quang-vao-danh-sach-xet-giai-nha-nuoc-4298361.html")
}
