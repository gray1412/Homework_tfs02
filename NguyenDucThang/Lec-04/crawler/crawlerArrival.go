package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	CrawlerF "crawler/ConnectDb"
)

type ProductArrival struct {
	ID     uint `gorm:"primaryKey"`
	Handle string
	Name   string
	Image  string
	Price  int
}
type Img struct {
	Src string `json:"src"`
}
type Product struct {
	Handle string `json:"handle"`
	Title  string `json:"title"`
	Price  int    `json:"price"`
	Images []Img  `json:"images"`
}
type Products struct {
	Products []Product `json:"products"`
}

func CrawlArrival() {
	fmt.Println("Start................")
	Database := CrawlerF.ConnectDb()
	defer Database.Close()
	if Database.HasTable(&ProductArrival{}) == false {
		Database.CreateTable(&ProductArrival{})
	}
	b := make(chan ProductArrival)
	q := make(chan bool)
	go craw(b)
	go sendData1(b, q)
	time.Sleep(time.Second * 20)
	q <- true

}

func craw(b chan ProductArrival) {
	resp, err := http.Get("https://template-homedecor.onshopbase.com/api/catalog/products_v2.json")
	if err != nil {
		return
	}
	body, _ := io.ReadAll(resp.Body)
	data := Products{}
	_ = json.Unmarshal(body, &data)
	for i := 0; i < len(data.Products); i++ {
		pArrival := ProductArrival{}
		pArrival.Handle = data.Products[i].Handle
		pArrival.Name = data.Products[i].Title
		pArrival.Price = data.Products[i].Price
		pArrival.Image = data.Products[i].Images[0].Src
		b <- pArrival
	}
}
func sendData1(b chan ProductArrival, q chan bool) {
	db := CrawlerF.ConnectDb()
	defer fmt.Println("Successfully connected to database Arrival")
	for {
		select {
		case s := <-b:
			user := ProductArrival{
				Handle: s.Handle,
				Name:   s.Name,
				Image:  s.Image,
				Price:  s.Price,
			}
			db.Create(&user)
		case <-q:
			fmt.Println("Quit!")
			return
		}
	}
}
