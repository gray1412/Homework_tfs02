package crawler

import (
	"fmt"
	"net/http"
	"sync"

	"crawling/pkgs/connectDB"
	"crawling/storage"

	"github.com/PuerkitoBio/goquery"
)

func CrawlFilm(wg *sync.WaitGroup, url string) {
	urlRoot := "https://www.imdb.com/"
	defer wg.Done()

	db := *connectDB.ConnectDB()
	defer db.Close()

	if (!db.HasTable(&storage.Film{})) {
		db.CreateTable(&storage.Film{})
	}

	resp, err := http.Get(url)
	if err != nil || (resp != nil && (resp.StatusCode > 299 || resp.StatusCode < 200)) {
		fmt.Println("failed to fetch URL")
		return
	}

	// close body after reading content
	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error loading HTTP response body. ", err)
		return
	}

	var id int64 = 0
	document.Find(".lister-list tr").Each(func(index int, element *goquery.Selection) {
		nameFilm := element.Find(".titleColumn a").Text()
		rate := element.Find(".imdbRating strong").Text()
		url, _ := element.Find(".titleColumn a").Attr("href")
		film := storage.Film{
			Name: nameFilm,
			Rate: rate,
			Url:  urlRoot + url,
		}
		id++

		//save to database
		db.Create(&film)
	})
}
