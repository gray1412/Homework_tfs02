package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello guys")
}

type Input struct {
	Num1     string `json:"num1"`
	Num2     string `json:"num2"`
	Operator string `json:"operator"`
	Result   int    `json:"result"`
}

func GetEquation(w http.ResponseWriter, r *http.Request) {
	jsonFeed, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}
	i := Input{}

	json.Unmarshal([]byte(jsonFeed), &i)
	fmt.Println("num1", i.Num1, "num2", i.Num2, "operator", i.Operator)

	num1, err := strconv.Atoi(i.Num1)
	if err != nil {
		fmt.Println(err)
	}

	num2, err := strconv.Atoi(i.Num2)
	if err != nil {
		fmt.Println(err)
	}

	if i.Operator == "+" {
		i.Result = num1 + num2
	}

	if i.Operator == "-" {
		i.Result = num1 - num2
	}

	if i.Operator == "*" {
		i.Result = num1 * num2
	}

	if i.Operator == "/" {
		i.Result = num1 / num2
	}

	if i.Operator == "" {
		i.Result = num1
	}

	fmt.Println(i)
	bb, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}
	fmt.Fprintln(w, string(bb))
}
