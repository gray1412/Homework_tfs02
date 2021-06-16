package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Thêm dữ liệu vào database
func Insert2DatabaseFilm(db *gorm.DB, ch chan Film) {
	for {
		select {
		case <-ch:
			s := <-ch
			db.Create(&Film{
				Name:  s.Name,
				Image: s.Image,
				Rate:  s.Rate,
				Url:   s.Url})
		}
	}
}
