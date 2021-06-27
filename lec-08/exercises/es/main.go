package main

import (
	doc "tfs-02/lec-08/exercises/es/document"
)

func main() {
	url := "http://localhost:9200"
	esclient, _ := doc.NewESClient(url)
	readFile(esclient)

}
