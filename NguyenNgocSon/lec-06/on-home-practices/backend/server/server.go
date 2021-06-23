package server

import (
	"fmt"
	"net/http"

	"lec06/handler"
	// "lec06/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func RunServer() {
	fmt.Println("Server opened at port 8000...")
	defer fmt.Println("Server stopped!")
	router := mux.NewRouter().StrictSlash(true)

	// router.Use(middleware.ContentTypeChecking)

	router.Methods("GET").Path("/api/members").HandlerFunc(handler.GetMembers)
	router.Methods("GET").Path("/api/members/{id:(?:\\d+)}").HandlerFunc(handler.GetOneMember)
	router.Methods("POST").Path("/api/members").HandlerFunc(handler.CreateMember)
	router.Methods("PUT").Path("/api/members/{id:(?:\\d+)}/profile").HandlerFunc(handler.UpdateMember)
	router.Methods("DELETE").Path("/api/members/{id:(?:\\d+)}/delete").HandlerFunc(handler.DeleteMember)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)

	http.ListenAndServe(":8000", handler)
}
