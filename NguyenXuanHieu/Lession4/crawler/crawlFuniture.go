package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"hieu/pkgs/connectDB"
	"hieu/pkgs/handle"
	"hieu/storage"
)

func CrawlItem(wg *sync.WaitGroup, url string, m storage.Memory_Storage_Item) {
	defer wg.Done()

	db := *connectDB.ConnectToDatabase()
	defer db.Close()

	if (!db.HasTable(&storage.Item{})) {
		db.CreateTable(&storage.Item{})
	}

	resp, err := http.Get(url)
	if err != nil || (resp != nil && (resp.StatusCode > 299 || resp.StatusCode < 200)) {
		fmt.Println("loi1")
		return
	}

	// close body after reading content
	defer resp.Body.Close()

	body, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		panic(err1)
	}

	data := storage.ListItemsCraw{}
	_ = json.Unmarshal(body, &data)

	for i, itemcraw := range data.ListItemsCraw {
		item := storage.ConvertItemCrawToItem(itemcraw)

		//save to database
		db.Create(&item)

		//save to memory
		m.M[int64(i)] = item

		//save to file
		handle.WriteToFile("data/DataItem.txt", storage.ConvertItemToString(item))
	}

}
