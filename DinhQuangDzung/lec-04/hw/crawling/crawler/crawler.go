package crawler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"lec-04/hw/crawling/storage"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var db = storage.ConnectDB()

func CrawlShop(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	data := storage.Products{}

	_ = json.Unmarshal(body, &data)
	for _, v := range data.Products {
		db.Create(v)
	}
	output, _ := json.Marshal(data.Products)
	_ = ioutil.WriteFile("./storage/shop_output.json", output, 0644)
}

func CrawlIMDB(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	movies := []storage.Movie{}

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find(".lister-list tr").Each(func(index int, element *goquery.Selection) {
		title := element.Find(".titleColumn a").Text()
		ranking := element.Find(".ratingColumn strong").Text()
		movie := storage.Movie{
			Name:    title,
			Ranking: ranking,
		}
		db.Create(movie)
		movies = append(movies, movie)
	})

	data, _ := json.Marshal(movies)
	_ = ioutil.WriteFile("./storage/imdb_output.json", data, 0644)
}

func RunCrawl() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go CrawlIMDB("https://www.imdb.com/chart/top/?ref_=nv_mv_250", &wg)
	go CrawlShop("https://template-homedecor.onshopbase.com/api/catalog/products_v2.json", &wg)
	wg.Wait()
}
