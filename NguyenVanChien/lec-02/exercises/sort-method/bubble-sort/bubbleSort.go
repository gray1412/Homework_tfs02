package main

import (
	"fmt"
)

func bubbleSort(a []int) []int {
	var len int = len(a)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
	}
	return a
}

func main() {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(a)
	a = bubbleSort(a)
	fmt.Println(a)
}
