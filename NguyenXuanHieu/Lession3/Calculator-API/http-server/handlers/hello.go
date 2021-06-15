package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello guys")
}

type Input struct {
	OP     string `json:"OP"`
	A      int32  `json:"a"`
	B      int32  `json:"b"`
	Result int32  `json:"ans"`
	Fault  string `json:"fault"`
}

func GetEquation(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	s := query.Get("s")
	a := query.Get("a")
	b := query.Get("b")
	w.WriteHeader(200)

	new_a, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
	}
	new_b, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println(err)
	}
	var ans int32
	var fault string
	if s == "sum" {
		ans = int32(new_a) + int32(new_b)
		fault = "No Fault"
	} else if s == "sub" {
		ans = int32(new_a) - int32(new_b)
		fault = "No Fault"
	} else if s == "add" {
		ans = int32(new_a) * int32(new_b)
		fault = "No Fault"
	} else if s == "div" {
		if new_b > 0 {
			ans = int32(new_a) / int32(new_b)
			fault = "No Fault"
		} else {
			ans = 0
			fault = "Wrong b"
		}
	} else {
		ans = 0
		fault = "Wrong equation"
	}
	i := Input{s, int32(new_a), int32(new_b), ans, fault}

	bb, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}
	fmt.Fprintln(w, string(bb))
}
