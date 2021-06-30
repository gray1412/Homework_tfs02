package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"book/book"
)

func main() {
	url := "http://localhost:9200"
	esclient, _ := book.NewESClient(url)
	bm := book.NewBookManager(esclient)
	start := time.Now()
	bm.SearchBooks("")
	end := time.Since(start)
	fmt.Print(end)
	// ReadData()
}

func ReadData() {
	url := "http://localhost:9200"
	esclient, _ := book.NewESClient(url)
	bm := book.NewBookManager(esclient)
	insertedBookCount := 0
	csvFile, err := os.Open("train.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for i, line := range csvLines {
		// Need to check this first to ensure that we have right data format
		// these called validation
		if len(line) < 3 {
			fmt.Println("Wrong row format: ", line, " at line: ", i)
			continue
		}
		// don't ignore error
		if err != nil {
			fmt.Println("Error when converting to int: ", err)
			continue
		}
		review := book.Book{
			ID:    string(insertedBookCount),
			Rate:  line[0],
			Title: line[1],
			Body:  line[2],
		}
		err = bm.AddBook(&review)
		if err != nil {
			fmt.Println("Cannot insert book: ", review, err)
			continue
		}
		// just for logging - we can see the application is running - and printing something
		if i%100 == 0 {
			fmt.Printf("Inserted %v row(s)\n", i)
		}
		insertedBookCount++
	}
}
