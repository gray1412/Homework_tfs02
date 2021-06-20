package main

import (
	"asslec6/pkgs/handler"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(handler.ContentTypeCheckingMiddleware)

	router.Methods(http.MethodGet).Path("/students/{id:[0-9]+}").HandlerFunc(handler.ReadbyID)
	router.Methods(http.MethodGet).Path("/students").HandlerFunc(handler.ReadAll)
	router.Methods(http.MethodPost).Path("/students").HandlerFunc(handler.Create)
	router.Methods(http.MethodPut).Path("/students").HandlerFunc(handler.Update)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
