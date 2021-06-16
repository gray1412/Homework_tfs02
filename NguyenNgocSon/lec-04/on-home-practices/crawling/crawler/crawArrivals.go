package crawler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Title string `json:"title"`
	Price int    `json:"price"`
}

type Products struct {
	Products []Product `json:"products"`
}

//SDHBFHSFHBEWS
func CrawlArivals() {
	resp, err := http.Get("https://template-homedecor.onshopbase.com/api/catalog/products_v2.json?sort_field=name&sort_direction=asc&limit=12&page=1&collection_ids=86733892580&fbclid=IwAR2aOwNN4HuCMN2V6Lv5A7MehtD47kpX9Bt4wBs--OC72vNllI5S-eRSLDw")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	data := Products{}

	_ = json.Unmarshal(body, &data)

	output, _ := json.Marshal(data.Products)
	_ = ioutil.WriteFile("./data/arrivals.txt", output, 0644)

}

func CrawArrivalsMain() {
	CrawlArivals()
}
