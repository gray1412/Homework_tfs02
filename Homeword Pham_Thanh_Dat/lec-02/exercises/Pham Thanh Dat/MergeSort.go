package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 11, 13, 4, 9, 5, 7, 6, 1}
	fmt.Print(Separate(s))
}
func Separate(s []int) []int {
	var len = len(s)
	if len == 1 {
		return s
	}
	node := int(len / 2)
	left := make([]int, node)
	right := make([]int, len-node)
	for i := 0; i < len; i++ {
		if i < node {
			left[i] = s[i]
		} else {
			right[i-node] = s[i]
		}
	}
	return mergeSort(Separate(left), Separate(right))
}
func mergeSort(left, right []int) (result []int) {
	result = make([]int, len(left) + len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return
}

