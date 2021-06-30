package handler

import (
	"asslec8/book"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func ReadFile(s *sync.WaitGroup, link string, c chan book.Book, quit chan bool) {
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
	// Check this first to ensure that we have right data format
	for i, line := range lines {
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

		newbook := book.Book{
			Type:  data,
			Title: line[1],
			Body:  line[2],
		}

		c <- newbook

		if i%1000 == 0 {
			fmt.Println(i)
		}
	}
	s.Done()
	quit <- true

}
