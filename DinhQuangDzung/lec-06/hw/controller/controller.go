package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"lec06-hw/database"

	"github.com/gorilla/mux"
)

var db = database.ConnectDb()

func init() {
	var students = []database.Student{}
	db.CreateTable(&students)
	newStudent := database.Student{
		Id:   1,
		Name: "test",
		Age:  5,
	}
	db.Create(&newStudent)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var students = []database.Student{}
	results := db.Find(&students)
	json.NewEncoder(w).Encode(results.Value)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	student := database.Student{}
	result := db.First(&student, id)

	if result.Error != nil {
		fmt.Fprintf(w, "No entry at id %v", id)
		return
	}

	json.NewEncoder(w).Encode(result.Value)
}

func AddOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newStudent := database.Student{}

	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		fmt.Fprintf(w, "error when parsing body %v", err)
		return
	}

	if result := db.Create(&newStudent); result.Error != nil {
		fmt.Fprintf(w, "Couldn't add %v. Error: %v", newStudent, result.Error)
		return
	}

	fmt.Fprintf(w, "Added %v", newStudent)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	student := database.Student{}
	result := db.First(&student, id)
	if result.Error != nil {
		fmt.Fprintf(w, "No entry at id %v", id)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Fprintf(w, "error when parsing body %v", err)
		return
	}

	db.Save(&student)

	fmt.Fprintf(w, "Updated id %v to %v", id, student)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	student := database.Student{}
	result := db.Delete(&student, id)
	if result.Error != nil {
		fmt.Fprintf(w, "No entry at id %v", id)
		return
	}
	fmt.Fprintf(w, "Deleted id %v", id)
}
