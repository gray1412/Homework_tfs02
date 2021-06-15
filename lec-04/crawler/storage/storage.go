package storage

import (
	"fmt"
	"os"
	"sync"
	"tfs-02/lec-04/crawler/crawler"

	"go.uber.org/zap"
)

//kiem tra ton tai
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//hàm tạo file chứa logger
func NewFileLogger(filepath string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	if filepath != "" {
		cfg.OutputPaths = []string{
			filepath,
		}
	}
	return cfg.Build()
}

func DownloadWithGoroutine(urls []string, sugar *zap.SugaredLogger) {
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		fmt.Println("crawling ", url)
		go crawler.Crawl(&wg, url, sugar)
	}
	wg.Wait()
}
