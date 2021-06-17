package crawler

import (
	"fmt"
	"net/http"
	"sync"
	"tfs-02/crawl/pkgs/connectDB"
	"tfs-02/crawl/pkgs/handleFile"
	"tfs-02/crawl/storage"

	"github.com/PuerkitoBio/goquery"
)

func CrawlYouTube(wg *sync.WaitGroup, url string, m storage.Memory_Storage_YouTube) {

	defer wg.Done()

	db := *connectDB.ConnectToDatabase()
	defer db.Close()

	if (!db.HasTable(&storage.YouTube{})) {
		db.CreateTable(&storage.YouTube{})
	}

	resp, err := http.Get(url)
	if err != nil || (resp != nil && (resp.StatusCode > 299 || resp.StatusCode < 200)) {
		fmt.Println("failed to fetch URL Youtube")
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
	document.Find("#contents ytd-rich-item-renderer").Each(func(index int, element *goquery.Selection) {
		name_video := element.Find("#video-title").Text()
		name_channel := element.Find("#text a").Text()
		view := element.Find("#metadata-line span").Text()
		youtube := storage.YouTube{
			Name_video:   name_video,
			Name_channel: name_channel,
			View:         view,
		}
		//save to memory
		m.M[id] = youtube
		id++

		//save to database
		db.Create(&youtube)

		//save to file
		handleFile.WriteToFile("dataCraw/DataYouTubeCraw.txt", storage.ConvertYouTubeToString(youtube))
	})
}
