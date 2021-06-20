package database

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Class       string `json:"class"`
	Price       int    `json:"price"`
	Img         string `json:"img"`
	Description string `json:"description"`
}
