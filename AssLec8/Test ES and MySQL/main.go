package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"

	"asslec8/book"
	"asslec8/database"
	"asslec8/handler"
)

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
func importToES(link, url string, bm *book.BookManager) {
	s := sync.WaitGroup{}
	s.Add(2)
	c := make(chan book.Book)
	q := make(chan bool)
	go handler.ReadFile(&s, link, c, q)
	go save(&s, c, q, bm)
	s.Wait()
}
func importToMYSQL(link string) {

	// Connect to Database
	db := database.ConnectToDatabase()
	defer db.Close()

	db.DropTableIfExists(&book.Book{})
	db.CreateTable(&book.Book{})

	file, err := os.Open(link)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File have been open!")
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	for i, line := range lines {
		// Check this first to ensure that we have right data format
		if len(line) < 3 {
			fmt.Println("Wrong row format", line, " at line: ", i)
			continue
		}

		data, err := strconv.Atoi(line[0])
		// don't ignore error
		if err != nil {
			fmt.Println("Error when converting to int: ", err)
			continue
		}

		newBook := book.Book{
			ID:    strconv.Itoa(i),
			Type:  data,
			Title: line[1],
			Body:  line[2],
		}

		db.Create(&newBook)
		if i%100 == 0 {
			fmt.Println(i)
		}
		// count = i
	}
	// fmt.Printf("There are %v had been writen", count)
	fmt.Println("Finish to write data to MySQL database and ES database")

}
func main() {
	link := "train.csv"

	// url := "http://localhost:9200"

	// esclient, _ := book.NewESClient(url)
	// bm := book.NewBookManager(esclient)

	//import to ES
	// importToES(link, url, bm)

	//import to Database
	importToMYSQL(link)

	//search ES
	// 	start := time.Now()

	// 	r := bm.SearchBooks("")

	// 	end := time.Since(start)
	// 	fmt.Println(end)
	// 	fmt.Println(len(r))
}
