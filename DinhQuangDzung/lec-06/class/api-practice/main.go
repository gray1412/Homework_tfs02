package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age" `
}

var Class []Student

func init() {
	student := Student{
		Name: "test",
		Age:  0,
	}
	Class = append(Class, student)
}

func main() {

	// create mux router
	router := mux.NewRouter().StrictSlash(true)

	router.Use(contentTypeChecking)
	// register handler to router
	router.Methods("GET").Path("/students").HandlerFunc(getHandler)
	router.Methods("POST").Path("/students").HandlerFunc(addHandler)
	router.Methods("PUT").Path("/students/{id:[0-9]+}").HandlerFunc(modifyHandler)

	// serve router on port
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		panic(err)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Class)
}

func addHandler(w http.ResponseWriter, r *http.Request) {

	newStudent := Student{}

	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		fmt.Fprintf(w, "error when parsing body %v", err)
		return
	}
	Class = append(Class, newStudent)

	fmt.Fprintf(w, "Added %v, Current: %v", newStudent, Class)
}

func modifyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	newStudent := Student{}

	//Check if id is out of range
	if id >= 0 && id >= len(Class) {
		fmt.Fprintf(w, "No element at id %v", id)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		fmt.Fprintf(w, "error when parsing body %v", err)
		return
	}

	Class[id] = newStudent

	fmt.Fprintf(w, "student at id %v changed to %v", id, Class[id])
}

func contentTypeChecking(next http.Handler) http.Handler {
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
