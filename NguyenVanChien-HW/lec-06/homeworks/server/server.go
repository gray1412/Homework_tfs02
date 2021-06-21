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
	db := *storage.ConnectDatabase()
	// defer db.Close()
	if (!db.HasTable(&storage.Product{})) {
		db.CreateTable(&storage.Product{})
	} else {
		db.DropTable(&storage.Product{})
		db.CreateTable(&storage.Product{})
	}

	router.Use(handler.ContenTypeCheckingMiddleware)

	router.HandleFunc("/products/{name}", handler.ReadProduc).Methods("GET")
	router.HandleFunc("/products", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id:[0-9]+}", handler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id:[0-9]+}", handler.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/header", handler.ContentHeaderType).Methods("GET")

	err := http.ListenAndServe("0.0.0.0:8080", router)

	if err != nil {
		panic(err)
	}
}
