package crawler

import (
	"fmt"
	"log"
	"time"

	"tfs-02/lec-04/crawler/connectDB"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
)

type Film struct {
	Id    uint `gorm:"primaryKey"`
	Title string
	// Year  string
	Link  string
	Rate  string
	Image string
}

func CrawlFilm() {
	db := connectDB.ConnectMySQL()
	// defer db.Close()
	if !db.HasTable(&Film{}) {
		db.CreateTable(&Film{})
	}
	d := make(chan Film)
	q := make(chan bool)
	go crawl(d)
	go sendData(d, q)
	time.Sleep(5 * time.Second)
	q <- true
	log.Println("Successfully connected to database Film!")

}
func sendData(d chan Film, q chan bool) {
	db := connectDB.ConnectMySQL()
	for {
		select {
		case s := <-d:
			user := Film{
				Title: s.Title,
				// Year:  s.Year,
				Link:  "https://www.imdb.com/" + s.Link,
				Rate:  s.Rate,
				Image: s.Image,
			}
			db.Create(&user)
		case <-q:
			fmt.Println("Quit!")
			return
		}
	}

}
func crawl(d chan Film) {
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
		film := Film{}
		film.Title = e.ChildText(".titleColumn a")
		// film.Year = e.ChildText(".titleColumn a")
		film.Link = e.ChildAttr("a", "href")
		film.Image = e.ChildAttr("img", "src")
		film.Rate = e.ChildText(".ratingColumn strong")
		d <- film
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")
}
