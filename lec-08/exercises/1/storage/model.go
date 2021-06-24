package storage

type Person struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `json: "name"`
	Age    int    `json: "age"`
	TypeID []Type `gorm:"foreignKey:TypeID;associationForeignKey:ID"`
}
type Type struct {
	ID          uint   `gorm:"primaryKey"`
	Type        string `json: "type"`
	Email       string `json: "email"`
	Description string `json: "description"`
}
type Registration struct {
	ID       uint   `gorm:"primaryKey"`
	PersonID []Person `gorm:"foreignKey:PersonID;associationForeignKey:ID"`
}
type Class struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json: "name"`
	LessonID []Lesson `gorm:"foreignKey:LessonID;associationForeignKey:ID"`
}
type Lesson struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `json: "name"`
	StartTime string  `json: "starttime"`
	EndTime   string  `json: "endtime"`
	SubjectID []Subject `gorm:"foreignKey:SubjectID;associationForeignKey:ID"`
	RoomID    []Room    `gorm:"foreignKey:RoomID;associationForeignKey:ID"`
	PersonID  []Person  `gorm:"foreignKey:PersonID;associationForeignKey:ID"`
}
type Subject struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json: "name"`
	Description string `json: "description"`
}
type Room struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json: "name"`
	PlaceID     string `json: "place"`
	Description string `json: "description"`
}
