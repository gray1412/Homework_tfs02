package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Input struct {
	Operator string `json:"op"`
	NumA     string `json:"a"`
	NumB     string `json:"b"`
	Result   string `json:"result"`
}

func calc(method string, a int, b int) int {
	var result int
	switch method {
	case "sum":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		result = a / b
	}
	return result
}

func handleCalc(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	op := params["op"][0]
	a, _ := strconv.Atoi(params["a"][0])
	b, _ := strconv.Atoi(params["b"][0])
	result := calc(op, a, b)

	// fmt.Println("Result: ", result)

	i := Input{
		Operator: op,
		NumA:     strconv.Itoa(a),
		NumB:     strconv.Itoa(b),
		Result:   strconv.Itoa(result),
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
	http.HandleFunc("/calc", handleCalc)
	http.ListenAndServe(":8090", nil)
}
