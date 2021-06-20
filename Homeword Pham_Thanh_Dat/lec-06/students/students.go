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

	get := r.Methods(http.MethodGet).Subrouter()
	get.Path("/students").HandlerFunc(viewStudent)
	get.Use(contentTextPlan)
	post := r.Methods(http.MethodPost).Subrouter()
	post.Path("/students").HandlerFunc(addStudent)
	post.Use(contentTypeJson)
	put := r.Methods(http.MethodPut).Subrouter()
	put.Path("/students").HandlerFunc(updateStudent)
	put.Use(contentTypeJson)
	delete := r.Methods(http.MethodDelete).Subrouter()
	delete.Path("/students").HandlerFunc(deleteStudent)
	delete.Use(contentTypeJson)
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

// func viewPerson(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Student: %+v", Api.Students)
// }
func viewStudent(w http.ResponseWriter, r *http.Request) {
	db := Api.Connect()
	student := &Api.Students{}
	db.Find(&student)
	fmt.Printf("%+v", student)
}
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
	db := Api.Connect()
	var Students = Api.Students{Id: s.Id}
	db.Delete(&Students)
}
func contentTypeJson(next http.Handler) http.Handler {
	const contentType = "application/json"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		if reqContentType != contentType {
			fmt.Fprintf(w, "Only allow request with content type %v", contentType)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func contentTextPlan(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
