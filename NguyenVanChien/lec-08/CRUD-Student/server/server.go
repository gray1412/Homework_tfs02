package server

import (
	"fmt"
	"net/http"
	"student/handler"
	"student/storage"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Runserver() {
	fmt.Println("Starting server. Please open http://localhost:8080")
	defer func() {
		fmt.Println("Server is stopped")
	}()

	router := mux.NewRouter().StrictSlash(true)

	router.Use(handler.ContenTypeCheckingMiddleware)

	storage.CreateTableProduct()

	router.HandleFunc("/students/{name}", handler.GetStudentByName).Methods("GET")
	router.HandleFunc("/students", handler.CreateStudent).Methods("POST")
	router.HandleFunc("/students/{id:[0-9]+}", handler.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id:[0-9]+}", handler.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/header", handler.ContentHeaderType).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)
	http.ListenAndServe(":8080", handler)
}
