package handler

import (
	"encoding/json"
	"io"
	"lec06/database"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var db = database.ConnectDb()


type Tfs2 struct {
	Id     int `gorm:"primary_key"`
	Name   string
	Age    int
	Gender string
	Phone  int
	Class  string
	// Status bool
}

func CreateMember(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))
	gender := r.FormValue("gender")
	phone, _ := strconv.Atoi(r.FormValue("phone"))
	// status, _ := strconv.ParseBool(r.FormValue("status"))
	// log.WithFields(log.Fields{"description": description}).Info("Add new TodoItem. Saving to database.")
	member := &Tfs2{Name: name, Age: age, Gender: gender, Phone: phone, Class: "Tfs2"}
	db.Create(&member)
	result := db.Last(&member).Value
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetMembers(w http.ResponseWriter, r *http.Request) {
	var members []Tfs2
	Tfs2 := "Tfs2"
	Tfs2Members := db.Where("class = ?", Tfs2).Find(&members).Value
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Tfs2Members)
}

func GetMemberByID(Id int) bool {
	member := &Tfs2{}
	result := db.First(&member, Id)
	if result.Error != nil {
		log.Warn("Member is not in class")
		return false
	}
	return true
}

func UpdateMember(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter from mux
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// Test if the TodoItem exist in DB
	err := GetMemberByID(id)
	if err == false {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"updated": false, "error": "Record Not Found"}`)
	} else {
		name := r.FormValue("name")
		age, _ := strconv.Atoi(r.FormValue("age"))
		gender := r.FormValue("gender")
		phone, _ := strconv.Atoi(r.FormValue("phone"))
		member := &Tfs2{}
		log.WithFields(log.Fields{"Id": id, "Name": member.Name}).Info("Updating TodoItem")
		db.First(&member, id)
		member.Name = name
		member.Phone = phone
		member.Gender = gender
		member.Age = age
		db.Save(&member)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"updated": true}`)
	}
}

func DeleteMember(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter from mux
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Test if the TodoItem exist in DB
	err := GetMemberByID(id)
	if err == false {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": false, "error": "Record Not Found"}`)
	} else {
		log.WithFields(log.Fields{"Id": id}).Info("Deleting Member")
		member := &Tfs2{}
		db.First(&member, id)
		db.Delete(&member)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": true}`)
	}
}
