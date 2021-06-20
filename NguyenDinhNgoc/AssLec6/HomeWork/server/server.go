package server

import (
	"asslec6p2/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer() {

	fmt.Println("Starting server. Please open http://localhost:8080")
	// Defer function will be called when process exits
	defer func() {
		fmt.Println("Server is stopped")
	}()

	router := mux.NewRouter().StrictSlash(true)

	router.Use(handlers.ContentTypeCheckingMiddleware)

	router.Methods(http.MethodGet).Path("/students/{id:[0-9]+}").HandlerFunc(handlers.ReadbyID)
	router.Methods(http.MethodGet).Path("/students").HandlerFunc(handlers.ReadAll)
	router.Methods(http.MethodPost).Path("/students").HandlerFunc(handlers.Create)
	router.Methods(http.MethodPut).Path("/students").HandlerFunc(handlers.Update)
	router.Methods(http.MethodDelete).Path("/students/{id:[0-9]+}").HandlerFunc(handlers.Delete)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
