package storage

import (
	"context"
	"encoding/json"

	// "errors"
	"fmt"
	"math/rand"
	"time"

	elastic "github.com/olivere/elastic/v7"
)

const (
	indexName = "book"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type BookManager struct {
	esClient *ESClient
}

func NewBookManager(es *ESClient) *BookManager {
	return &BookManager{esClient: es}
}

func (bm *BookManager) SearchBooks(title string) []*Book {
	ctx := context.Background()

	if bm.esClient == nil {
		fmt.Println("Nil es client")
		return nil
	}

	// build query to search for title
	query := elastic.NewSearchSource()
	query.Query(elastic.NewMatchQuery("title", title))

	// get search's service
	searchService := bm.esClient.Search().Index(indexName).SearchSource(query)

	// perform search query
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("Cannot perform search with ES", err)
		return nil
	}

	// get result

	var books []*Book

	for _, hit := range searchResult.Hits.Hits {
		var book Book
		err := json.Unmarshal(hit.Source, &book)
		if err != nil {
			fmt.Println("Get data error: ", err)
			continue
		}
		fmt.Println(&book)
		books = append(books, &book)
	}

	return books
}

