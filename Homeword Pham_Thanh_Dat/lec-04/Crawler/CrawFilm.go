package Crawl

import (
	Storage "crawler/Storage"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
	// Connect "crawler/dbConnect"
	// _ "github.com/go-sql-driver/mysql"
)

func CrawFilm() {
	Storage.MigrateFilm()
	quit := make(chan bool)
	channel := make(chan Storage.Imdb)
	go crawlImdb(channel)
	go insertDBImdb(channel, quit)
	time.Sleep(5 * time.Second)
	quit <- true
	// quit <- true
	time.Sleep(2 * time.Second)
}

func insertDBImdb(channel chan Storage.Imdb, quit chan bool) {
	db := Storage.Connect()
	for {
		select {
		case s := <-channel:
			var text string
			a := s.Name
			for i := 0; i < len(strings.Split(a, "\n")); i++ {
				text = text + strings.Split(a, "\n")[i]
			}
			var Imdb = Storage.Imdb{Name: text, Rate: s.Rate, Link: "https://www.imdb.com/" + s.Link}
			db.Create(&Imdb)
		case <-quit:
			// fmt.Println(<-quit)
			fmt.Println("Quiting sum func")
			return
		}
	}
}
func crawlImdb(channel chan Storage.Imdb) {

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting: %s\n", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Visited: %s\n", r.Request.URL)
	})
	c.OnHTML(".lister-list tr", func(e *colly.HTMLElement) {
		Imdb := Storage.Imdb{}
		Imdb.Name = e.ChildText(".titleColumn")
		Imdb.Rate = e.ChildText(".ratingColumn strong")
		Imdb.Link = e.ChildAttr("a", "href")
		channel <- Imdb

	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")
}
