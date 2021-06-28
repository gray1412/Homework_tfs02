package main

import (
	"ftxMysql/handler"
	// "ftxMysql/storage"
)

func main() {
	link := "data/train.csv"

	// url := "http://localhost:9200"
	// esclient, _ := storage.NewESCLient(url)

	// bm := storage.NewBookManager(esclient)
	

	handler.ReadFile(link)
}
