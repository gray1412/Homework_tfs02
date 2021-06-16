package main

import (
	sorting "Callpkgs/pkgs/sort"
	"fmt"
)

func main() {
	arr1 := []int{1, 9, 5, 3, 7, 6, 4}
	sorting.QuickSort(arr1, 0, len(arr1)-1)
	fmt.Println(arr1)

	arr2 := []int{1, 9, 5, 3, 7, 6, 4}
	sorting.BubbleSort(arr2)
	fmt.Println(arr2)
}
