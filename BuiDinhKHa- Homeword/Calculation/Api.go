package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Input struct {
	Result  int `json:"result"`
	Number1 int `json: "number1"`
	Number2 int `json: "number2"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func hello(w http.ResponseWriter, rep *http.Request) {
	enableCors(&w)
	Method := rep.URL.Query().Get("type")
	number1 := rep.URL.Query().Get("number1")
	number2 := rep.URL.Query().Get("number2")
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)

	switch Method {
	case "summation":
		i := Input{
			Result:  int(int1) + int(int2),
			Number1: int(int1),
			Number2: int(int2),
		}
		bb, _ := json.Marshal(&i)
		fmt.Fprintln(w, string(bb))
		break
	case "subtraction":
		i := Input{
			Result:  int(int1) - int(int2),
			Number1: int(int1),
			Number2: int(int2),
		}
		bb, _ := json.Marshal(&i)
		fmt.Fprintln(w, string(bb))
		break
	case "multiplication":
		i := Input{
			Result:  int(int1) * int(int2),
			Number1: int(int1),
			Number2: int(int2),
		}
		bb, _ := json.Marshal(&i)
		fmt.Fprintln(w, string(bb))
		break
	case "division":
		i := Input{
			Result:  int(int1) / int(int2),
			Number1: int(int1),
			Number2: int(int2),
		}
		bb, _ := json.Marshal(&i)
		fmt.Fprintln(w, string(bb))
		break
	}
}
func main() {
	fmt.Println("Start....")
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
	defer func() {
		fmt.Println("end")
	}()
}
