package crawler

import (
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
)

// type filmInfo struct {
// 	ID     int    `db:"id"`
// 	Poster string `db:"poster"`
// 	Name   string `db:"name"`
// 	Year   string `db:"year"`
// 	Rate   string `db:"rate"`
// }

func Crawl(wg *sync.WaitGroup, url string, sugar *zap.SugaredLogger) {

	if wg != nil {
		defer wg.Done()
	}

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
	var i int = 1
	document.Find(".lister-list tr").Each(func(index int, element *goquery.Selection) {
		posterFilm, _ := element.Find(".posterColumn img").Attr("src")
		nameFilm := element.Find(".titleColumn a").Text()
		yearFilm := element.Find(".titleColumn span").Text()
		rate := element.Find(".imdbRating strong").Text()
		// film := filmInfo{ID: i, Poster: posterFilm, Name: nameFilm, Year: yearFilm, Rate: rate}
		sugar.Info(zap.String("imgSrc:", posterFilm), zap.String("nameFilm:", nameFilm), zap.String("yearFilm:", yearFilm), zap.String("rate:", rate))
		i++
	})
}
