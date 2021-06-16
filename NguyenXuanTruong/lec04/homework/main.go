package main

import (
	"log"
	"time"

	C "Crawl/crawler"
	S "Crawl/storage"
)

func main() {
	// Crawl Film from IMDb
	ch := make(chan S.Film)
	db, err := S.DBConnection() //Khởi tạo biến conection
	if err != nil {             //Catch error trong quá trình thực thi
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()

	log.Printf("Successfully connected to database")

	go C.CrawlerFilm(ch) //Thực thi crawl
	go S.Insert2DatabaseFilm(db, ch)
	time.Sleep(10 * time.Second)

	// Crawl HomeDecor from Shopbase
	ch2 := make(chan S.HomeDecor)
	db2, err := S.DBConnection()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db2.Close()

	go C.CrawlHomeDecor(ch2)
	go S.Insert2DatabaseHomeDecor(db, ch2)
	time.Sleep(10 * time.Second)
}
