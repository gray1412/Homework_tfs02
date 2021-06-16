package server

import (
	"fmt"
	"net/http"
	"tfs-02/lec-03/exercises/caculator/backend/handlers"
)

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:5500/calculate")
	// Defer function will be called when process exits
	defer func() {
		fmt.Println("Server is stopped")
	}()

	http.HandleFunc("/calculate", handlers.HandlersCalculate)

	if err := http.ListenAndServe(":5500", nil); err != nil {
		panic("Error when running server")
	}
}
