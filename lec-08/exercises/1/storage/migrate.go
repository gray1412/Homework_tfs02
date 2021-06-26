package storage

func Migrate() {
	CreateTable()
}

//tao bang
func CreateTable() {
	db := ConnectToDatabase()
	db.AutoMigrate(&Role{}, &Person{}, &Room{}, &Course{}, &Lession{}, &Person{}, &Class{}, &Registration{}, &Slot{})
}
func GetAllPeople() {
	db := ConnectToDatabase()
	var Person Person
	// SELECT * FROM `pepple`
	db.Find(&Person)

}
