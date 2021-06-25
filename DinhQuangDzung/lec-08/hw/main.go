package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"lec08-hw/database"

	"gorm.io/gorm"
)

type CsvLine struct {
	Type  string
	Title string
	Body  string
}

var s []CsvLine

func main() {

	lines, err := ReadCsv("./train.csv")
	if err != nil {
		panic(err)
	}

	db := database.ConnectDB()

	db.Debug().Migrator().DropTable(&CsvLine{})
	db.AutoMigrate(&CsvLine{})

	// Loop through lines & turn into object
	for i, line := range lines {
		data := CsvLine{
			Type:  line[0],
			Title: line[1],
			Body:  line[2],
		}
		if i%1000 == 0 {
			fmt.Printf("Added row %vth\n", i)
		}
		s = append(s, data)
	}
	db = db.Session(&gorm.Session{CreateBatchSize: 1000})
	db.Create(s)

	fmt.Println("Done...!")
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
