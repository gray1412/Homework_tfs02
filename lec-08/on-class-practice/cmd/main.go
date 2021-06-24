package main

import (
	"fmt"
	"sync"
	"tfs-02/lec-08/on-class-practice/pkg/handlers"
	"tfs-02/lec-08/on-class-practice/pkg/storage"
)

func downloadWithGoroutine(urls []string) (MapProduct storage.Memory_Storage_Product) {

	MapProduct.M = make(map[int64]storage.Product)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go handlers.Crawl(&wg, urls[0], MapProduct)
	wg.Wait()
	return
}

func main() {
	urls := []string{"https://www.imdb.com/chart/top/?ref_=nv_mv_250"}

	MapProduct := downloadWithGoroutine(urls)
	fmt.Println(MapProduct)
}
