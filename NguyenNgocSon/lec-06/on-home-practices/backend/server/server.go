package server

import (
	"fmt"
	"net/http"

	"lec06/handler"
	"lec06/middleware"

	"github.com/gorilla/mux"
)

func RunServer() {
	fmt.Println("Server opened at port 8000...")
	defer fmt.Println("Server stopped!")
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middleware.ContentTypeChecking)

	router.Methods("GET").Path("/api/members").HandlerFunc(handler.GetMembers)
	router.Methods("POST").Path("/api/members").HandlerFunc(handler.CreateMember)
	router.Methods("PUT").Path("/api/members/{id:(?:\\d+)}/profile").HandlerFunc(handler.UpdateMember)
	router.Methods("DELETE").Path("/api/members/{id:(?:\\d+)}/delete").HandlerFunc(handler.DeleteMember)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}
