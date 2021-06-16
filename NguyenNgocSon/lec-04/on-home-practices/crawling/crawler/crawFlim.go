package crawler

import (
	"fmt"
	"net/http"
	"sync"
	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
)



//ham craw data
func Crawl(wg sync.WaitGroup, url string, sugar *zap.SugaredLogger) {

	// if wg != nil {
	defer wg.Done()
	// }

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
		fmt.Println(nameFilm)
		rate := element.Find(".imdbRating strong").Text()

		sugar.Info(zap.String("nameFilm:", nameFilm), zap.String("rate:", rate))
		fmt.Println(i)

		i++
	})
}

func CrawFilm() {

	logger, _ := NewFileLogger("./data/film.txt")
	sugar := logger.Sugar()

	urls := []string{"https://www.imdb.com/chart/top/?ref_=nv_mv_250"}

	downloadWithGoroutine(urls, sugar)
}
