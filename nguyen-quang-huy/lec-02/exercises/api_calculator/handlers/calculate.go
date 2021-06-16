package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"tfs-02/lec-02/exercises/api_calculator/models"
)

func HandlersCalculate(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	op := params["op"]
	rawA := params["a"]
	rawB := params["b"]

	a, err := strconv.Atoi(rawA[0])
	if err != nil {
		checkError(w, errors.New("a must be number"))
		return
	}

	b, err := strconv.Atoi(rawB[0])
	if err != nil {
		checkError(w, errors.New("b must be number"))
		return
	}

	resultCalculate, err := calculate(op[0], a, b)
	if err != nil {
		checkError(w, err)
		return
	}

	input := models.Input{
		Op:     op[0],
		A:      a,
		B:      b,
		Result: resultCalculate,
	}

	respone, err := json.Marshal(&input)
	if err != nil {
		checkError(w, err)
		return
	}
	fmt.Fprintln(w, string(respone))
}
func checkError(w http.ResponseWriter, err error) {
	fmt.Println(err)
	fmt.Fprintln(w, err)
}

func calculate(op string, a, b int) (float64, error) {
	switch op {
	case "add":
		return float64(a + b), nil
	case "sub":
		return float64(a - b), nil
	case "mul":
		return float64(a * b), nil
	case "div":
		if b == 0 {
			return 0, errors.New("b must be different from 0")
		}
		return float64(a) / float64(b), nil
	default:
		return 0, errors.New("op doesn't match")
	}
}
