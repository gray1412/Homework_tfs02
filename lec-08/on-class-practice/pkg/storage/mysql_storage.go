package storage

import "fmt"

type Product struct {
	Id         uint   `gorm:"primary_key"`
	Name       string `db:"name"`
	Img_src    string `db:"img_src"`
	Created_at string `db:"created_at"`
}

func ConvertProductToString(p Product) (s string) {
	s = "ID: " + fmt.Sprintf("%d", p.Id) + ", NAME: " + p.Name + ", RATE: " + p.Img_src + "CREATED_AT: " + p.Created_at
	return
}
