package database

import (
	"time"
)

// func init() {
// 	db := ConnectToDatabase()
// 	db.AutoMigrate(&People{}, &Class{}, &Course{}, &Lesson{}, &Subject{}, &Room{})
// 	db.Close()
// }
func Migrate() {
	db := ConnectToDatabase()
	db.AutoMigrate(&Room{}, &Course{}, &Subject{}, &Class{}, &People{}, &Lesson{})

}

type People struct {
	ID         int      `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	Name       string   `json:"name" gorm:"type:varchar(51)"`
	Phone      string   `json:"phone" gorm:"type:varchar(11)"`
	Role       string   `json:"role" gorm:"type:varchar(7)"`
	ClassId    int      `json:"class_id" gorm:"default:null; type:mediumint"`
	ListLesson []Lesson `gorm:"foreignKey:TeacherId;ASSOCIATION_FOREIGNKEY:ID"`
}
type Class struct {
	ID            int      `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	Name          string   `json:"name" gorm:"type:varchar(50)"`
	NumberStudent int      `json:"number_student" gorm:"type:tinyint"`
	CourseId      int      `json:"course_id" gorm:"default:null; type:mediumint"`
	ListStudent   []People `gorm:"foreignKey:ClassId;ASSOCIATION_FOREIGNKEY:ID"`
	ListLesson    []Lesson `gorm:"foreignKey:ClassId;ASSOCIATION_FOREIGNKEY:ID"`
}
type Course struct {
	ID          int       `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	Name        string    `json:"name" gorm:"type:varchar(50)"`
	Description string    `json:"description" gorm:"type:mediumtext"`
	ListClass   []Class   `gorm:"foreignKey:CourseId;ASSOCIATION_FOREIGNKEY:ID"`
	ListSubject []Subject `gorm:"foreignKey:CourseId;ASSOCIATION_FOREIGNKEY:ID"`
}
type Lesson struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	Name      string    `json:"name" gorm:"type:varchar(50)"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	RoomId    int       `json:"room" gorm:"default:null; type:mediumint"`
	SubjectId int       `json:"subject_id" gorm:"default:null; type:mediumint"`
	TeacherId int       `json:"teacher_id" gorm:"default:null; type:mediumint"`
	ClassId   int       `json:"class_id" gorm:"default:null; type:mediumint"`
}
type Subject struct {
	ID          int      `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	Name        string   `json:"name" gorm:"type:varchar(50)"`
	Description string   `json:"description" gorm:"type:mediumtext"`
	CourseId    int      `json:"course_id" gorm:"default:null; type:mediumint"`
	ListLesson  []Lesson `gorm:"foreignKey:SubjectId;ASSOCIATION_FOREIGNKEY:ID"`
}
type Room struct {
	ID         int      `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	Name       string   `json:"name" gorm:"type:varchar(50)"`
	Address    string   `json:"address" gorm:"type:varchar(255)"`
	ListLesson []Lesson `gorm:"foreignKey:RoomId;ASSOCIATION_FOREIGNKEY:ID"`
}
