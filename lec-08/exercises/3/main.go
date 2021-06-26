package main

import (
	"fmt"

	"./book"
)

func main() {
	url := "http://localhost:9200"
	esclient, _ := book.NewESClient(url)
	bm := book.NewBookManager(esclient)
	// create new book
	books := createSomeBooks()
	// insert to es
	var err error
	insertedBookCount := 0
	for _, b := range books {
		err = bm.AddBook(b)
		if err != nil {
			fmt.Println("Cannot insert book: ", b, err)
			continue
		}
		insertedBookCount++
	}
	fmt.Printf("Inserted: %v books \n", insertedBookCount)

	// search
	// case 1: found
	resultSearchBooksSuccess := bm.SearchBooks("One Stop")
	fmt.Println("Found books: ", resultSearchBooksSuccess)
	// case 2: not found
	resultSearchBooksFailed := bm.SearchBooks("opencommerce group")
	fmt.Println("Found books: ", resultSearchBooksFailed)
	// delete
	bm.DeleteBook(&book.Book{ID: "1"})
	// search again, but not found
	resultSearchBooksSuccess = bm.SearchBooks("One Stop")
	fmt.Println("Found books: ", resultSearchBooksSuccess)
}

func createSomeBooks() []*book.Book {
	return []*book.Book{}
}
