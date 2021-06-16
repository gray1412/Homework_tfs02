package crawler

import (
	S "Crawl/storage"
	"encoding/json"
	"io"
	"net/http"
)

func CrawlHomeDecor(ch chan S.HomeDecor) {
	resp, err := http.Get("https://template-homedecor.onshopbase.com/api/catalog/products_v2.json?sort_field=name&sort_direction=asc&limit=12&page=1&collection_ids=86733892580&fbclid=IwAR2aOwNN4HuCMN2V6Lv5A7MehtD47kpX9Bt4wBs--OC72vNllI5S-eRSLDw")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	data := S.HomeDecors{}

	_ = json.Unmarshal(body, &data)

	//output, _ := json.Marshal(data.HomeDecors)
	for i := 0; i < len(data.HomeDecors); i++ {
		ch <- data.HomeDecors[i]
	}
}
