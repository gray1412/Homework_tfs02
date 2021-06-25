package memory

type Data struct {
	Id    uint   `gorm:"primary_key"`
	Name  string `db:"name"`
	Age   int    `db:"age"`
	Phone string `db:"phone"`
}

type Memory_Storage_Data struct {
	M map[int64]Data
}

var Datas []Data
