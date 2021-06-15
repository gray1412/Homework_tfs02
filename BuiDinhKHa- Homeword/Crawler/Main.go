package main

import (
	//Thư viện để print ra màn hình
	//Thư viện để log

	"fmt"
	db "gomod/CrawlerDataweb"
)

func main() {
	db.Crawlerdata()
	fmt.Println("Finish process !!")

}
