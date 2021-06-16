package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Input struct {
	Op string  `json:"op"`
	A  string  `json:"a"`
	B  string  `json:"b"`
	R  float64 `json:"result"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	op := params["s"]
	a := params["a"]
	b := params["b"]
	fmt.Printf("%v %v %v\n", op, a, b)

	m, err := strconv.ParseFloat(a[0], 64) //convert string to float
	n, err := strconv.ParseFloat(b[0], 64)

	//sum
	sum := Input{
		Op: op[0],
		A:  a[0],
		B:  b[0],
		R:  (m + n),
	}
	sum1, err := json.Marshal(&sum)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}

	//sub
	sub := Input{
		Op: op[0],
		A:  a[0],
		B:  b[0],
		R:  (m - n),
	}
	sub1, err := json.Marshal(&sub)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}

	//mul
	mul := Input{
		Op: op[0],
		A:  a[0],
		B:  b[0],
		R:  (m * n),
	}
	mul1, err := json.Marshal(&mul)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}

	//div
	div := Input{
		Op: op[0],
		A:  a[0],
		B:  b[0],
		R:  (m / n),
	}
	div1, err := json.Marshal(&div)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}

	switch op[0] {
	case "sum":
		fmt.Fprintln(w, string(sum1))
	case "sub":
		fmt.Fprintln(w, string(sub1))
	case "mul":
		fmt.Fprintln(w, string(mul1))
	case "div":
		fmt.Fprintln(w, string(div1))
	default:
		fmt.Fprintln(w, "ha")
	}

}

func main() {
	fmt.Println("start")
	defer func() {
		fmt.Println("end")
	}()
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
