package main

import (
	"fmt"
	"math/rand"
)

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left := 0
	right := len(a) - 1
	pivot := rand.Int() % len(a)
	//Move the pivot to the right
	a[pivot], a[right] = a[right], a[pivot]
	//Move elements smaller than the pivot to the left
	for i := 0; i < len(a); i++ {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	//Swap the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])
	return a
}

func main() {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(a)
	fmt.Println(len(a))
	a = quickSort(a)
	fmt.Println(a)
}
