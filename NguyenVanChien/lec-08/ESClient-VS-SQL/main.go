package main

import (
	"ex8/storage"
	"fmt"
	"time"
)

func main() {
	// link := "data/train.csv"

	// url := "http://localhost:9200"
	// esclient, _ := storage.NewESCLient(url)
	// bm := storage.NewBookManager(esclient)
	// // save to ESClient
	// handler.SaveToEsClient(link, bm)
	// // save tp SQL
	// handler.SaveToSQL(link)

	// check time find book with title is "Disapointed"
	// find in SQL
	db := storage.ConnectDatabase()
	defer db.Close()
	db.Debug().Where("title LIKE ?", "%Disapointed%").Find(&storage.Book{})

	// find in ESclient
	start := time.Now()

	url := "http://localhost:9200"
	esclient, _ := storage.NewESCLient(url)
	bm := storage.NewBookManager(esclient)
	bm.SearchBooks("isapointed")

	end := time.Since(start)
	fmt.Println(end)

}
