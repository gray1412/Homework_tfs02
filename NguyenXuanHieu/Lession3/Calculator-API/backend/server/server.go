package server

import (
	"fmt"
	"hieu/handlers" // import local package
	"log"
	"net/http"
)

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:5000/hello")
	// Defer function will be called when process exits
	defer func() {
		fmt.Println("Server is stopped")
	}()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", handlers.GetEquation) // simple hello
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}

	// run server
	// if err := http.ListenAndServe(":5500", nil); err != nil {
	// 	panic("Error when running server")
	// }
}
