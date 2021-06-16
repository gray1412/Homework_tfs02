package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
)

type Math struct {
	Type string
	A int
	B int
	Result int
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " %s!", r.URL.Path[1:])
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	Type :=  r.URL.Query().Get("type")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	numa, _ := strconv.Atoi(a)
	numb, _ := strconv.Atoi(b)
	switch Type {
	case "sum":
		math := &Math{Type: Type, Result: numa + numb, A:numa, B:numb}
		re, err := json.Marshal(math)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%v", string(re))
	case "sub":
		math := &Math{Type: Type, Result: numa - numb, A:numa, B:numb}
		re, err := json.Marshal(math)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%v", string(re))
	case "mul":
		math := &Math{Type: Type, Result: numa * numb, A:numa, B:numb}
		re, err := json.Marshal(math)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%v", string(re))
	case "div":
		math := &Math{Type: Type, Result: numa / numb, A:numa, B:numb}
		re, err := json.Marshal(math)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%v", string(re))
	case "mod":
		math := &Math{Type: Type, Result: numa % numb, A:numa, B:numb}
		re, err := json.Marshal(math)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%v", string(re))
	}
}

func main() {
	http.HandleFunc("/", queryParams)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
