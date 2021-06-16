package storage

import (
	"fmt"
)

type Item struct {
	ID        int64  `db:"id"`
	Title     string `db:"title"`
	Price     int64  `db:"price"`
	ImgSrc    string `db:"imgSrc"`
	ImgWidth  int64  `db:"imgWidth"`
	ImgHeight int64  `db:"imgHeight"`
}
type ImageCraw struct {
	Src    string `json:"src"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}
type ItemCraw struct {
	ID    int64       `json:"id"`
	Title string      `json:"title"`
	Img   []ImageCraw `json:"images"`
	Price int64       `json:"price"`
}
type ListItemsCraw struct {
	ListItemsCraw []ItemCraw `json:"products"`
}
type Film struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Rate string `db:"rate"`
}

func ConvertItemCrawToItem(i ItemCraw) (r Item) {
	r.ID = i.ID
	r.Title = i.Title
	r.Price = i.Price
	r.ImgSrc = i.Img[0].Src
	r.ImgWidth = i.Img[0].Width
	r.ImgHeight = i.Img[0].Height
	return
}

func ConvertItemToString(i Item) (s string) {
	s = "ID: " + fmt.Sprintf("%d", i.ID) + ", TITLE: " + i.Title + ", PRICE:" + fmt.Sprintf("%d", i.Price) + ", IMG: " + i.ImgSrc
	return
}
func ConvertFilmToString(f Film) (s string) {
	s = "ID: " + fmt.Sprintf("%d", f.ID) + ", NAME: " + f.Name + ", RATE: " + f.Rate
	return
}
