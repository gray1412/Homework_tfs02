package server

import (
	"fmt"
	"homework/handler"
	"homework/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:8080")
	defer func() {
		fmt.Println("Server is stopped")
	}()

	router := mux.NewRouter().StrictSlash(true)
	//create product table
	// storage.CreateTableProduct()
	storage.CreateTableStudent()
	router.Use(handler.ContenTypeCheckingMiddleware)

	router.HandleFunc("/students", handler.ReadStudent).Methods("GET")
	router.HandleFunc("/students", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/students/{id:[0-9]+}", handler.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id:[0-9]+}", handler.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/header", handler.ContentHeaderType).Methods("GET")

	err := http.ListenAndServe("0.0.0.0:8080", router)

	if err != nil {
		panic(err)
	}
}
