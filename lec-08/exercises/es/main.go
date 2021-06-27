package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	doc "tfs-02/lec-08/exercises/es/document"

	elastic "github.com/olivere/elastic/v7"
)

func main() {
	url := "http://localhost:9200"
	esclient, _ := doc.NewESClient(url)
	readFile(esclient)

	// search
	bm := doc.NewBookManager(esclient)
	// case 1: found
	resultSearchBooksSuccess := bm.SearchBooks("Stuning even for the non-gamer")
	fmt.Println("Found books: ", resultSearchBooksSuccess)
	// case 2: not found
	resultSearchBooksFailed := bm.SearchBooks("The best soundtrack")
	fmt.Println("Found books: ", resultSearchBooksFailed)
	// delete
	bm.DeleteBook(&doc.Book{ID: "1"})
	// search again, but not found
	resultSearchBooksSuccess = bm.SearchBooks("Stuning even for the non-gamer")
	fmt.Println("Found books: ", resultSearchBooksSuccess)

}

const (
	FILE       = "../train.csv"
	TOTAL_ROWS = 50000
)

func readFile(esclient *doc.ESClient) {
	ctx := context.Background()
	file, err := os.Open(FILE)
	printError(err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true // bo dau ""

	rows, err := reader.ReadAll()
	printError(err)

	bulkRequest := esclient.Bulk()
	for n, col := range rows {
		n++
		// id := uuid.NewV4().String()
		if n <= TOTAL_ROWS {
			document := doc.Book{
				ID:    strconv.Itoa(n),
				Rate:  col[0],
				Title: col[1],
				Body:  col[2],
			}
			req := elastic.NewBulkIndexRequest().Index("documents").Type("_doc").Id(document.ID).Doc(document)
			bulkRequest = bulkRequest.Add(req)
			if n%10000 == 0 {
				fmt.Printf("%v:\n", n)
			}
		}
	}
	bulkResponse, err := bulkRequest.Do(ctx)
	printError(err)

	indexed := bulkResponse.Indexed()
	if len(indexed) != 1 {
		fmt.Printf("\n Indexed documents: %v  \n", len(indexed))
	}

}

func printError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
