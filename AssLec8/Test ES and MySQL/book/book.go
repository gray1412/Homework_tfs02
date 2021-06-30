package book

import (
	"encoding/json"
	"fmt"
)

// Book defines a book's attributes and basic operation and interaction with Elasticsearch
type Book struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
	Type  int    `json:"type"`
	Body  string `json:"body" gorm:"type:text"`
}

// String returns object's string representation
func (book *Book) String() string {
	if book == nil {
		return ""
	}
	b, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Cannot convert to json: ", err)
		return ""
	}
	return string(b)
}
