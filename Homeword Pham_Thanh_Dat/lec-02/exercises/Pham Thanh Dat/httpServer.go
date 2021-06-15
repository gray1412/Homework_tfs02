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
	Result int
	A int
	B int
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func queryParams(w http.ResponseWriter, r *http.Request) {
	Type :=  r.URL.Query().Get("type")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	var c = 0
	var d = 0
	c, _ = strconv.Atoi(a)
	d, _ = strconv.Atoi(b)
	switch Type {
	case "sum":
		math := &Math{Type: Type, Result: c + d, A:c, B:d}
		re, err := json.Marshal(math)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%v", string(re))
	case "sub":
		math := &Math{Type: Type, Result: c - d, A:c, B:d}
		re, err := json.Marshal(math)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%v", string(re))
	case "mul":
		math := &Math{Type: Type, Result: c / d, A:c, B:d}
		re, err := json.Marshal(math)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%v", string(re))
	case "div":
		math := &Math{Type: Type, Result: c * d, A:c, B:d}
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
/*func main() {
	type Person struct {
		Sum string `json:"sum"`
		A int `json:"a"`
		B int `json:"b"`
	}
	var listPerson []Person
	str := fmt.Sprintf(`[{"sum":"%v", "a":%v, "b":%v}]`, 4,2,3)
	err := json.Unmarshal([]byte(str), &listPerson)
	if err != nil {
		return
	}
	var data []byte
	data, err = json.Marshal(&listPerson)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}*/
