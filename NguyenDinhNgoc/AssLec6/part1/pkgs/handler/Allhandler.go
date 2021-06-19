package handler

import (
	"asslec6/pkgs/data"
	"asslec6/pkgs/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//response with http Status Code and Message or Object
func responseWithJson(w http.ResponseWriter, httpStatusCode int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(httpStatusCode)
	//response with json
	json.NewEncoder(w).Encode(object)
}

func ReadbyID(w http.ResponseWriter, r *http.Request) {
	//get ID
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//find
	student, err := data.Students[id]
	if !err {
		fmt.Println(err)
		fmt.Fprintf(w, "No rerult !")
		return
	}

	//response
	responseWithJson(w, http.StatusOK, student)

	fmt.Println(student)
}

func ReadAll(w http.ResponseWriter, r *http.Request) {
	//response
	responseWithJson(w, http.StatusOK, data.Students)

	fmt.Println(data.Students)
}
func Create(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var newStudent storage.Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	idCreate := newStudent.ID
	//check
	_, err2 := data.Students[idCreate]
	if err2 {
		responseWithJson(w, http.StatusOK, map[string]string{"message": "ID has been used !"})
		return
	}

	//create
	data.Students[idCreate] = newStudent

	//response
	responseWithJson(w, http.StatusCreated, newStudent)

	//print
	fmt.Println(data.Students)
}
func Update(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var updateStudent storage.Student
	err := json.NewDecoder(r.Body).Decode(&updateStudent)
	if err != nil {
		responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	idNeedUpdate := updateStudent.ID
	//check
	_, err2 := data.Students[idNeedUpdate]
	if !err2 {
		responseWithJson(w, http.StatusOK, map[string]string{"message": "There is no such student !"})
		return
	}

	//update
	data.Students[idNeedUpdate] = updateStudent

	//response
	responseWithJson(w, http.StatusOK, updateStudent)

	//print
	fmt.Println(data.Students)
}

//middleware
const JsonContentType = "application/json"

func ContentTypeCheckingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		if reqContentType != JsonContentType {
			responseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Request only allow content type: application/json !"})
			return
		}
		next.ServeHTTP(w, r)
	})

}
