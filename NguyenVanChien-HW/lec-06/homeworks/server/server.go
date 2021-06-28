package server

import (
	"fmt"
	"homework/handler"
	"homework/storage"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	}
	//else {
	// 	db.DropTable(&storage.Product{})
	// 	db.CreateTable(&storage.Product{})
	// }

	// router.Use(handler.ContenTypeCheckingMiddleware)

	router.HandleFunc("/products", handler.ReadProduc).Methods("GET")
	router.HandleFunc("/products", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id:[0-9]+}", handler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id:[0-9]+}", handler.DeleteProduct).Methods("DELETE")
	// router.HandleFunc("/header", handler.ContentHeaderType).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
	}).Handler(router)
	http.ListenAndServe(":8000", handler)

}
