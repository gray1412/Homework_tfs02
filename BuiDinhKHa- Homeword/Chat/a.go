package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Record struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func main() {
	var allRecords []Record

	input := []byte(`[{
      "author": "Nirvana",
      "title":  "Smells like teen spirit"
    }, {
      "author": "The Beatles",
      "title":  "Help"
    }]`)

	var tmpRecords []Record
	err := json.Unmarshal(input, &tmpRecords)
	if err != nil {
		log.Fatal(err)
	}

	allRecords = append(allRecords, tmpRecords...)

	fmt.Println("RECORDS:", allRecords)
}
