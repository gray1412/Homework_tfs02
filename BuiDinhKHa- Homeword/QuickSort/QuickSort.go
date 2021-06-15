// Package main provides ...
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := [6]int{1, 3, 2, 4, 5, 6}
	sline := array[0:6]
	fmt.Println(sline)
	fmt.Println(QuickSort(sline))
}
func QuickSort(a []int) []int {

	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])

	return a
}
