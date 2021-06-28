package handler

import (
	"encoding/csv"
	"fmt"
	"ftxMysql/storage"
	"os"
	"strconv"
)

// ReadFile train.csv and save to SQLDatabase and EsClient
func ReadFile(link string) {
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
	}

	// Connect to Database
	db := storage.ConnectDb()
	defer db.Close()
	// defer db.Close()

	db.Debug().DropTableIfExists(&storage.Book{})
	db.Debug().AutoMigrate(&storage.Book{})
	if (!db.HasTable(&storage.Book{})) {
		db.CreateTable(&storage.Book{})
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

		newBook := storage.Book{
			Type:  data,
			Title: line[1],
			Body:  line[2],
		}
		db.Create(&newBook)

		if i%10000 == 0 {
			fmt.Println("Done!")
			// return
		}
	}
	// fmt.Printf("There are %v had been writen", count)
	fmt.Println("Finish to write data to MySQL database and ES database")

}
