package storage

type Data struct {
	Id   uint    `gorm:"primaryKey"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}
