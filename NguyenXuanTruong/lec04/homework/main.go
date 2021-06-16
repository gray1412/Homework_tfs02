package main

import (
	Craw "Crawl/crawler"
	Storage "Crawl/storage"
)

func main() {
	db := Storage.Connect()
	defer db.Close()
	// Craw.CrawFilm()
	Craw.Crawl()
}
