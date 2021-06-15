package Crawl

import (
	Storage "crawler/Storage"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Crawl() {
	Storage.MigrateHomedecor()

	quit := make(chan bool)
	channelHomedecor := make(chan Storage.Product)
	go crawlHomedecor(channelHomedecor)
	go insertDBHomedecor(channelHomedecor, quit)
	time.Sleep(5 * time.Second)
	quit <- true
	// quit <- true
	time.Sleep(2 * time.Second)

}
func insertDBHomedecor(channelHomedecor chan Storage.Product, quit chan bool) {
	db := Storage.Connect()
	for {
		select {
		case s := <-channelHomedecor:
			var Product = Storage.Product{Title: s.Title, Price: s.Price}
			db.Create(&Product)
			fmt.Println(s)
		case <-quit:
			// fmt.Println(<-quit)
			fmt.Println("Quiting sum func")
			return
		}
	}
}
func crawlHomedecor(channelHomedecor chan Storage.Product) {
	resp, err := http.Get("https://template-homedecor.onshopbase.com/api/catalog/products_v2.json")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	data := Storage.Products{}

	_ = json.Unmarshal(body, &data)
	for i := 0; i < len(data.Products); i++ {
		channelHomedecor <- data.Products[i]
	}
}
