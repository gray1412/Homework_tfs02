package main

import (
	// storage "crawling/storage"
	"crawling/crawler"
)

func main() {

	// db := storage.Connect()
	// defer db.Close()
	crawler.CrawFilm()
	crawler.CrawArrivalsMain()
}
