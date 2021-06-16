package server

import (
	"fmt"
	"net/http"
	"tfs-02/lec-02/exercises/api_calculator/handlers"
)

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:8080/")
	// Defer function will be called when process exits
	defer func() {
		fmt.Println("Server is stopped")
	}()

	http.HandleFunc("/calculate", handlers.HandlersCalculate) 

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Error when running server")
	}
}
