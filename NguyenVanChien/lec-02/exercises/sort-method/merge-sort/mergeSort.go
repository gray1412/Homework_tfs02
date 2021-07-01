package main

import (
	"fmt"
)

func mergeSort(a []int) []int {
	var len = len(a)
	if len == 1 {
		return a
	}

	mid := int(len / 2)
	var left = make([]int, mid)
	var right = make([]int, len-mid)

	for i := 0; i < mid; i++ {
		left[i] = a[i]
	}
	for i := mid; i < len; i++ {
		right[i-mid] = a[i]
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) (result []int) {
	result = make([]int, len(left)+len(right))
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
	if len(left) == 0 {
		for j := 0; j < len(right); j++ {
			result[i] = right[j]
			i++
		}
	} else {
		for j := 0; j < len(left); j++ {
			result[i] = left[j]
			i++
		}
	}

	return
}

func main() {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(a)
	a = mergeSort(a)
	fmt.Println(a)
}
