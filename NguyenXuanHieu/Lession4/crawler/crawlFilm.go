package crawler

import (
	"fmt"
	"hieu/pkgs/connectDB"
	"hieu/pkgs/handle"
	"hieu/storage"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func CrawlFilm(wg *sync.WaitGroup, url string, m storage.Memory_Storage_Film) {
	defer wg.Done()

	db := *connectDB.ConnectToDatabase()
	defer db.Close()

	if (!db.HasTable(&storage.Movie{})) {
		db.CreateTable(&storage.Movie{})
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
		year := element.Find(".titleColumn span").Text()
		rate := element.Find(".imdbRating strong").Text()
		film := storage.Movie{
			Name:   nameFilm,
			Year:   year,
			Rating: rate,
		}
		//save to memory
		m.M[id] = film
		id++

		//save to database
		db.Create(&film)

		//save to file
		handle.WriteToFile("data/DataFilm.txt", storage.ConvertFilmToString(film))
	})
}
