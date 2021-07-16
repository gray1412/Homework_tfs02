package main

import (
	"fmt"
	"hi/query"
	"hi/storage"
	"time"
	// "logger"
)

func main() {
	pathFile := "train.csv"

	//1.Đọc file csv và luu data vao mysql
	query.ReadCSV(pathFile)
	fmt.Println("done!")

	// query.CreateDataEs(pathFile)
	// var s []storage.Comment
	// //search sql
	// fmt.Println("sql start time: ", time.Now())
	// db := query.ConnectDB()
	// db.Debug().Where("Body LIKE ?", "%game music! I have played the game Chrono Cross but out of all of the games%").Find(&s)
	// fmt.Println("sql end time: ", time.Now())

	// // search es
	// fmt.Println("sql start time: ", time.Now())
	// query.QueryEs("Body", "game music! I have played the game Chrono Cross but out of all of the games")
	// fmt.Println("sql end time: ", time.Now())

}
