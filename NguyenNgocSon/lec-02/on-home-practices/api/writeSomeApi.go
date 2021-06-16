package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/sum", apiSum)
	http.HandleFunc("/sub", apiSub)
	http.HandleFunc("/mul", apiMul)
	http.HandleFunc("/div", apiDiv)
	http.ListenAndServe(":8080", nil)
}

type Input struct {
	A string `json:"a"`
	B string `json:"b"`
}

type Sum struct {
	Sum float64 `json:"value"`
}

type Sub struct {
	Sub float64 `json:"value"`
}

type Mul struct {
	Mul float64 `json:"value"`
}

type Div struct {
	Div float64 `json:"value"`
}

func apiSum(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	a := params["a"]
	b := params["b"]
	s, err := strconv.ParseFloat(string(a[0]), 32)
	s1, err := strconv.ParseFloat(string(b[0]), 32)
	i := Sum{
		Sum: s + s1,
	}
	bb, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}
	fmt.Fprintln(w, string(bb))
}

func apiSub(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	a := params["a"]
	b := params["b"]
	// intA, err := strconv.Atoi(a[0])
	// intB, err := strconv.Atoi(b[0])
	s, err := strconv.ParseFloat(string(a[0]), 32)
	s1, err := strconv.ParseFloat(string(b[0]), 32)
	i := Sub{
		Sub: s - s1,
	}
	bb, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}
	fmt.Fprintln(w, string(bb))
}

func apiMul(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	a := params["a"]
	b := params["b"]
	s, err := strconv.ParseFloat(string(a[0]), 32)
	s1, err := strconv.ParseFloat(string(b[0]), 32)
	i := Mul{
		Mul: s * s1,
	}
	bb, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}
	fmt.Fprintln(w, string(bb))
}

func apiDiv(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	a := params["a"]
	b := params["b"]
	s, err := strconv.ParseFloat(string(a[0]), 32)
	s1, err := strconv.ParseFloat(string(b[0]), 32)
	i := Div{
		Div: s / s1,
	}
	bb, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}
	fmt.Fprintln(w, string(bb))
}
