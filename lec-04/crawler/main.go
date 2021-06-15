package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
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

func Crawl(wg *sync.WaitGroup, url string, sugar *zap.SugaredLogger) {

	if wg != nil {
		defer wg.Done()
	}

	sugar.Infof("Crawling url %s", url)
	resp, err := http.Get(url)
	if err != nil || (resp != nil && (resp.StatusCode > 299 || resp.StatusCode < 200)) {
		sugar.Errorw("failed to fetch URL",
			"url", url,
			"error", err,
		)
		return
	}
	// close body after reading content
	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		sugar.Fatal("Error loading HTTP response body. ", err)
	}


	document.Find(".lister-list tr").Each(func(index int, element *goquery.Selection) {
		posterFilm, exists := element.Find(".posterColumn img").Attr("src")
		nameFilm := element.Find(".titleColumn a").Text()
		yearFilm := element.Find(".titleColumn span").Text()
		rate := element.Find(".imdbRating strong").Text()
		if exists {
			sugar.Info(zap.String("imgSrc:", posterFilm), zap.String("nameFilm:", nameFilm), zap.String("yearFilm:", yearFilm), zap.String("rate:", rate))
		}

	})
}

func downloadWithGoroutine(urls []string, sugar *zap.SugaredLogger) {
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		fmt.Println("crawling ", url)
		go Crawl(&wg, url, sugar)
	}
	wg.Wait()
}

func main() {

	logger, _ := NewFileLogger("D.txt")
	sugar := logger.Sugar()

	urls := []string{"https://www.imdb.com/chart/top/?ref_=nv_mv_250"}

	downloadWithGoroutine(urls, sugar)
}
