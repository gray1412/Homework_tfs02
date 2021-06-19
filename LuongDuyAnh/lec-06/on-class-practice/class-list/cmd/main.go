package main

import (
	"log"
	"net/http"
	"tfs-02/lec-06/on-class-practice/class-list/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create/student", handlers.CreateData).Methods(http.MethodPost)
	router.HandleFunc("/delete/student/{id}", handlers.DeleteData).Methods(http.MethodDelete)
	router.HandleFunc("/get/students", handlers.GetAllData).Methods(http.MethodGet)
	router.HandleFunc("/get/student/{id}", handlers.GetData).Methods(http.MethodGet)
	router.HandleFunc("/update/student/{id}", handlers.UpdateData).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":8082", router))
}
