package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Input struct {
	Op string `json:"op"`
	A  string `json:"a"`
	B  string `json:"b"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	op := params["s"]
	a := params["a"]
	b := params["b"]
	fmt.Printf("%v %v %v\n", op, a, b)
	i := Input{
		Op: op[0],
		A:  a[0],
		B:  b[0],
	}
	bb, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}
	fmt.Fprintln(w, string(bb))
}
func main() {
	fmt.Println("start")
	defer func() {
		fmt.Println("end")
	}()
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
