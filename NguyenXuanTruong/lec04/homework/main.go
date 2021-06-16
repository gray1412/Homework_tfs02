package main

import (
	"log"
	"time"

	C "Crawl/crawler"
	S "Crawl/storage"
)

func main() {
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
}
