package main

import (
	"fmt"
	"sync"

	"tfs-02/crawl/crawler"
	"tfs-02/crawl/storage"
)

func downloadWithGoroutine(urls []string) (MapFilm storage.Memory_Storage_Film, MapItem storage.Memory_Storage_Item, MapYouTube storage.Memory_Storage_YouTube) {

	MapFilm.M = make(map[int64]storage.Film)
	MapItem.M = make(map[int64]storage.Item)
	MapYouTube.M = make(map[int64]storage.YouTube)

	wg := sync.WaitGroup{}

	wg.Add(3)
	go crawler.CrawlFilm(&wg, urls[0], MapFilm)
	go crawler.CrawlItem(&wg, urls[1], MapItem)
	go crawler.CrawlYouTube(&wg, urls[2], MapYouTube)
	wg.Wait()
	return
}
func main() {

	urls := []string{"https://www.imdb.com/chart/top/?ref_=nv_mv_250",
		"https://template-homedecor.onshopbase.com/api/catalog/products_v2.json?sort_field=name&sort_direction=asc&limit=12&page=1&collection_ids=86733892580&fbclid=IwAR2aOwNN4HuCMN2V6Lv5A7MehtD47kpX9Bt4wBs--OC72vNllI5S-eRSLDw",
		"https://www.youtube.com/"}

	MapFilm, MapItem, MapYouTube := downloadWithGoroutine(urls)
	fmt.Println(MapFilm)
	fmt.Println(MapItem)
	fmt.Println(MapYouTube)

}
