package main

import (
	"tfs-02/lec-04/crawler/storage"
)

func main() {
	urls := []string{"https://www.imdb.com/chart/top/?ref_=nv_mv_250"}
	if !storage.Exists("crawl.txt") {
		logger, _ := storage.NewFileLogger("crawl.txt")
		sugar := logger.Sugar()
		storage.DownloadWithGoroutine(urls, sugar)
	}

}
