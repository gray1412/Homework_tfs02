package main

import (
	"fmt"
)

func calc(method string, a int, b int) int {
	var result int
	switch method {
	case "sum":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		result = a / b
	}
	return result
}

func main() {
	cal1 := calc("sum", 10, 5)
	cal2 := calc("sub", 10, 5)
	cal3 := calc("mul", 10, 5)
	cal4 := calc("div", 10, 5)

	fmt.Println("Result of calculation is", cal1)
	fmt.Println("Result of calculation is", cal2)
	fmt.Println("Result of calculation is", cal3)
	fmt.Println("Result of calculation is", cal4)
}
