package main

import (
	"tfs-02/lec-08/exercises/storage"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectToDatabase() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:123456@/ocg_tfs_02?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database") // Kiểm tra kết nối tới databse
	}
	return
}

func main() {
	db := *ConnectToDatabase()
	defer db.Close()

	if (!db.HasTable(&storage.People{})) { //neu khong ton tai bang
		db.CreateTable(&storage.People{}) // thi tao bang
	} else {
		db.DropTable(&storage.People{})   // neu da ton tai bang thi xoa bang
		db.CreateTable(&storage.People{}) // sau do tao lai bang
	}

	if (!db.HasTable(&storage.PersonRole{})) {
		db.CreateTable(&storage.PersonRole{})
	} else {
		db.DropTable(&storage.PersonRole{})
		db.CreateTable(&storage.PersonRole{})
	}

	


}
