package main

import (
	"log"
	"net/http"

	"NDT/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/create/student", handlers.CreateData).Methods(http.MethodPost)
	r.HandleFunc("/delete/student/{id}", handlers.DeleteData).Methods(http.MethodDelete)
	r.HandleFunc("/getall/student", handlers.GetAllData).Methods(http.MethodGet)
	r.HandleFunc("/get/student/{id}", handlers.GetData).Methods(http.MethodGet)
	r.HandleFunc("/update/student/{id}", handlers.UpdateData).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":8081", r))
}
