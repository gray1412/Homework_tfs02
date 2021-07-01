package handler

import (
	"encoding/csv"
	"ex8/storage"
	"fmt"
	"os"
	"strconv"
)

// ReadFile train.csv and save to SQLDatabase and EsClient
func ReadFile(link string) [][]string {
	// open file
	file, err := os.Open(link)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("File have been open!")
	defer file.Close()

	// get all lines in file
	lines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return lines
}

func SaveToSQL(link string) {
	lines := ReadFile(link)

	db := storage.ConnectDatabase()
	defer db.Close()

	storage.CreateTableBook()

	for i, line := range lines {
		if len(line) < 3 {
			fmt.Println("Wrong row format", line, " at line: ", i)
			continue
		}

		data, err := strconv.Atoi(line[0])

		if err != nil {
			fmt.Println("Error when converting to int: ", err)
			continue
		}

		newBook := storage.Book{
			ID:    strconv.Itoa(i),
			Type:  data,
			Title: line[1],
			Body:  line[2],
		}

		db.Create(&newBook)
		if i > 10000 {
			fmt.Println("Have write 10000 data row to table books in SQL")
			break
		}
	}
}

func SaveToEsClient(link string, bm *storage.BookManager) {
	lines := ReadFile(link)
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

		newBook := storage.Book{
			Type:  data,
			Title: line[1],
			Body:  line[2],
		}

		bm.AddBook(&newBook)

		if i > 10000 {
			fmt.Println("Have write 10000 row to EsClient!")
			break
		}
	}
}
