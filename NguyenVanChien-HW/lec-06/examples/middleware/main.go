package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods(http.MethodGet).Path("/hello").HandlerFunc(helloHander)

	router.Use(contentTypeCheckingMiddleware)

	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		panic(err)
	}
}

func helloHander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

const JsonContentType = "application/json"

func contentTypeCheckingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")

		if reqContentType != JsonContentType {
			fmt.Fprintf(w, "request only allow content type application/json")
			return
		}

		next.ServeHTTP(w, r)
		fmt.Fprint(w, "hello")
	})
}
