package main

import (
	"fmt"
	doc "tfs-02/lec-08/exercises/es/document"
)

func main() {
	url := "http://localhost:9200"
	esclient, _ := doc.NewESClient(url)
	readFile(esclient)

	// search
	// case 1: found
	resultSearchBooksSuccess := doc.SearchBooks("Stuning even for the non-gamer")
	fmt.Println("Found books: ", resultSearchBooksSuccess)
	// case 2: not found
	resultSearchBooksFailed := doc.SearchBooks("The best soundtrack")
	fmt.Println("Found books: ", resultSearchBooksFailed)
	// delete
	doc.DeleteBook(&doc.Document{ID: "1"})
	// search again, but not found
	resultSearchBooksSuccess = doc.SearchBooks("Stuning even for the non-gamer")
	fmt.Println("Found books: ", resultSearchBooksSuccess)
}
