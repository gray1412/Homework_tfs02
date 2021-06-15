package main

import (
	"fmt"
	"net/http"
	"sync"
	"tfs-02/lec-04/crawler/storage"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Film struct {
	ID     int    `db:"id"`
	Poster string `db:"poster"`
	Name   string `db:"name"`
	Year   string `db:"year"`
	Rate   string `db:"rate"`
}

func NewFileLogger(filepath string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	if filepath != "" {
		cfg.OutputPaths = []string{
			filepath,
		}
	}
	return cfg.Build()
}

func main() {
	urls := []string{"https://www.imdb.com/chart/top/?ref_=nv_mv_250"}
	logger, _ := storage.NewFileLogger("testcrawl.txt")
	sugar := logger.Sugar()
	DownloadWithGoroutine(urls, sugar)

	db := *ConnectDatabase()
	defer db.Close()
	var films []Film
	db.Find(&films)
	time.Sleep(time.Second * 10)
}

func ConnectDatabase() (db *gorm.DB) {
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
	// defer db.Close()
	if err != nil {
		panic("Failed to connect database") // Kiểm tra kết nối tới database
	}
	return
}

func Crawl(wg sync.WaitGroup, url string, sugar *zap.SugaredLogger) {

	defer wg.Done()
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

	db := *ConnectDatabase()
	defer db.Close()

	var i int = 1
	document.Find(".lister-list tr").Each(func(index int, element *goquery.Selection) {
		posterFilm, _ := element.Find(".posterColumn img").Attr("src")
		nameFilm := element.Find(".titleColumn a").Text()
		yearFilm := element.Find(".titleColumn span").Text()
		rate := element.Find(".imdbRating strong").Text()
		film := Film{ID: i, Poster: posterFilm, Name: nameFilm, Year: yearFilm, Rate: rate}
		//ghi vao file txt
		sugar.Info(zap.String("posterFilm:", posterFilm), zap.String("nameFilm:", nameFilm), zap.String("yearFilm:", yearFilm), zap.String("rate:", rate))
		db.Create(&film)
		i++
	})
}

func DownloadWithGoroutine(urls []string, sugar *zap.SugaredLogger) {
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go Crawl(wg, url, sugar)
	}
	wg.Wait()
}
