package main

import (
	"fmt"
	"sync"

	"chien/crawler"

	"chien/storage"

	"go.uber.org/zap"
)

func DownloadWithGoroutineFilm(url string, sugar *zap.SugaredLogger) map[int64]storage.Film {
	m := make(map[int64]storage.Film)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go crawler.CrawlFilm(&wg, url, sugar, &m)
	wg.Wait()
	return m
}

func DownloadWithGoroutineItem(url string, sugar *zap.SugaredLogger) map[int64]storage.Item {
	m := make(map[int64]storage.Item)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go crawler.CrwalItem(&wg, url, sugar, &m)
	wg.Wait()
	return m
}

func main() {

	logger1, _ := storage.NewFileLogger("DataFilm.txt")
	sugar1 := logger1.Sugar()

	urlFilm := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	mmapFilm := make(map[int64]storage.Film)
	mmapFilm = DownloadWithGoroutineFilm(urlFilm, sugar1)
	fmt.Println(mmapFilm)

	logger2, _ := storage.NewFileLogger("DataItem.txt")
	sugar2 := logger2.Sugar()
	urlItem := "https://template-homedecor.onshopbase.com/collections/new-arrivals"
	mapItem := make(map[int64]storage.Item)
	mapItem = DownloadWithGoroutineItem(urlItem, sugar2)
	fmt.Print(mapItem)
}
