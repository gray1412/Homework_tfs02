package main

import (
	"CallPkgs/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// create mux router
	router := mux.NewRouter().StrictSlash(true)

	// register handler to router
	router.HandleFunc("/create/student", handlers.Create).Methods(http.MethodPost)
	router.HandleFunc("/delete/student/{id}", handlers.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/get/students", handlers.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/get/student/{id}", handlers.GetSingle).Methods(http.MethodGet)
	router.HandleFunc("/update/student/{id}", handlers.Update).Methods(http.MethodPut)

	// serve router on port
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		panic(err)
	}

}
