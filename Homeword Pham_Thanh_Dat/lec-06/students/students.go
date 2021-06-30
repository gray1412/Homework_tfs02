package students

import (
	Api "Api/Storage"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type DataFile struct {
	Name string
	File string
}

func PathStudent() {
	r := mux.NewRouter().StrictSlash(true)
	get := r.Methods(http.MethodGet).Subrouter()
	get.Path("/students").HandlerFunc(viewStudent)
	get.Use(contentTypeJson)
	post := r.Methods(http.MethodPost).Subrouter()
	post.Path("/students").HandlerFunc(addStudent)
	r.Methods(http.MethodPost).Path("/upload").HandlerFunc(uploadFile)
	post.Use(contentTypeJson)
	put := r.Methods(http.MethodPut).Subrouter()
	put.Path("/students").HandlerFunc(updateStudent)
	put.Use(contentTypeJson)
	delete := r.Methods(http.MethodDelete).Subrouter()
	delete.Path("/students").HandlerFunc(deleteStudent)
	delete.Use(contentTypeJson)
	http.Handle("/", r)
	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
	}).Handler(r)
	// log.Fatal(http.ListenAndServe(":8082", handler))
	http.ListenAndServe(":8000", handler)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	var dataFile DataFile
	err := json.NewDecoder(r.Body).Decode(&dataFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	f, err := os.Create(dataFile.Name)
	if err != nil {
		fmt.Print("Aaaaaaaa")
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(dataFile.File)
	if err2 != nil {
		log.Fatal(err2)
	}
}
func viewStudent(w http.ResponseWriter, r *http.Request) {
	db := *Api.Connect()
	var student []Api.Students
	db.Find(&student)
	b, _ := json.Marshal(student)
	fmt.Fprint(w, string(b))

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
	var student []Api.Students
	db.Find(&student)
	b, _ := json.Marshal(student)
	fmt.Fprint(w, string(b))
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
	var student []Api.Students
	db.Find(&student)
	b, _ := json.Marshal(student)
	fmt.Fprint(w, string(b))
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
	var student []Api.Students
	db.Find(&student)
	b, _ := json.Marshal(student)
	fmt.Fprint(w, string(b))
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
