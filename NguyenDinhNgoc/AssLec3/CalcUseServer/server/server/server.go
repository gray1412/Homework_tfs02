package server

import (
	"fmt"
	"net/http"
	calc "ngoc/handlers"
)

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:5500/calc")
	// Defer function will be called when process exits
	defer func() {
		fmt.Println("Server is stopped")
	}()

	http.HandleFunc("/calc", calc.Calc)
	if err := http.ListenAndServe(":5500", nil); err != nil {
		panic("Error when running server")
	}
}
