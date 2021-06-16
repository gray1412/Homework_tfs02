package main

import (
	f "CallPkg/pkgs/findNumber"
	"fmt"
)

func printResult(a bool, b int64) {
	if a {
		fmt.Println("Tìm thấy ", b, " trong file !")
	} else {
		fmt.Println("Không tìm thấy ", b, " trong file !")
	}
}

func main() {
	var a int64 = 10
	ret := f.FindNumbyWay1(a, "file/test.txt")
	printResult(ret, a)

	var b int64 = 6
	ret2 := f.FindNumByHashing(b, "file/test.txt")
	printResult(ret2, b)
}
