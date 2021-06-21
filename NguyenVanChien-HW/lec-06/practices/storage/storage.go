package storage

type Student struct {
	Id   int    `json:"id"`
	Name string `json: "name"`
	Age  int    `json: "age"`
}

var Students = map[int]Student{}
