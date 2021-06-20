package server

import (
	"fmt"
	"net/http"

	"lec06-hw/handler"
	"lec06-hw/middleware"

	"github.com/gorilla/mux"
)

func RunServer() {
	fmt.Println("Server opened at port 8080...")
	defer fmt.Println("Server stopped!")
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middleware.ContentTypeChecking)

	router.Methods("GET").Path("/students").HandlerFunc(handler.GetAll)
	router.Methods("GET").Path("/students/{id:[0-9]+}").HandlerFunc(handler.GetById)
	router.Methods("POST").Path("/students").HandlerFunc(handler.AddOne)
	router.Methods("PUT").Path("/students/{id:[0-9]+}").HandlerFunc(handler.UpdateById)
	// router.Methods("DELETE").Path("/students/{id:[0-9]+}").HandlerFunc(handler.DeleteById)
	// router.Methods("DELETE").Path("/students}").HandlerFunc(handler.DeleteAll)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
