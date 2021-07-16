package service

import (
	"course/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func ServerRun() {
	routerParent := mux.NewRouter().PathPrefix("/api/v1")
	routerUser := routerParent.PathPrefix("/users").Subrouter()


	routerUser.Methods("GET").Path("").HandlerFunc(handler.GetAllUsers)
	routerUser.Methods("GET").Path("/{id}").HandlerFunc(handler.GetUserById)
	routerUser.Methods("POST").Path("").HandlerFunc(handler.CreateUser)
	routerUser.Methods("PUT").Path("/{id}").HandlerFunc(handler.UpdateUser)
	routerUser.Methods("DELETE").Path("/{id}").HandlerFunc(handler.DeleteUser)

	err := http.ListenAndServe(":8080", routerUser)
	if err != nil {
		panic(err)
	}
}
