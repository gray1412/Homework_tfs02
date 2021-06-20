package students

import (
	Api "Api/Storage"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func PathStudent() {
	r := mux.NewRouter().StrictSlash(true)
	r.Methods(http.MethodPost).Path("/students").HandlerFunc(addStudent)
	r.Methods(http.MethodPut).Path("/students").HandlerFunc(updateStudent)
	r.Methods(http.MethodDelete).Path("/students").HandlerFunc(deleteStudent)
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

// func viewPerson(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Student: %+v", Api.Students)
// }
func addStudent(w http.ResponseWriter, r *http.Request) {
	var s Api.Students
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := Api.Connect()
	var Studens = Api.Students{Name: s.Name, Address: s.Address, Phone: s.Phone, Age: s.Age}
	db.Create(&Studens)
}
func updateStudent(w http.ResponseWriter, r *http.Request) {
	var s Api.Students
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Print(s)
	db := Api.Connect()
	// var Studens = Api.Students{Id:1}
	// db.Model(&Studens).Update(s.Name, s.Address, s.Phone, s.Age)
	var Students = Api.Students{Id: s.Id}
	db.First(&Students)
	Students.Name = s.Name
	Students.Age = s.Age
	Students.Address = s.Address
	Students.Phone = s.Phone
	db.Save(&Students)
}
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	var s Api.Students
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Print(s)
	db := Api.Connect()
	var Students = Api.Students{Id: s.Id}
	db.Delete(&Students)
}
