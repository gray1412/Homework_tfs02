package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.Methods(http.MethodGet).Path("/").HandlerFunc(testHandler)
	r.Use(contentTypeCheckingMiddleware)

	err := http.ListenAndServe("0.0.0.0:8090", r)
	if err != nil {
		panic(err)
	}
}
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

const JsonContentType = "application/json"

// middleware func
func contentTypeCheckingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")

		if reqContentType != JsonContentType {
			fmt.Fprintf(w, "request only allow content type application/json")
			return
		}

		next.ServeHTTP(w, r)
	})
}
