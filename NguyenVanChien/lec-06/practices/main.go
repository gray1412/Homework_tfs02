package main

import (
	"net/http"
	"practice/handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(handler.ContenTypeCheckingMiddleware)

	router.HandleFunc("/students/{name}", handler.GetStudentByName).Methods("GET")
	router.HandleFunc("/students", handler.CreateStudent).Methods("POST")
	router.HandleFunc("/students/{id:[0-9]+}", handler.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id:[0-9]+}", handler.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/header", handler.ContentHeaderType).Methods("GET")

	err := http.ListenAndServe("0.0.0.0:8080", router)

	if err != nil {
		panic(err)
	}
}
