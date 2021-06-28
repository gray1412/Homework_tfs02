package storage

import (
	"encoding/json"
	"fmt"
)

type Comment struct {
	Id      int
	Number  int
	Title   string
	Content string
}

func (comment *Comment) String() string {
	if comment == nil {
		return ""
	}
	b, err := json.Marshal(comment)
	if err != nil {
		fmt.Println("Cannot convert to json: ", err)
		return ""
	}
	return string(b)
}
