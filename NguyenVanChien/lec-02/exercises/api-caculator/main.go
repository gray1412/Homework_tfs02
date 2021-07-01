package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Input struct {
	Result  int64 `json:"result"`
	Number1 int64 `json: "number1"`
	Number2 int64 `json: "number2"`
}

func Caculator(w http.ResponseWriter, rep *http.Request) {
	Method := rep.URL.Query().Get("type")
	number1 := rep.URL.Query().Get("number1")
	number2 := rep.URL.Query().Get("number2")
	int1, _ := strconv.ParseInt(number1, 6, 12)
	int2, _ := strconv.ParseInt(number2, 6, 12)

	switch Method {
	case "sum":
		i := Input{
			Result:  int64(int1) + int64(int2),
			Number1: int64(int1),
			Number2: int64(int2),
		}
		bb, _ := json.Marshal(&i)
		fmt.Fprintln(w, string(bb))
		break
	case "sub":
		i := Input{
			Result:  int64(int1) - int64(int2),
			Number1: int64(int1),
			Number2: int64(int2),
		}
		bb, _ := json.Marshal(&i)
		fmt.Fprintln(w, string(bb))
		break
	case "mul":
		i := Input{
			Result:  int64(int1) * int64(int2),
			Number1: int64(int1),
			Number2: int64(int2),
		}
		bb, _ := json.Marshal(&i)
		fmt.Fprintln(w, string(bb))
		break
	case "div":
		i := Input{
			Result:  int64(int1) / int64(int2),
			Number1: int64(int1),
			Number2: int64(int2),
		}
		bb, _ := json.Marshal(&i)
		fmt.Fprintln(w, string(bb))
		break
	}
}
func main() {
	fmt.Println("Start....")
	http.HandleFunc("/caculator", Caculator)
	http.ListenAndServe(":8080", nil)
	defer func() {
		fmt.Println("end")
	}()
}
