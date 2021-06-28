package storage

import (
	"encoding/json"
	"fmt"
)

// Book defines a book's attributes and basic operation and interaction with Elasticsearch
type Book struct {
	Id     int `gorm:"primary_key"`
	Title string `json:"title"`
	Type  int    `json:"type"`
	Body  string `json:"body"`
}

// String returns object's string representation
func (book *Book) String() string {
	if book == nil {
		return ""
	}

	// Convert struct to json
	b, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Cannot convert to json: ", err)
		return ""
	}
	return string(b)
}
