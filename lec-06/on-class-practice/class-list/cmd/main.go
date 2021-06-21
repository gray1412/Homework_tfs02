package main

import (
	"log"
	"net/http"
	"tfs-02/lec-06/on-class-practice/class-list/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/students", handlers.CreateData).Methods(http.MethodPost)
	router.HandleFunc("/api/students/{id:(?:\\d+)}", handlers.DeleteData).Methods(http.MethodDelete)
	router.HandleFunc("/api/students", handlers.GetAllData).Methods(http.MethodGet)
	router.HandleFunc("/api/students/{id:(?:\\d+)}", handlers.GetData).Methods(http.MethodGet)
	router.HandleFunc("/api/students/{id:(?:\\d+)}", handlers.UpdateData).Methods(http.MethodPut)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)

	log.Fatal(http.ListenAndServe(":8082", handler))

}
