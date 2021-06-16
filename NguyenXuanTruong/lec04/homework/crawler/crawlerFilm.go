package crawler

import (
	"fmt"
	"log"

	S "Crawl/storage"
	//Thao tác với SQL
	//Xử lý thời gian
	//_ "github.com/go-sql-driver/mysql" //Tạo driver kết nối mysql
	"github.com/gocolly/colly"
)

func CrawlerFilm(ch chan S.Film) {
	c := colly.NewCollector()

	//Đang gửi request get HTML
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	//Handle error trong quá trình craw html
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) { //Sau khi đã lấy được HTML
		fmt.Printf("Visited: %sn", r.Request.URL)
	})

	c.OnHTML(".lister-list tr", func(e *colly.HTMLElement) {
		f := S.Film{}
		f.Name = e.ChildText(".titleColumn")         // Tìm đến thẻ con .titleColumn và lấy nội dung
		f.Image = e.ChildAttr(" img", "src")         // Tìm đến thẻ con .posterColumn img và lấy nội dung trong "src"
		f.Rate = e.ChildText(".ratingColumn strong") // Tìm đến thẻ con .ratingColumn có thuộc tính strong và lấy nội dung
		f.Url = e.ChildAttr(" a", "href")            // Tìm đến thẻ con .titleColumn a và lấy nội dung trong "href"
		fmt.Printf("Name : %s \n Image : %s \n Rate : %s \n URL : %s \n", f.Name, f.Image, f.Rate, f.Url)
		ch <- f
	})

	// Thông báo thực hiện xong crawl
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished.", r.Request.URL)
	})

	// Truy cập URL
	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")
}
