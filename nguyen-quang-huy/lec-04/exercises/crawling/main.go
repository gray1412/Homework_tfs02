package main

import (
	"fmt"
	"sync"

	"crawling/crawler"
)

func downloadWithGoroutine(url string) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go crawler.CrawlFilm(&wg, url)
	wg.Wait()
}
func main() {

	url := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	fmt.Println("Crawler ", url)
	downloadWithGoroutine(url)
	fmt.Println("Done!")

}
