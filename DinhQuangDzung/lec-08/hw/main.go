package main

import (
	"lec08-hw/database"
	"lec08-hw/query"
)

var db = database.ConnectSQL()

var s []database.CsvLine

func main() {
	// query.AddToSql()
	// query.AddToEs()

	// db.Debug().Find(&s)

	db.Debug().Where("Title LIKE ?", "%computer%").Find(&s)

	query.QueryEs("Title", "computer")
}
