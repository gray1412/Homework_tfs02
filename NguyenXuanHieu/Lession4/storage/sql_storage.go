package storage

import (
	"fmt"
)

type Item struct {
	Title string `db:"title"`
	Price int64  `db:"price"`
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

type Movie struct {
	Name   string `json:"name"`
	Year   string `json:"year"`
	Rating string `json:"rating"`
}

func ConvertItemCrawToItem(i ItemCraw) (r Item) {
	r.Title = i.Title
	r.Price = i.Price
	return
}

func ConvertItemToString(i Item) (s string) {
	s = "TITLE: " + i.Title + ", PRICE:" + fmt.Sprintf("%d", i.Price)
	return
}
func ConvertFilmToString(f Movie) (s string) {
	s = "Name: " + f.Name + ", YEAR: " + f.Year + ", RATING: " + f.Rating
	return
}
