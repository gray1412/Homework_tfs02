package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	Crawl("https://vnexpress.net/nhac-si-phu-quang-vao-danh-sach-xet-giai-nha-nuoc-4298361.html")
}

func Crawl(url string) {
	res, err := http.Get(url)
	if err != nil || (res != nil && (res.StatusCode > 299 || res.StatusCode < 200)) {
		fmt.Println("failed to fetch URL")
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	_ = ioutil.WriteFile("./hi.html",body,0644)
	fmt.Println(body)

}

type Blog struct {
	Id         int
	Url        string
	Content    string
	createdAt time.Time
	updatedAt time.Time
}

// resp, err := http.Get(url)
// 	if err != nil || (resp != nil && (resp.StatusCode > 299 || resp.StatusCode < 200)) {
// 		fmt.Println("failed to fetch URL")
// 		return
// 	}

// 	// close body after reading content
// 	defer resp.Body.Close()
// 	document, err := goquery.NewDocumentFromReader(resp.Body)

// 	if err != nil {
// 		fmt.Println("Error loading HTTP response body. ", err)
// 		return
// 	}

// 	var id int64 = 0
// 	document.Find(".lister-list tr").Each(func(index int, element *goquery.Selection) {
// 		nameFilm := element.Find(".titleColumn a").Text()
// 		rate := element.Find(".imdbRating strong").Text()
// 		url, _ := element.Find(".titleColumn a").Attr("href")
// 		film := storage.Film{
// 			Name: nameFilm,
// 			Rate: rate,
// 			Url:  urlRoot + url,
// 		}
// 		id++

// 		//save to database
// 		db.Create(&film)
// 	})
