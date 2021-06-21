package storage

import "github.com/jinzhu/gorm"

type Storage interface {
	ConnectDatabase() (db *gorm.DB)
	CreateTableProduct()
}
