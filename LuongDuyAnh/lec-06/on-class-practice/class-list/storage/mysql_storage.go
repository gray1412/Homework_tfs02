package storage

type Data struct {
	Id   uint   `gorm:"primary_key"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}
