package storage

type flimDb struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Rate string
	Link string
}
type Product struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

type Products struct {
	Products []Product `json:"products"`
}
