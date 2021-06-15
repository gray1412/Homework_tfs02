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
	var i int64 = 0
	document.Find(".lister-list tr").Each(func(index int, element *goquery.Selection) {
		nameFilm := element.Find(".titleColumn a").Text()
		yearFilm := element.Find(".titleColumn span").Text()
		rate := element.Find(".imdbRating strong").Text()
		sugar.Info(zap.String("nameFilm:", nameFilm), zap.String("yearFilm:", yearFilm), zap.String("rate:", rate))
		fmt.Println(i)
		i++
	})
}
func downloadWithGoroutine(urls []string, sugar *zap.SugaredLogger) {
	wg := sync.WaitGroup{}
	for _, url := range urls {
		fmt.Println("crawling ", url)
		wg.Add(1)
		go Crawl(&wg, url, sugar)
	}
	wg.Wait()
}

func main() {

	logger, _ := NewFileLogger("D.txt")
	sugar := logger.Sugar()

	urls := []string{"https://www.imdb.com/chart/top/?ref_=nv_mv_250",
		"https://www.imdb.com/chart/moviemeter?pf_rd_m=A2FGELUUNOQJNL&pf_rd_p=4da9d9a5-d299-43f2-9c53-f0efa18182cd&pf_rd_r=7MX3QNK9X7EJXGPMK7GR&pf_rd_s=right-4&pf_rd_t=15506&pf_rd_i=top&ref_=chttp_ql_2"}

	downloadWithGoroutine(urls, sugar)
}
