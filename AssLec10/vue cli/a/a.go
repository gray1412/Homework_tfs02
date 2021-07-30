package main

import (
	"fmt"
	// "strings"
)

type date struct {
	a, b int
}

func f(m map[int]date) []date {
	arr := make([]date, 25)
	for k, v := range m{
		fmt.Println(k)
		arr[k] = v
	}
	return arr
}

func main() {
	m := make(map[int]date)
	for i:= 1; i<25; i++{
		m[i] = date{i, i+1}
	}
	
	// var arr = make([]date, 25)
	// for key, val := range m {
	// 	fmt.Println(key)
	// 	arr[key] = val
		
	// }
	fmt.Println(f(m))
}