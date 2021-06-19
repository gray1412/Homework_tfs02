package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const JsonContentType = "application/json"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// //http handler
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	req := Person{}
	//decode json body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Fprintf(w, "error when parse body")
		return
	}
	fmt.Fprintf(w, "welcome %v, %v year old", req.Name, req.Age)
}

//middleware func
func ContentTypeCheckingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		if reqContentType != JsonContentType {
			fmt.Fprintf(w, "request only allow content type application/json")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	reqContentType := r.Header.Get("Content-Type")
	fmt.Fprintf(w, "hi with content-type %v", reqContentType)
	fmt.Println(reqContentType)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "hello %s", name)
}

func main() {
	//create router
	router := mux.NewRouter().StrictSlash(true)

	//register handler to router
	router.Methods(http.MethodGet).Path("/hi").HandlerFunc(hiHandler)
	router.Methods(http.MethodGet).Path("/hello/{name:{?:\\w+}}").HandlerFunc(helloHandler)
	router.Methods(http.MethodGet, http.MethodPost).Path("/welcome").HandlerFunc(welcomeHandler)

	//using middleware function
	// router.Use(ContentTypeCheckingMiddleware)

	//serve router on port
	err := http.ListenAndServe("localhost:8082", router)
	if err != nil {
		panic(err)
	}
	fmt.Println("server started")
}
