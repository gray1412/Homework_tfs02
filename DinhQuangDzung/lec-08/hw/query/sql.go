package query

import (
	"fmt"

	"lec08-hw/database"
	"lec08-hw/reader"

	"gorm.io/gorm"
)

func AddToSql() {
	s := []database.Review{}

	db := database.ConnectSQL()

	lines, err := reader.ReadCsv("./train.csv")
	if err != nil {
		panic(err)
	}

	db.Debug().Migrator().DropTable(&database.Review{})
	db.AutoMigrate(&database.Review{})

	// Loop through lines & turn into object
	for i, line := range lines {
		data := database.Review{
			Type:  line[0],
			Title: line[1],
			Body:  line[2],
		}
		if i == 100000 {
			break
		}
		s = append(s, data)
		fmt.Println("Added row", i)
	}
	db = db.Session(&gorm.Session{CreateBatchSize: 1000})
	db.Create(s)

	fmt.Println("Done...!")
}
