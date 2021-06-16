package crawler

import (
	"fmt"
	"sync"
	"go.uber.org/zap"
)

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

func downloadWithGoroutine(urls []string, sugar *zap.SugaredLogger) {
	wg := sync.WaitGroup{}
	for _, url := range urls {
		fmt.Println("crawling ", url)
		wg.Add(1)
		go Crawl(wg, url, sugar)
	}
	wg.Wait()
}

