package storage

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Kết nối với database
func dbConnection() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:1234@/dataFilm")
	if err != nil {
		log.Printf("Error %s when opening DBn", err)
		return nil, err
	}

	// db.SetMaxOpenConns(20)
	// db.SetMaxIdleConns(20)
	// db.SetConnMaxLifetime(time.Minute * 5)

	return db, nil
}
