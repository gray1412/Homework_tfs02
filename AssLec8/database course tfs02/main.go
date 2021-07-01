package main

import (
	"tfs02/database"
)

func main() {
	database.Migrate()

	// db := database.ConnectToDatabase()
	// defer db.Close()

	// // a := database.People{
	// // 	ID:   3,
	// // 	Name: "ngoc",

	// // 	ClassId: 3,
	// // }
	// // db.Create(&a)

	// // var b database.People
	// // db.First(&b, 3)
	// // fmt.Println(b)
	// fmt.Println("done")
}
