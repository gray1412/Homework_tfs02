package main

import (
	"lec08-hw/database"
	"lec08-hw/query"
)

var db = database.ConnectSQL()

var s []database.Review

func main() {
	// db.AutoMigrate(&database.Role{}, &database.Person{}, &database.Room{}, &database.Course{}, &database.Lession{}, &database.Person{}, &database.Class{}, &database.Registration{}, &database.Slot{})

	// query.AddToSql()
	// query.AddToEs()

	db.Debug().Where("Body LIKE ?", "%computer%").Find(&s)

	query.QueryEs("Body", "computer")
}
