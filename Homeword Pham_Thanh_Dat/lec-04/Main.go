package main

import (
	Craw "crawler/Crawler"
	Storage "crawler/Storage"
)

func main() {
	db := Storage.Connect()
	defer db.Close()
	// Craw.CrawFilm()
	Craw.Crawl()
}
