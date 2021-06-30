package data

import (
	"encoding/csv"
	"fmt"
	storage "fultextsearch/storage"
	"os"
	"strconv"
)

func ReadData() {
	csvFile, err := os.Open("data/train.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	db := storage.Connect()
	for i, line := range csvLines {
		// Need to check this first to ensure that we have right data format
		// these called validation
		if len(line) < 3 {
			fmt.Println("Wrong row format: ", line, " at line: ", i)
			continue
		}
		data, err := strconv.Atoi(line[0])
		// don't ignore error
		if err != nil {
			fmt.Println("Error when converting to int: ", err)
			continue
		}
		review := storage.Review{
			Rate:  data,
			Title: line[1],
			Body:  line[2],
		}
		db.Create(&review)
		// just for logging - we can see the application is running - and printing something
		if i%100 == 0 {
			fmt.Printf("Inserted %v row(s)\n", i)
		}
	}
}
