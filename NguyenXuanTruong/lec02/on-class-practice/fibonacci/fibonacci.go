package main

import "fmt"

// Phép tính với 2 số tự nhiên
func caculate(s string, a int, b int) int {
	switch s {
	case "add":
		return a + b
	case "sub":
		return a - b
	case "mul":
		return a * b
	case "div":
		return a / b
	default:
		return 0
	}
}

// Hàm in ra dãy fibnacci
func showFibonaci(n int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}

// // Kiểm tra panic có defer không
// func showSt(arr [2]int) {
// 	defer fmt.Println("loi")
// 	fmt.Println(arr[2])
// }

func main() {
	fmt.Println(caculate("add", 2, 3))
	showFibonaci(9)
}
