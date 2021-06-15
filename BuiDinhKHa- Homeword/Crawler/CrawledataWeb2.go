package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

func Crawl() {
	resp, err := http.Get("https://template-homedecor.onshopbase.com/api/catalog/products_v2.json")
	if err != nil {
		return
	}
	// defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	data := Products{}
	_ = json.Unmarshal(body, &data)
	for i := 0; i < len(data.Products); i++ {
		fmt.Println(data.Products[i].Title)
		fmt.Println(data.Products[i].Price)
		fmt.Println(data.Products[i].Images[0].Src)
	}
}
func main() {
	Crawl()
}
