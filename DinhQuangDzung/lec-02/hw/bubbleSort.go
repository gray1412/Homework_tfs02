package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}

func main() {
	arr := []int{15, 17, 20, 5, 8, 1, 6, 10, 11, 16}
	fmt.Println("Before sort:", arr)

	sortedArr := bubbleSort(arr)
	fmt.Println("After sort:", sortedArr)
}
