package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Input struct {
	OP int
	X  int
	Y  int
}

func main() {
	fmt.Println("Start!")
	defer func() {
		fmt.Println("End!")
	}()
	http.HandleFunc("/operation", operation)
	http.ListenAndServe(":8080", nil)
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func operation(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	x := r.URL.Query().Get("x")
	y := r.URL.Query().Get("y")
	op := r.URL.Query().Get("types")

	int1, _ := strconv.Atoi(x)
	int2, _ := strconv.Atoi(y)

	switch op {
	case "sum":
		i := Input{
			OP: int(int1) + int(int2),
			X:  int(int1),
			Y:  int(int2),
		}
		b, _ := json.Marshal(&i)
		fmt.Fprintln(w, string(b))
		break
	case "sub":
		i := Input{
			OP: int(int1) - int(int2),
			X:  int(int1),
			Y:  int(int2),
		}
		b, _ := json.Marshal(&i)
		fmt.Fprintln(w, string(b))
		break
	case "mul":
		i := Input{
			OP: int(int1) * int(int2),
			X:  int(int1),
			Y:  int(int2),
		}
		b, _ := json.Marshal(&i)
		fmt.Fprintln(w, string(b))
		break
	case "div":
		i := Input{
			OP: int(int1) / int(int2),
			X:  int(int1),
			Y:  int(int2),
		}
		b, _ := json.Marshal(&i)
		fmt.Fprintln(w, string(b))
		break
	}

}
