package storage

import (
	"course/database"
	"time")


func init() {
	db := database.ConnectDB()
	db.AutoMigrate(&Role{}, &User{}, &Room{}, &Course{}, &Class{}, &UserClass{}, &Lesson{}, &Slot{})
}

type User struct {
	Id            int         `json:"id"`
	FirstName     string      `json:"firstname" gorm:"type:varchar(20)"`
	LastName      string      `json:"lastname" gorm:"type:varchar(20)"`
	Phone         string      `json:"phone" gorm:"type:varchar(11)"`
	Email         string      `json:"email" gorm:"type:varchar(50)"`
	Address       string      `json:"address" gorm:"type:varchar(50)"`
	Username      string      `json:"username" gorm:"type:varchar(20)"`
	Password      string      `json:"password" gorm:"type:varchar(20)"`
	RoleId        int         `json:"roleid"`
	UserClassList []UserClass `json:"userclasslist" gorm:"foreignKey:UserId"`
}
type Role struct {
	Id       int    `json:"id"`
	Name     string `json:"lastname" gorm:"type:varchar(20)"`
	UserList []User `json:"userlist" gorm:"foreignKey:RoleId"`
}
type UserClass struct {
	Id      int `json:"id"`
	UserId  int `json:"userid"`
	ClassId int `json:"classid"`
}
type Class struct {
	Id            int         `json:"id"`
	Name          string      `json:"name" gorm:"type:varchar(20)"`
	Description   string      `json:"description"`
	CourseId      int         `json: courseid`
	UserClassList []UserClass `json:"userclasslist" gorm:"foreignKey:ClassId"`
	SlotList      []Slot      `json:"slotlist" gorm:"foreignKey:ClassId"`
}
type Course struct {
	Id          int      `json:"id"`
	Name        string   `json:"name" gorm:"type:varchar(20)"`
	Description string   `json:"description"`
	ClassList   []Class  `json:"classlist" gorm:"foreignKey:CourseId"`
	LessonList  []Lesson `json:"lessonlist" gorm:"foreignKey:CourseId"`
}
type Lesson struct {
	Id          int    `json:"id"`
	Name        string `json:"name" gorm:"type:varchar(20)"`
	Description string `json:"description"`
	CourseId    int    `json:courseId`
	SlotList    []Slot `json:"lessonlist" gorm:"foreignKey:LessonId"`
}
type Slot struct {
	Id        int       `json:"id"`
	StartTime time.Time `json:"starttime"`
	EndTime   time.Time `json:"endtime"`
	ClassId   int       `json:"classID"`
	LessonId  int       `json:"lessonid"`
	RoomId    int       `json:"roomid"`
}
type Room struct {
	Id       int    `json:"id"`
	Address  string `json:"address" gorm:"type:varchar(50)"`
	SlotList []Slot `json:"slotlist" gorm:"foreignKey:RoomId"`
}
