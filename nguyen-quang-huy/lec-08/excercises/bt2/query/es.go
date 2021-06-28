package query

import (
	"context"
	"encoding/csv"
	"fmt"
	"hi/storage"
	"os"
	"strconv"

	"github.com/olivere/elastic/v7"
)

func CreateDataEs(path string) {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}

	info, code, err := client.Ping("http://localhost:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	for i, line := range lines {
		number, err := strconv.Atoi(line[0])
		// don't ignore error
		if err != nil {
			fmt.Println("Error when converting to int: ", err)
			continue
		}
		data := storage.Comment{
			Number:  number,
			Title:   line[1],
			Content: line[2],
		}

		if i%1000 == 0 {
			fmt.Printf("Added row %vth\n", i)
		}

		put1, err := client.Index().
			Index("train").
			Id(strconv.Itoa(i)).
			BodyJson(data).
			Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}

		fmt.Printf("Indexedline %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	}

	fmt.Println("Done...!")

}

func QueryEs(field string, val string) {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}
	termQuery := elastic.NewTermQuery(field, val)
	searchResult, err := client.Search().
		Index("train").   // search in index "train"
		Query(termQuery). // specify the query
		Do(ctx)           // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query with term {%v: %v} took %d milliseconds\n", field, val, searchResult.TookInMillis)
}
