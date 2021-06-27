package document

import (
	"encoding/json"
	"fmt"
)

type Document struct {
	ID    string `json:"id"`
	Rate  string `json:"rate"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// String returns object's string representation
func (book *Document) String() string {
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
