package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practice/storage"
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

func GetStudentByName(w http.ResponseWriter, r *http.Request) {
	// check header content type
	w.Header().Set("Content-Type", JsonContentType)

	// get name from Url
	vars := mux.Vars(r)
	name := vars["name"]
	// find student has name is vars["name"] in map students
	for _, student := range storage.Students {
		if student.Name == name {
			fmt.Fprintf(w, "Student %v has name is: %v and age is %v", student.Id, student.Name, student.Age)
			json.NewDecoder(r.Body).Decode(&student)
		}
		return
	}
	// notify if don't have student
	fmt.Fprintf(w, "Can't find student has name %v", name)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	student := storage.Student{}
	// get data from json body
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		fmt.Fprintf(w, "error when parse body")
		return
	}

	// check if id has available
	_, studentCheck := storage.Students[student.Id]
	if studentCheck {
		fmt.Fprintln(w, "Id has been used")
		return
	}
	// add new student to memory-storage
	storage.Students[student.Id] = student
	fmt.Fprintf(w, "Student: %v, has name is: %v and old is: %v", student.Id, student.Name, student.Age)
	fmt.Println("Created succesful !")
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	// update student by student Id
	vars := mux.Vars(r)
	idUpdate, _ := strconv.Atoi(vars["id"])

	// get student infor form body
	var student storage.Student
	json.NewDecoder(r.Body).Decode(&student)

	// check if the same student in storage
	if student == storage.Students[idUpdate] {
		fmt.Fprintln(w, "Same as the student in the storage")
		return
	}
	// update student
	storage.Students[idUpdate] = student
	fmt.Fprintf(w, "Student: %v, has new name is: %v and new old is: %v", student.Id, student.Name, student.Age)

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	// get id of student will be delete
	vars := mux.Vars(r)
	idDelete, _ := strconv.Atoi(vars["id"])

	// check if hasn't student to delete
	_, studentCheck := storage.Students[idDelete]
	if !studentCheck {
		fmt.Fprintf(w, "Don't have student with id: %v", idDelete)
		return
	}

	// delete student has idDelete
	delete(storage.Students, idDelete)
	fmt.Fprintf(w, "Student with id %v has been delete", idDelete)
}
