package handler

import (
	"encoding/json"
	"fmt"
	"homework/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const JsonContentType = "application/json"

// middleware checking
func ContenTypeCheckingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")

		if reqContentType != JsonContentType {
			fmt.Fprintf(w, "request only allow content type application/json")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// print content header type
func ContentHeaderType(w http.ResponseWriter, r *http.Request) {
	contentHeaderType := r.Header.Get("Content-Type")
	fmt.Fprintf(w, "Hello with content type %v", contentHeaderType)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// get data from json body
	student := storage.Student{}
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		fmt.Fprintf(w, "error when parse body")
		return
	}
	// connect database and add new student to it
	db := *storage.ConnectDatabase()
	defer db.Close()

	// check if id has available
	_, productCheck := storage.Products[student.Id]
	if productCheck {
		fmt.Fprintln(w, "Id has been used")
		return
	}
	// save to database and memory
	storage.Products[student.Id] = student
	db.Create(&student)
	fmt.Println("Created succesful !")
}

func ReadStudent(w http.ResponseWriter, r *http.Request) {
	//connect database
	db := *storage.ConnectDatabase()
	defer db.Close()

	//find list item in database
	var ListStudents []storage.Student
	db.Find(&ListStudents)

	//response
	json.NewEncoder(w).Encode(ListStudents)
	fmt.Println(ListStudents)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	// get infor form body
	student := storage.Student{}
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		fmt.Fprintf(w, "error when parse body")
		return
	}

	// connect database to update product
	db := *storage.ConnectDatabase()
	defer db.Close()

	// find product by ID
	studentUpdate := storage.Student{}
	db.Find(&studentUpdate, student.Id)

	// check if the same student in storage
	if studentUpdate == student {
		fmt.Fprintln(w, "Same as the student in the storage")
		fmt.Println(w, "There is no update product")
		return
	}

	// save
	studentUpdate = student
	db.Save(&studentUpdate)
	storage.Products[student.Id] = student

	// response
	json.NewEncoder(w).Encode(studentUpdate)
	fmt.Println(studentUpdate)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	// get id of student will be delete
	vars := mux.Vars(r)
	idDelete, _ := strconv.Atoi(vars["id"])

	// check if hasn't product to delete
	_, productCheck := storage.Products[idDelete]
	if !productCheck {
		fmt.Fprintf(w, "Don't have student with id: %v", idDelete)
		fmt.Printf("Don't have student with id: %v", idDelete)
		return
	}

	// connect database to delete product
	db := *storage.ConnectDatabase()
	defer db.Close()

	// find product will be delete
	productDelete := storage.Student{}
	db.Find(&productDelete, idDelete)

	// delete student has idDelete
	delete(storage.Products, idDelete)
	db.Delete(&storage.Student{}, idDelete)
	fmt.Fprintf(w, "Student with id %v has been delete", idDelete)
	fmt.Printf("Student with id %v has been delete", idDelete)
}
