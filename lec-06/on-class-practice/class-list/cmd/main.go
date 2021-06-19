package main

import (
	"log"
	"net/http"
	"tfs-02/lec-06/on-class-practice/class-list/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create", handlers.CreateData).Methods(http.MethodPost)
	router.HandleFunc("/delete/{id}", handlers.DeleteData).Methods(http.MethodDelete)
	router.HandleFunc("/get", handlers.GetAllData).Methods(http.MethodGet)
	router.HandleFunc("/get/{id}", handlers.GetData).Methods(http.MethodGet)
	router.HandleFunc("/update/{id}", handlers.UpdateData).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":8082", router))
}
