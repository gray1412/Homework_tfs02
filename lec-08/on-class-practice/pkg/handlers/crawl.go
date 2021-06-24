package handlers

import (
	"fmt"
	"net/http"
	"sync"
	"tfs-02/lec-08/on-class-practice/pkg/storage"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Crawl(wg *sync.WaitGroup, url string, m storage.Memory_Storage_Product) {
	defer wg.Done()

	db := *storage.ConnectToDatabase()
	defer db.Close()

	if (!db.HasTable(&storage.Product{})) {
		db.CreateTable(&storage.Product{})
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

	var id int64 = 1
	document.Find(".lister-list tr").Each(func(index int, element *goquery.Selection) {
		nameProduct := element.Find(".titleColumn a").Text()
		img_src := element.Find(".imdbRating strong").Text()
		dt := time.Now()
		product := storage.Product{
			Id:         uint(id),
			Name:       nameProduct,
			Img_src:    img_src,
			Created_at: dt.String(),
		}
		//save to memory
		m.M[id] = product
		id++

		//save to database
		db.Create(&product)

		//save to file
		WriteToFile("repository/text/DataProduct.txt", storage.ConvertProductToString(product))
	})
}
