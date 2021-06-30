package main

import (
	"fmt"
	"sync"

	"asslec8/book"
	"asslec8/handler"
)

// func print(s *sync.WaitGroup, c chan handler.Comment, quit chan bool) {
// 	for {
// 		select {
// 		case cc := <-c:
// 			fmt.Println(cc.Title)
// 			fmt.Println(cc.Body)
// 			fmt.Println(cc.Type)
// 		case <-quit:
// 			s.Done()
// 		}

// 	}
// }
// save data to esClient
func save(s *sync.WaitGroup, c chan book.Book, quit chan bool, bm *book.BookManager) {
	for {
		select {
		case cc := <-c:
			// fmt.Println(cc)
			err := bm.AddBook(&cc)
			if err != nil {
				fmt.Println(err)
			}

		case <-quit:
			s.Done()
		}

	}
}

func main() {
	link := "data/train1.csv"

	url := "http://localhost:9200"

	esclient, _ := book.NewESClient(url)
	bm := book.NewBookManager(esclient)

	s := sync.WaitGroup{}
	s.Add(2)
	c := make(chan book.Book)
	q := make(chan bool)
	go handler.ReadFile(&s, link, c, q)
	go save(&s, c, q, bm)
	s.Wait()

	// r := bm.SearchBooks("Stuning")
	// fmt.Println(r)
}
