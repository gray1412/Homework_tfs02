package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"lec06-hw/db"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.Class)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//Check if id is out of range
	if id >= 0 && id >= len(db.Class) {
		fmt.Fprintf(w, "No element at id %v", id)
		return
	}

	json.NewEncoder(w).Encode(db.Class[id])
}

func AddOne(w http.ResponseWriter, r *http.Request) {

	newStudent := db.Student{}

	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		fmt.Fprintf(w, "error when parsing body %v", err)
		return
	}
	db.Class = append(db.Class, newStudent)

	fmt.Fprintf(w, "Added %v, Current: %v", newStudent, db.Class)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	newStudent := db.Student{}

	//Check if id is out of range
	if id >= 0 && id >= len(db.Class) {
		fmt.Fprintf(w, "No element at id %v", id)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		fmt.Fprintf(w, "error when parsing body %v", err)
		return
	}

	db.Class[id] = newStudent

	fmt.Fprintf(w, "student at id %v changed to %v", id, db.Class[id])
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//Check if id is out of range
	if id >= 0 && id >= len(db.Class) {
		fmt.Fprintf(w, "No element at id %v", id)
		return
	}
}
