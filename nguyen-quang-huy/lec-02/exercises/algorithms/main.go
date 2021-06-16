package main

import (
	"fmt"
	sortAlgorithm "tfs-02/lec-02/exercises/algorithms/sort-algorithm"
)

func main() {
	arr := []int{9, 1, 4, 5, 8, 2, 4, 6}
	// BubbleSort
	sortAlgorithm.BubbleSort(arr)
	fmt.Println(arr)

	// QuickSort
	sortAlgorithm.QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	// MergeSort
	arr = sortAlgorithm.MergeSort(arr)
	fmt.Println(arr)
}
