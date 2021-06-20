package server

import (
	"net/http"
	"tfs/tfs-api-mysql/handlers"
	"tfs/tfs-api-mysql/middleware"

	"github.com/gorilla/mux"
)

func ServerRun() {
	routerParent := mux.NewRouter()
	router := routerParent.PathPrefix("/students").Subrouter()
	router.Use(middleware.ContentTypeCheckingMiddleware)
	router.Methods("POST").Path("").HandlerFunc(handlers.CreatePerson)
	router.Methods("GET").Path("").HandlerFunc(handlers.GetAllPersons)
	router.Methods("PUT").Path("/{id:[0-9]+}").HandlerFunc(handlers.UpdatePerson)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
