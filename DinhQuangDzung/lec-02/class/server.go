package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Input struct {
	Operator string `json:"op"`
	NumA     string `json:"a"`
	NumB     string `json:"b"`
}

func sayHello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello WORLD!")
}

func handleQuery(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query() //params is a map
	// fmt.Println(params)
	op := params["op"][0] //value from map is a slice, [0] to get value from slice
	a := params["a"][0]
	b := params["b"][0]

	fmt.Printf("Params: %v, %v, %v\n", op, a, b)

	i := Input{
		Operator: op,
		NumA:     a,
		NumB:     b,
	}

	data, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(res, "Error")
		return
	}
	fmt.Fprintln(res, string(data))
}

func main() {
	fmt.Println("Server started!")
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/query", handleQuery)
	http.ListenAndServe(":8090", nil)
}
