package main

import (
	findNum "Callpkgs/pkgs/findNum"
	"fmt"
)

func showResult(filename string, element int) {
	if findNum.Checking(filename, element) == true {
		fmt.Println("%v tồn tại trong file", element)
	} else {
		fmt.Println("%v không tồn tại trong file", element)
	}
}

func main() {
	e := 9
	showResult("input.txt", e)
}
