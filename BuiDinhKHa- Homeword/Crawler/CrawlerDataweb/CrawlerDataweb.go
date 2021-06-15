package CrawlerDataweb

import (
	"encoding/json"
	"fmt" //Thư viện để print ra màn hình
	"io"
	"log" //Thư viện để log
	"net/http"
	"strings"
	"time"

	db "gomod/databaseCrawler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
)

type Data struct { //Khởi tạo struct Data chứa dữ liệu craw
	ID     uint `gorm:"primaryKey"`
	Name   string
	Img    string
	Link   string
	Rating string
}

func Crawlerdata() {
	Database := db.Connect()
	defer Database.Close()
	if Database.HasTable(&Data{}) == false {
		Database.CreateTable(&Data{})
	}
	if Database.HasTable(&DataWEb2{}) == false {
		Database.CreateTable(&DataWEb2{})
	}
	d := make(chan Data)
	b := make(chan DataWEb2)
	quit := make(chan bool)
	go crawl(d)
	go CrawlDataweb2(b)
	go sendData(d, quit)
	go SaveDataweb2(b, quit)
	time.Sleep(5 * time.Second)
	quit <- true

}
func sendData(d chan Data, quit chan bool) {
	defer fmt.Println("successfull add data web1")
	Database := db.Connect()
	for {
		select {
		case s := <-d:
			var NewName string
			text := s.Name
			for i := 0; i < len(strings.Split(text, "\n")); i++ { // edit data input delete \n
				NewName = NewName + strings.Split(text, "\n")[i]
			}
			user := Data{
				Name:   NewName,
				Img:    s.Img,
				Link:   s.Link,
				Rating: s.Rating,
			}
			Database.Create(&user)
		case <-quit:
			fmt.Println("quit now")
			return
		}
	}
}

func crawl(d chan Data) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) { //Đang gửi request get HTML
		fmt.Printf("Visiting: %s\n", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) { //Handle error trong quá trình craw html
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) { //Sau khi đã lấy được HTML
		fmt.Printf("Visited: %s\n", r.Request.URL)
	})

	c.OnHTML(".lister-list tr", func(e *colly.HTMLElement) { //Bóc tách dữ liệu từ HTML lấy được
		data := Data{}
		data.Name = e.ChildText(".titleColumn")
		data.Link = e.ChildAttr("a", "href")
		data.Img = e.ChildAttr("img", "src")
		data.Rating = e.ChildText(".ratingColumn strong")
		d <- data

	})

	c.OnScraped(func(r *colly.Response) { //Hoàn thành job craw
		fmt.Println("Finished", r.Request.URL)
	})
	c.Visit("https://www.imdb.com/chart/top/") //Trình thu thập truy cập URL đó

}

// crawler dataweb2
type DataWEb2 struct { //Khởi tạo struct Data chứa dữ liệu craw
	ID    uint `gorm:"primaryKey"`
	Name  string
	Img   string
	Price int
}
type Image struct {
	Src string `json:"src"`
}
type Product struct {
	Title  string  `json:"title"`
	Price  int     `json:"price"`
	Images []Image `json:"images"`
}
type Products struct {
	Products []Product `json:"products"`
}

func CrawlDataweb2(b chan DataWEb2) {
	resp, err := http.Get("https://template-homedecor.onshopbase.com/api/catalog/products_v2.json")
	if err != nil {
		return
	}
	// defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	data := Products{}
	_ = json.Unmarshal(body, &data)
	for i := 0; i < len(data.Products); i++ {
		inforproduct := DataWEb2{}
		inforproduct.Name = data.Products[i].Title
		inforproduct.Price = data.Products[i].Price
		inforproduct.Img = data.Products[i].Images[0].Src
		b <- inforproduct
	}
}
func SaveDataweb2(b chan DataWEb2, quit chan bool) {
	Database := db.Connect()
	defer fmt.Println("successfull add data web2")
	for {
		select {
		case s := <-b:
			user := DataWEb2{
				Name:  s.Name,
				Img:   s.Img,
				Price: s.Price,
			}
			Database.Create(&user)
		case <-quit:
			fmt.Println("quit now")
			return
		}
	}
}
