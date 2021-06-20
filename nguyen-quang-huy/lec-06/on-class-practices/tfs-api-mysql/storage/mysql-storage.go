package storage


type Person struct {
	ID   int    `db: "id"`
	Name string `db: "name"`
	Age int `db: "age"`
}