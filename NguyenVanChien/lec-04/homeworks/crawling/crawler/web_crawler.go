package crawler

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"chien/storage"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func CrawlFilm(wg *sync.WaitGroup, url string, sugar *zap.SugaredLogger, m *map[int64]storage.Film) {

	defer wg.Done()
	sugar.Infof("Crawling url %s", url)
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
	db := *storage.ConnectDatabase()
	defer db.Close()

	if (!db.HasTable(&storage.Film{})) {
		db.CreateTable(&storage.Film{})
	} else {
		db.DropTable(&storage.Film{})
		db.CreateTable(&storage.Film{})
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		sugar.Fatal("Error loading HTTP response body. ", err)
	}
	var i int64 = 1
	filmMap := *m
	document.Find(".lister-list tr").Each(func(index int, element *goquery.Selection) {
		nameFilm := element.Find(".titleColumn a").Text()
		rate := element.Find(".imdbRating strong").Text()
		film := storage.Film{
			ID:   i,
			Name: nameFilm,
			Rate: rate,
		}
		db.Create(&film)
		filmMap[i] = film
		sugar.Info(zap.String("nameFilm:", nameFilm), zap.String("rate:", rate))
		i++
	})
}

func CrwalItem(wg *sync.WaitGroup, url string, sugar *zap.SugaredLogger, m *map[int64]storage.Item) {
	defer wg.Done()
	sugar.Infof("Crawling url %s", url)
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
	db := *storage.ConnectDatabase()
	defer db.Close()

	if (!db.HasTable(&storage.Item{})) {
		db.CreateTable(&storage.Item{})
	} else {
		db.DropTable(&storage.Item{})
		db.CreateTable(&storage.Item{})
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	datas := storage.ListItemsCraw{}
	_ = json.Unmarshal(body, &datas)

	var i int64 = 1
	Itemmap := *m
	for _, itemcrawl := range datas.ListItemsCraw {

		item := storage.ConvertIteamCrawltoItem(itemcrawl)
		Itemmap[i] = item
		db.Create(&item)
		sugar.Info(zap.String("title:", item.Title), zap.String("price:", item.Price))
		i++
	}

}
