// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"

// 	"github.com/PuerkitoBio/goquery"
// 	"go.uber.org/zap"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type Movie struct {
// 	Name   string `json:"name"`
// 	Year   string `json:"year"`
// 	Rating string `json:"rating"`
// }

// func NewFileLogger(filepath string) (*zap.Logger, error) {
// 	cfg := zap.NewProductionConfig()
// 	if filepath != "" {
// 		cfg.OutputPaths = []string{
// 			filepath,
// 		}
// 	}
// 	return cfg.Build()
// }

// func CrawlPage1(wg sync.WaitGroup, url string, sugar *zap.SugaredLogger) {

// 	dsn := "root:anhvahieu2k@tcp(127.0.0.1:3306)/TestDB?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer wg.Done()

// 	sugar.Infof("Crawling url %s", url)
// 	resp, err := http.Get(url)
// 	if err != nil || (resp != nil && (resp.StatusCode > 299 || resp.StatusCode < 200)) {
// 		sugar.Errorw("failed to fetch URL",
// 			"url", url,
// 			"error", err,
// 		)
// 		return
// 	}
// 	// close body after reading content
// 	defer resp.Body.Close()

// 	document, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		sugar.Fatal("Error loading HTTP response body. ", err)
// 	}
// 	var movies []Movie
// 	document.Find(".lister-list tr").Each(func(index int, element *goquery.Selection) {

// 		nameFilm := element.Find(".titleColumn a").Text()
// 		// fmt.Println(nameFilm)
// 		rate := element.Find(".imdbRating strong").Text()
// 		// fmt.Println(rate)
// 		year := element.Find(".titleColumn span").Text()
// 		// fmt.Println(year)
// 		movie := Movie{Name: nameFilm, Year: year, Rating: rate}

// 		sugar.Info(zap.String("nameFilm:", nameFilm), zap.String("rate:", rate))

// 		movies = append(movies, movie)
// 		// fmt.Println(movie)
// 	})
// 	db.AutoMigrate()
// 	// db.Select("Name", "Year", "Rating").Create(&movies)
// }

// type Decorations struct {
// 	Name  string `json:"name"`
// 	Price string `json:"price"`
// }

// func CrawlPage2(wg sync.WaitGroup, url string, sugar *zap.SugaredLogger) {

// 	dsn := "root:anhvahieu2k@tcp(127.0.0.1:3306)/TestDB?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer wg.Done()

// 	sugar.Infof("Crawling url %s", url)
// 	resp, err := http.Get(url)
// 	if err != nil || (resp != nil && (resp.StatusCode > 299 || resp.StatusCode < 200)) {
// 		sugar.Errorw("failed to fetch URL",
// 			"url", url,
// 			"error", err,
// 		)
// 		return
// 	}
// 	// close body after reading content
// 	defer resp.Body.Close()

// 	document, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		sugar.Fatal("Error loading HTTP response body. ", err)
// 	}
// 	var decorations []Decorations
// 	document.Find(".collection-detail__product-details text-align-left div").Each(func(index int, element *goquery.Selection) {

// 		name := element.Find(".title d-block cl-black span").Text()
// 		// fmt.Println(name)
// 		price := element.Find(".has-text-weight-medium money money-original cl-black span").Text()
// 		// fmt.Println(price)
// 		decoration := Decorations{Name: name, Price: price}

// 		sugar.Info(zap.String("name:", name), zap.String("price:", price))

// 		decorations = append(decorations, decoration)
// 		fmt.Println(decoration)
// 	})
// 	db.AutoMigrate()
// 	// db.Select("Name", "Year", "Rating").Create(&movies)
// }

// func downloadWithGoroutinePage1(urls string, sugar *zap.SugaredLogger) {
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	go CrawlPage1(wg, urls, sugar)
// 	wg.Wait()
// }

// func downloadWithGoroutinePage2(urls string, sugar *zap.SugaredLogger) {
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)
// 	go CrawlPage2(wg, urls, sugar)
// 	wg.Wait()
// }
// func main() {
// 	// logger1, _ := NewFileLogger("A.txt")
// 	// sugar1 := logger1.Sugar()

// 	logger2, _ := NewFileLogger("B.txt")
// 	sugar2 := logger2.Sugar()

// 	// urls := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"

// 	urls2 := "https://template-homedecor.onshopbase.com/collections/new-arrivals"

// 	// downloadWithGoroutinePage1(urls, sugar1)
// 	downloadWithGoroutinePage1(urls2, sugar2)
// }

package main

import (
	"fmt"
	"hieu/crawler"
	"hieu/storage"
	"sync"
)

func downloadWithGoroutine(urls []string) (MapFilm storage.Memory_Storage_Film, MapItem storage.Memory_Storage_Item) {
	m := make(map[int64]storage.Movie)
	MapFilm.M = m

	mm := make(map[int64]storage.Item)
	MapItem.M = mm

	wg := sync.WaitGroup{}

	wg.Add(2)
	go crawler.CrawlFilm(&wg, urls[0], MapFilm)
	go crawler.CrawlItem(&wg, urls[1], MapItem)
	wg.Wait()
	return
}
func main() {

	urls := []string{"https://www.imdb.com/chart/top/?ref_=nv_mv_250",
		"https://template-homedecor.onshopbase.com/api/catalog/products_v2.json?sort_field=name&sort_direction=asc&limit=12&page=1&collection_ids=86733892580&fbclid=IwAR2aOwNN4HuCMN2V6Lv5A7MehtD47kpX9Bt4wBs--OC72vNllI5S-eRSLDw"}

	MapFilm, MapItem := downloadWithGoroutine(urls)
	fmt.Println(MapFilm)
	fmt.Println(MapItem)

}
