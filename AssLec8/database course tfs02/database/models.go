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
	db.AutoMigrate(&People{}, &Class{}, &Course{}, &Lesson{}, &Subject{}, &Room{})
	db.Close()
}

type People struct {
	ID         int      `json:"id" gorm:"PRIMARY_KEY"`
	Name       string   `json:"name"`
	Phone      string   `json:"phone"`
	Role       string   `json:"role"`
	ClassId    int      `json:"class_id"`
	ListLesson []Lesson `gorm:"FOREIGNKEY:TeacherID;ASSOCIATION_FOREIGNKEY:ID"`
}
type Class struct {
	ID            int      `json:"id" gorm:"PRIMARY_KEY"`
	Name          string   `json:"name"`
	NumberStudent int      `json:"number_student"`
	CourseId      int      `json:"course_id"`
	ListStudent   []People `gorm:"FOREIGNKEY:ClassID;ASSOCIATION_FOREIGNKEY:ID"`
}
type Course struct {
	ID          int       `json:"id" gorm:"PRIMARY_KEY"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ListClass   []Class   `gorm:"FOREIGNKEY:CourseID;ASSOCIATION_FOREIGNKEY:ID"`
	ListSubject []Subject `gorm:"FOREIGNKEY:CourseID;ASSOCIATION_FOREIGNKEY:ID"`
}
type Lesson struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	RoomId    int       `json:"room"`
	SubjectId int       `json:"subject_id"`
	TeacherId int       `json:"teacher_id"`
	ClassId   int       `json:"class_id"`
}
type Subject struct {
	ID          int      `json:"id" gorm:"PRIMARY_KEY"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CourseId    int      `json:"course_id"`
	ListLesson  []Lesson `gorm:"FOREIGNKEY:SubjectID;ASSOCIATION_FOREIGNKEY:ID"`
}
type Room struct {
	ID         int      `json:"id" gorm:"PRIMARY_KEY"`
	Name       string   `json:"name"`
	Address    string   `json:"address"`
	ListLesson []Lesson `gorm:"FOREIGNKEY:RoomID;ASSOCIATION_FOREIGNKEY:ID"`
}
