package Storage

type Students struct {
	Id      uint   `gorm:"primaryKey"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   int    `json:"phone"`
	Age     int    `json:"age"`
}
