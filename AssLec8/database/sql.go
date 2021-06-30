package query

import (
	"hi/storage"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
var db = ConnectDB()
func init(){
	db.AutoMigrate(&storage.Comment{})

}
func ReadCSV(path string) {
	
	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	count := 0
	csvs := csv.NewReader(csvFile)
	data := []storage.Comment{}
	totalcount := 0
	line, err := csvs.Read()
	for {
		if err != nil {
			//luu data
			db.Create(data)
			fmt.Println("end: ", time.Now())
			fmt.Println("Tong: ", totalcount)

			break
		}
		if len(line) < 3 {
			continue
		}
		number, err := strconv.Atoi(line[0])
		// don't ignore error
		if err != nil {
			fmt.Println("Error when converting to int: ", err)
			continue
		}
		data = append(data, storage.Comment{
			Number:  number,
			Title:   line[1],
			Content: line[2],
		})
		count++
		totalcount++
		if count == 1000 {
			db.Create(data)
			count = 0
			data = []storage.Comment{}

		}
		if totalcount%100000 == 0 {

			fmt.Println(totalcount)
			fmt.Println(time.Now())
		}
		if totalcount == 3600000 {
			fmt.Println("Day")
			break
		}
		line, err = csvs.Read()
	}

}
func ConnectDB() *gorm.DB {
	// dsn := "root:@tcp(127.0.0.1:3306)/testdb"

	dsn := "root:password@tcp(127.0.0.1:3306)/testdb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000,
		Logger:          logger.Default.LogMode(logger.Silent),
	})
	db, err := gorm.Open("mysql", "root:ngochd246@/test_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conected db")
	return db
}

