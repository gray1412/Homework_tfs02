package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tfs-02/lec-04/crawler/connectDB"
	"time"
)

type Product struct {
	Title string `json:"title"`
	Price int    `json:"price"`
	// Image    []Images   `json:"images"`
	// Variants []Variants `json:"variants"`
}

type Variants struct {
	Fulfillment_service string `json:"fulfillment_service"`
}
type Products struct {
	Products []Product `json:"products"`
}
type Images struct {
	Src string `json:"src"`
}

func CrawlProduct() {
	db := connectDB.ConnectMySQL()
	defer db.Close()
	if !db.HasTable(&Product{}) {
		db.CreateTable(&Product{})
	}
	b := make(chan Product)
	q := make(chan bool)
	go craw1(b)
	go sendData1(b, q)
	time.Sleep(time.Second * 20)
	q <- true

}

func craw1(b chan Product) {
	resp, err := http.Get("https://template-homedecor.onshopbase.com/api/catalog/products_v2.json")
	if err != nil {
		return
	}
	body, _ := io.ReadAll(resp.Body)
	data := Products{}
	_ = json.Unmarshal(body, &data)
	for i := 0; i < len(data.Products); i++ {
		product := Product{}
		product.Title = data.Products[i].Title
		product.Price = data.Products[i].Price
		// product.Variants = data.Products[i].Variants
		// product.Image = data.Products[i].Image
		b <- product
	}
}
func sendData1(b chan Product, q chan bool) {
	db := connectDB.ConnectMySQL()
	defer fmt.Println("Successfully connected to database product")
	for {
		select {
		case s := <-b:
			user := Product{
				Title: s.Title,
				Price: s.Price,
				// Variants: s.Variants,
				// Image:    s.Image,
			}
			db.Create(&user)
		case <-q:
			fmt.Println("Quit!")
			return
		}
	}
}
