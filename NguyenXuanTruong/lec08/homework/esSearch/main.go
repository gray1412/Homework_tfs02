package main

import (
	"context"
	doc "document/document"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"time"

	elastic "github.com/olivere/elastic/v7"
)

func main() {
	//time to run func
	defer TimeTrack(time.Now())
	url := "http://localhost:9200"
	esclient, _ := doc.NewESClient(url)
	// ReadFile(esclient)

	// search
	bm := doc.NewBookManager(esclient)
	// case 1: found
	bm.SearchBooks("beautiful")
	// fmt.Println("Found books: ", resultSearchBooksSuccess)
	// // case 2: not found
	// resultSearchBooksFailed := bm.SearchBooks("8")
	// fmt.Println("Found books: ", resultSearchBooksFailed)
	// // delete
	// bm.DeleteBook(&doc.Book{ID: "5"})
	// // search again, but not found
	// resultSearchBooksSuccess = bm.SearchBooks("5")
	// fmt.Println("Found books: ", resultSearchBooksSuccess)

}

const (
	FILE       = "./train.csv"
	TOTAL_ROWS = 116100
)

func ReadFile(esclient *doc.ESClient) {
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
			req := elastic.NewBulkIndexRequest().Index("documents").Type("doc").Id(document.ID).Doc(document)
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

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)

	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	log.Println(fmt.Sprintf("%s took %s", name, elapsed))
}
