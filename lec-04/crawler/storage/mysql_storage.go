package storage

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type data_film struct {
	ID     int    `db:"id"`
	Poster string `db:"poster"`
	Name   string `db:"name"`
	Year   string `db:"year"`
	Rate   string `db:"rate"`
}

// func SaveDB() {

// 	// db:=*pkg.ConnectMySQL()
// 	if !db.HasTable(&data_film{}) {
// 		db.CreateTable(&data_film{})
// 	}
// 	fmt.Printf("%v", db)
// }

func SaveDB(wg *sync.WaitGroup, url string, sugar *zap.SugaredLogger) (films []data_film) {
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"123456",
		"tcp",
		"localhost",
		"3306",
		"crawl",
	)
	db, err := gorm.Open("mysql", mysqlCredentials)

	if err != nil {
		panic("Failed to connect database") // Kiểm tra kết nối tới database
	}
	defer db.Close()

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
	var i int = 0
	document.Find(".lister-list tr").Each(func(index int, element *goquery.Selection) {
		posterFilm, _ := element.Find(".posterColumn img").Attr("src")
		nameFilm := element.Find(".titleColumn a").Text()
		yearFilm := element.Find(".titleColumn span").Text()
		rate := element.Find(".imdbRating strong").Text()
		films[i] = data_film{ID: i, Poster: posterFilm, Name: nameFilm, Year: yearFilm, Rate: rate}
		// sugar.Info(zap.String("imgSrc:", posterFilm), zap.String("nameFilm:", nameFilm), zap.String("yearFilm:", yearFilm), zap.String("rate:", rate))
		i++
	})
	return films
}
