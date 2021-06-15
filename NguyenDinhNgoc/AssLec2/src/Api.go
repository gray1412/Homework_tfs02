package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Input struct {
	Op     string  `json:"op"`
	A      string  `json:"a"`
	B      string  `json:"b"`
	Result float64 `json:"Result"`
}

func handelQuery(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	op := params["s"]
	a := params["a"]
	b := params["b"]

	i := Input{
		Op:     op[0],
		A:      a[0],
		B:      b[0],
		Result: 0,
	}
	//convert string to int64
	aa, err := strconv.ParseInt(i.A, 10, 64)
	bb, err2 := strconv.ParseInt(i.B, 10, 64)
	if err == nil && err2 == nil { //convert successful
		switch i.Op {
		case "add":
			i.Result = float64(aa) + float64(bb)
		case "sub":
			i.Result = float64(aa) - float64(bb)
		case "mul":
			i.Result = float64(aa) * float64(bb)
		case "div":
			i.Result = float64(aa) / float64(bb)

		default:
			i.Op = i.Op + ": exception Operation !"
		}
	} else {
		//mark exception
		if err != nil {
			i.A = i.A + ": exception value !"
			fmt.Println(i.Op, " ", i.A, " ", i.B)
		}
		if err2 != nil {
			i.B = i.B + ": exception value !"
		}
	}
	fmt.Println(i.Op, " ", i.A, " ", i.B, " ", i.Result)

	//convert to json
	r, err3 := json.Marshal(&i)
	fmt.Println(string(r))

	//ResponseWriter
	if err3 != nil {
		fmt.Println(err3)
		fmt.Fprintln(w, "Error")
		return
	} else {
		fmt.Fprintln(w, string(r))
	}

}
func main() {
	fmt.Println("start")
	defer func() {
		fmt.Println("end")
	}()

	http.HandleFunc("/calc", handelQuery)

	http.ListenAndServe(":8090", nil)
}
