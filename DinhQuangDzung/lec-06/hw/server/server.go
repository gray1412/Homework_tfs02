package server

import (
	"fmt"
	"lec06-hw/controller"
	"lec06-hw/middleware"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func RunServer() {
	fmt.Println("Server opened at port 8080...")
	defer fmt.Println("Server stopped!")
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middleware.ContentTypeChecking)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})
	handler := c.Handler(router)

	router.Methods("GET").Path("/students").HandlerFunc(controller.GetAll)
	router.Methods("GET").Path("/students/{id:[0-9]+}").HandlerFunc(controller.GetById)
	router.Methods("POST").Path("/students").HandlerFunc(controller.AddOne)
	router.Methods("PUT").Path("/students/{id:[0-9]+}").HandlerFunc(controller.UpdateById)
	router.Methods("DELETE").Path("/students/{id:[0-9]+}").HandlerFunc(controller.DeleteById)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}
