package storage

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Kết nối với database
func DBConnection() (*gorm.DB, error) {
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
