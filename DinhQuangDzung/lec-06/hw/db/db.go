package db

type Student struct {
	Name string `json:"name" db:"name"`
	Age  int    `json:"age"  db:"age"`
}

var Class []Student

func init() {
	student := Student{
		Name: "test",
		Age:  0,
	}
	Class = append(Class, student)
}
