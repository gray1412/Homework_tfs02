package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// create mux router
	router := mux.NewRouter().StrictSlash(true)

	// register handler to router
	router.Methods(http.MethodGet).Path("/hello").HandlerFunc(helloHander)

	// serve router on port
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		panic(err)
	}
}

func helloHander(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // vars is map[string]string

	name := vars["name"]

	fmt.Fprintf(w, "hi %s", name)
}
